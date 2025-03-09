package yt

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	// "os"

	// "fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	t "github.com/Sayan-995/vidquizgen/types"
	st "github.com/Sayan-995/vidquizgen/store"
	"github.com/gocolly/colly"
	"github.com/google/generative-ai-go/genai"
	// "google.golang.org/api/option"
)

func GetBaseURL(url string)(string,error){
	c:=colly.NewCollector()
	var BaseURL string;
	c.OnHTML("script",func(h *colly.HTMLElement) {
		if(strings.Contains(h.Text,"captionTracks")){
			re:=regexp.MustCompile(`"captionTracks":\[(.*?)\]`)
			match:=re.FindStringSubmatch(h.Text);
			if(len(match)>1){
				BaseURL=match[1];
			}
		}
	})
	err:=c.Visit(url)
	re:=regexp.MustCompile(`"baseUrl":"([^"]+)"`);
	match:=re.FindStringSubmatch(BaseURL)
	BaseURL=match[1]
	// fmt.Println(BaseURL);
	if(err!=nil){
		return "",err;
	}
	return BaseURL,nil;
}

func GetTranscript(BaseURL string)(string, error){
	BaseURL=strings.ReplaceAll(BaseURL, `\u0026`, "&")
	resp,err:=http.Get(BaseURL)
	if err!=nil{
		return "",err;
	}
	defer resp.Body.Close()
	body,err:=io.ReadAll(resp.Body)
	if(err !=nil){
		return "",err;
	}
	var ts t.Transcript
	err=xml.Unmarshal(body,&ts)
	if(err !=nil){
		return "",err
	}
	var trans string
	for _,t:=range(ts.Text){
		trans+=t.Content;
	}
	trans=html.UnescapeString(trans)
	return trans,nil
}

func SummerizeTranscript(text string)(string,error){
	model:=st.Client.GenerativeModel("gemini-2.0-flash")
	resp,err:=model.GenerateContent(context.Background(),genai.Text(fmt.Sprintf(t.SummerizationPrompt,text)))
	if(err !=nil){
		return "",err
	}
	return fmt.Sprintf("%v\n", resp.Candidates[0].Content.Parts[0]),nil;
}