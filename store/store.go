package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	pb "github.com/Sayan-995/vidquizgen/bindings"
	t "github.com/Sayan-995/vidquizgen/types"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client *genai.Client
	em     *genai.EmbeddingModel
)

func init() {
	var err error
	godotenv.Load()
	Client, err = genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalf("Failed to create Generative AI Client: %v", err)
	}

	em = Client.EmbeddingModel("text-embedding-004")
	if em == nil {
		log.Fatalf("Failed to initialize embedding model")
	}
}


type Pgstore struct {
	DB *sql.DB
}

func ScanIntoQuestion(row *sql.Rows) (t.Question, error) {
	var q t.Question
	err := row.Scan(
		&q.QID,
		&q.Title,
		&q.TitleSlug,
		&q.Difficulty,
		&q.AcceptanceRate,
		&q.PaidOnly,
		&q.TopicTags,
		&q.CategorySlug,
	)
	return q, err
}

func (s *Pgstore) GetTopQuestions(text string) ([]string, error) {
	res, err := em.EmbedContent(context.Background(), genai.Text(text))
	if err != nil {
		return nil, err
	}
	query := `SELECT content FROM embeddings ORDER BY embedding <-> $1 LIMIT 5`
	rows, err := s.DB.Query(query, floatSliceToVector(res.Embedding.Values))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		resp = append(resp, content)
	}
	return resp, nil
}

func Create() (*Pgstore, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	queries := []string{
		`CREATE TABLE IF NOT EXISTS Questions (
			qid SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			titleslug TEXT NOT NULL,
			difficulty TEXT CHECK (difficulty IN ('Easy', 'Medium', 'Hard')),
			acceptancerate DECIMAL(5,2),
			paidonly BOOLEAN DEFAULT FALSE,
			topictags TEXT,
			categoryslug TEXT
		)`,
		`CREATE EXTENSION IF NOT EXISTS vector`,
		`CREATE TABLE IF NOT EXISTS embeddings (
			id SERIAL PRIMARY KEY,
			content TEXT,
			embedding vector(768)
		)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, err
		}
	}
	return &Pgstore{DB: db}, nil
}

func GetDB() (*sql.DB, error) {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URI"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (s *Pgstore) GetText() ([]string, error) {
	query := `select * from questions`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	Client := pb.NewStatementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	var problems []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	cnt:=0;
	for rows.Next() {
		cnt++
		q, err := ScanIntoQuestion(rows)
		if err != nil {
			return nil, err
		}
		if q.PaidOnly != nil && *q.PaidOnly {
			continue
		}

		wg.Add(1)
		go func(q t.Question) {
			defer wg.Done()
			resp, err := Client.GetStatement(ctx, &pb.ProblemRequest{TitleSlug: *q.TitleSlug})
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}
			title:=""
			if(q.Title==nil){
				// fmt.Println(q);
			}else{
				title=*q.Title
			}
			tt:=""
			if(q.TopicTags==nil){
				// fmt.Println(q);
			}else{
				tt=*q.TopicTags
			}
			problem := fmt.Sprintf("Title: %s\nStatement: %s\nTopics: %s", title, resp.Statement, tt)
			mu.Lock()
			problems = append(problems, problem)
			mu.Unlock()
		}(q)
	}
	fmt.Println(cnt);
	go func() {
		wg.Wait()
		close(errCh)
	}()
	if err := <-errCh; err != nil {
		return nil, err
	}
	return problems, nil
}
func floatSliceToVector(slice []float32) string {
    strSlice := make([]string, len(slice))
    for i, v := range slice {
        strSlice[i] = fmt.Sprintf("%f", v)
    }
    return "[" + strings.Join(strSlice, ",") + "]"
}
func (s *Pgstore) AddEmbeddings(texts []string) error {
	if len(texts) == 0 {
		return fmt.Errorf("no texts to embed")
	}
	ctx := context.Background()
	_, err := s.DB.Exec("TRUNCATE TABLE embeddings")
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(texts))

	for i, text := range texts {
		if(i%1000==0){
			time.Sleep(time.Minute+20*time.Second);
		}
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			res, err := em.EmbedContent(ctx, genai.Text(text))
			if err != nil {
				errChan <- fmt.Errorf("embedding error: %w", err)
				return
			}
			val := floatSliceToVector(res.Embedding.Values)
			_, err = s.DB.Exec("INSERT INTO embeddings (content, embedding) VALUES ($1, $2)", text, val)
			if err != nil {
				errChan <- fmt.Errorf("insert error: %w", err)
			}
		}(text)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err // Return the first error encountered
		}
	}

	println("Completed")
	return nil
}
