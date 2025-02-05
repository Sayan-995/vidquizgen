package main

import (
	"fmt"

	"github.com/Sayan-995/vidquizgen/store"
)

func main(){
	_,err:=store.Create()
	if(err!=nil){
		fmt.Println(err)
	}
	_,err =store.ScrapProblemDetails("two-sum");
	if(err!=nil){
		fmt.Println(err)
	}
}