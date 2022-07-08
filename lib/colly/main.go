package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// 抓取豆瓣top250电影列表
func main() {
	c := colly.NewCollector(
		// 按排序爬虫所以关闭异步爬虫方式
		//colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.Limits([]*colly.LimitRule{
		{
			DomainGlob:  "*.douban.*",
			Parallelism: 5,
		},
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		fmt.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))

		file, _ := os.OpenFile("./douban.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		defer file.Close()
		str := fmt.Sprintf("id: %s    name: %s\n", strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))

		file.WriteString(str)
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}
