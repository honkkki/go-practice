package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	c *colly.Collector
)

func init() {
	c = colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.Limits([]*colly.LimitRule{
		{
			//DomainGlob: "",
			Parallelism: 5,
		},
	})
}

func downloadImg(val, movieName string) {
	err := os.MkdirAll("./img", 0777)
	if err != nil {
		log.Fatal("make dir failed: ", err)
	}

	// save images as files
	resp, _ := http.Get(val)
	//savePath := "./img/" + strings.Split(val, "/")[7]
	savePath := "./img/" + movieName + ".jpg"
	file, _ := os.Create(savePath)
	buf := bufio.NewWriter(file)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal("download img failed: ", err)
	}

	resp.Body.Close()
	buf.Flush()
}

func spiderDb() {
	c.OnHTML(".item", func(e *colly.HTMLElement) {
		// get movie name
		movieName := e.DOM.Find(".info .hd span").Eq(0).Text()

		// download images
		val, bool := e.DOM.Find(".pic img").Attr("src")
		if bool == true {
			downloadImg(val, movieName)
		}
	})

	// 执行下一页
	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("正在抓取：", r.URL)
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}

func main() {
	spiderDb()
}
