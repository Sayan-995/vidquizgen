package store

import (
	"context"
	"database/sql"
	"fmt"

	// "net/http"
	"os"
	"time"

	// "github.com/PuerkitoBio/goquery"
	pb "github.com/Sayan-995/vidquizgen/bindings"
	t "github.com/Sayan-995/vidquizgen/types"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Pgstore struct{
	DB *sql.DB
}
func Create()(*Pgstore,error){
	db,err:=GetDB()
	if(err!=nil){
		return nil,err
	}
	return &Pgstore{
		DB: db,
	},nil
}
func GetDB()(*sql.DB,error) {
	godotenv.Load()
	db, err := sql.Open("postgres",os.Getenv("DATABASE_URI"))
	if(err!=nil){
		return nil,err;
	}
	err=db.Ping()
	if(err!=nil){
		return nil,err;
	}
	return db,nil;
}
func ScanIntoQuestion(row *sql.Rows)(t.Question,error){
	var q t.Question
	err:=row.Scan(&q)
	return q,err;
}
func ScrapProblemDetails(TitleSlug string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	conn,err:=grpc.NewClient("localhost:50051",grpc.WithInsecure())
	if(err!=nil){
		return "",err
	}
	defer conn.Close()
	client:=pb.NewStatementServiceClient(conn)
	resp,err:=client.GetStatement(ctx,&pb.ProblemRequest{
		TitleSlug: TitleSlug,
	})
	if(err !=nil){
		return "",err
	}
	fmt.Println(resp.Statement);
	return resp.Statement,nil
}



// func (s *Pgstore)GetText()([]string,error){
// 	var text []string
// 	query:=`SELECT * FROM QUESTIONS`
// 	res,err:=s.DB.Query(query)
// 	if(err!=nil){
// 		return nil,err
// 	}
// 	for(res.Next()){
// 		q,err:=ScanIntoQuestion(res)
// 		if(err!=nil){
// 			return nil,err
// 		}

// 	}
// }