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
	// 获取get参数
	query := r.URL.Query()
	name := query.Get("name")
	fmt.Fprintf(w, "hello! %s", name)

	// 获取post参数
	if r.Method == "POST" {
		r.ParseForm()
		r.FormValue("name")
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "path = %q \n", r.URL.Path)


}
