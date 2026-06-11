package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}

// https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
//https://riptutorial.com/go/topic/1422/http-client
//https://chenyitian.gitbooks.io/gin-tutorials/content/gin/6.html
//https://zetcode.com/all/#go
//https://golangbyexample.com/
//https://freshman.tech/web-development-with-go/#formatting-the-published-date
//https://livebook.manning.com/book/go-web-programming/chapter-1/50
//https://pkg.go.dev/github.com/gin-gonic/gin@v1.7.7#section-readme
