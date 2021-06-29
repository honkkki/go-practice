package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handleWorker(w http.ResponseWriter, r *http.Request)  {
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


func main()  {
	http.HandleFunc("/", handleWorker)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}
