package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
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
			Parallelism: 6,
		},
	})
}

func downloadImg(imgUrl, fileName, dirName string) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Fatal("make dir failed: ", err)
	}

	// save images as files
	resp, _ := http.Get(imgUrl)
	//savePath := "./img/" + strings.Split(val, "/")[7]
	if !strings.HasSuffix(fileName, ".jpg") {
		fileName = fileName + ".jpg"
	}
	savePath := dirName + "/" + fileName
	file, _ := os.Create(savePath)
	defer file.Close()
	buf := bufio.NewWriter(file)
	_, err = io.Copy(file, resp.Body)
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
			downloadImg(val, movieName, "./img/")
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

func spiderLhc() {
	imgC := c.Clone()
	c.OnHTML(".photo-wp", func(e *colly.HTMLElement) {
		imgUrl := e.ChildAttr("a img", "src")
		fmt.Println(imgUrl)
		//fileName := strings.Split(imgUrl, "/")[7]
		//downloadImg(imgUrl, fileName, "./img/lhc")
		imgC.Visit(imgUrl)
	})

	// 执行下一页
	c.OnHTML(".photo-wp", func(e *colly.HTMLElement) {
		e.Request.Visit(e.ChildAttr("a", "href"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("正在抓取：", r.URL)
	})

	imgC.OnResponse(func(response *colly.Response) {
		err := os.MkdirAll("./img/lhc", 0777)
		if err != nil {
			log.Fatal("make dir failed: ", err)
		}
		f, _ := os.Create("./img/lhc/" + strconv.Itoa(int(rand.Int31())) + ".jpg")
		defer f.Close()
		io.Copy(f, bytes.NewReader(response.Body))
	})

	c.Visit("https://movie.douban.com/celebrity/1410267/photo/2548187252/")
	c.Wait()
	imgC.Wait()
}

func main() {
	//spiderDb()
	spiderLhc()
}
