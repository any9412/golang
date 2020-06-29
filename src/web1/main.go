package main

import (
	"net/http"
	"github.com/any9412/golang/src/web1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler()) // web server 실행 & request를 3000 port에서 기다림.
}