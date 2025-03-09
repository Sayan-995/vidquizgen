
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
	"strings"

	"github.com/Sayan-995/vidquizgen/store"
	"github.com/Sayan-995/vidquizgen/yt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Question struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type QuizResponse struct {
	Questions []Question `json:"questions"`
}
func main() {
	s, err := store.Create()
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	http.HandleFunc("/",hello)
	http.HandleFunc("/api/quiz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == "POST" {
			generateQuizHandler(w, r, s)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
	})
	
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Open http://localhost%s in your browser\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func generateQuizHandler(w http.ResponseWriter, r *http.Request, s *store.Pgstore) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid form data"})
		return
	}
	fmt.Println("came");
	videoURL := r.FormValue("url")
	if videoURL == "" {
		var reqBody struct {
			URL string `json:"url"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err == nil {
			videoURL = reqBody.URL
		}
	}
	
	if videoURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "YouTube URL is required"})
		return
	}
	fmt.Printf("Processing video: %s\n", videoURL)
	start := time.Now()
	baseURL, err := yt.GetBaseURL(videoURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Failed to process video URL: %v", err)})
		return
	}
	
	transcript, err := yt.GetTranscript(baseURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Failed to get transcript: %v", err)})
		return
	}
	summarizedTranscript, err := yt.SummerizeTranscript(transcript)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Failed to summarize transcript: %v", err)})
		return
	}
	problems, err := s.GetTopQuestions(summarizedTranscript)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Failed to generate questions: %v", err)})
		return
	}
	questions := formatQuestions(problems)
	elapsed := time.Since(start)
	fmt.Printf("Processed in %s\n", elapsed)
	json.NewEncoder(w).Encode(QuizResponse{Questions: questions})
}
func formatQuestions(problems []string) []Question {
	var questions []Question
	titleRegex := regexp.MustCompile(`Title:\s*(.+)`)
	
	for _, problem := range problems {
		match := titleRegex.FindStringSubmatch(problem)
		if len(match) > 1 {
			title := match[1]
			// Convert title to leetcode URL format
			urlSlug := toURLSlug(title)
			
			questions = append(questions, Question{
				Title: title,
				URL:   fmt.Sprintf("https://leetcode.com/problems/%s", urlSlug),
			})
		}
	}
	
	return questions
}

func toURLSlug(title string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s-]`)
	slug := re.ReplaceAllString(title, "")
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")
	
	return strings.ToLower(slug)
}
func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HELLOW"))
}