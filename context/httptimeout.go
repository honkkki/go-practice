package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const TRACEID = "trace_id" 

func HandleWorker(w http.ResponseWriter, r *http.Request)  {
	ctx := context.WithValue(context.Background(), TRACEID, rand.Int31())
	a(ctx)
	if r.URL.Path != "/favicon.ico" {
		randNum := rand.Intn(2)
		if randNum == 0 {
			time.Sleep(5 * time.Second)
			fmt.Fprintln(w, "slow response")
			fmt.Println("slow response")
			return
		}

		fmt.Fprintln(w, "fast response")
		fmt.Println("fast response")
		return
	} else {
		return
	}
}

func a(ctx context.Context)  {
	id := ctx.Value(TRACEID)
	fmt.Printf("id = %v function: a\n", id)
	b(ctx)
}

func b(ctx context.Context)  {
	id := ctx.Value(TRACEID)
	fmt.Printf("id = %v function: b\n", id)
}

func main()  {
	http.HandleFunc("/", HandleWorker)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}
