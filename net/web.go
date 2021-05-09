package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	// 配置路由与处理逻辑
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello golang http")
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "path = %q \n", r.URL.Path)


}
