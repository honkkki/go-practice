package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync/atomic"

	"github.com/gocolly/colly"
)

var (
	co          *colly.Collector
	downLoadNum int32
	cookie      string
)

func init() {
	// 通过配置文件获取cookie值
	file, _ := os.Open("./cookie.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	cookie, _ = reader.ReadString('\n')

	co = colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	co.Limits([]*colly.LimitRule{
		{
			//DomainGlob: "",
			Parallelism: 10,
		},
	})
}

func spiderWithCookie() {
	imgC := co.Clone()

	co.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", cookie)
	})

	co.OnHTML(".update", func(e *colly.HTMLElement) {
		// 找出包含update的class 且 class name = update magnifier
		if e.Attr("class") == "update magnifier" {
			rawUrl := e.ChildAttr("a", "href")
			//fmt.Println(rawUrl)
			imgC.Visit(rawUrl)
		}
	})

	// 绕过豆瓣防盗链机制
	imgC.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", "https://movie.douban.com/")
		fmt.Println("正在抓取：", r.URL)
	})

	imgC.OnResponse(func(response *colly.Response) {
		err := os.MkdirAll("./img/miyeon/", 0777)
		if err != nil {
			log.Fatal("make dir failed: ", err)
		}
		fileName := strings.Split(response.Request.URL.String(), "/")[7]
		f, _ := os.Create("./img/miyeon/" + fileName)
		defer f.Close()
		io.Copy(f, bytes.NewReader(response.Body))
		atomic.AddInt32(&downLoadNum, 1)
		fmt.Println("success!", downLoadNum)
		if downLoadNum == 500 {
			fmt.Println("下载完成！")
			os.Exit(1)
		}
	})

	// 执行下一页
	co.OnHTML(".photo-wp", func(e *colly.HTMLElement) {
		e.Request.Visit(e.ChildAttr("a", "href"))
	})

	co.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	co.Visit("https://movie.douban.com/celebrity/1419037/photo/2646824033/#photo")
	co.Wait()
	imgC.Wait()
}

func main() {
	//fmt.Println(cookie)
	//os.Exit(1)
	spiderWithCookie()
}
