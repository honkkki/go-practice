package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

var (
	c *colly.Collector
	downLoadNum int32
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

func spiderWithCookie() {
	imgC := c.Clone()

	c.OnHTML(".update", func(e *colly.HTMLElement) {
		// 找出包含update的class 且 class name = update magnifier
		if e.Attr("class") == "update magnifier" {
			rawUrl := e.ChildAttr("a", "href")
			imgC.Visit(rawUrl)
		}
	})

	// 执行下一页
	c.OnHTML(".photo-wp", func(e *colly.HTMLElement) {
		e.Request.Visit(e.ChildAttr("a", "href"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "douban-fav-remind=1; _ga=GA1.2.1015094152.1570013182; gr_user_id=4f75482a-0275-4bd2-b398-a5028da6f213; _vwo_uuid_v2=D341AF5900A68B0DA7A3E2D5E99CF7481|bb77b52d8102568b896dfa42543d6cb8; __gads=ID=92200105b4f0065e:T=1573922086:S=ALNI_MZcxF7PXjOyYb_rp8d0bSqCEIZcsQ; __utmv=30149280.11974; douban-profile-remind=1; bid=vq6g5Qa2HBw; ll=\"108309\"; viewed=\"27044219_5414086_3843879_4204727_10727672_3844259_30178432_27091993_2361462_10760397\"; _vwo_uuid_v2=D341AF5900A68B0DA7A3E2D5E99CF7481|bb77b52d8102568b896dfa42543d6cb8; push_noty_num=0; push_doumail_num=0; __utmc=30149280; __utmc=223695111; __utmz=223695111.1621013791.85.80.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmz=30149280.1621151053.161.153.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; ap_v=0,6.0; ct=y; dbcl2=\"119749516:28p9I+gw/38\"; ck=qG4w; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1621694762%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DuPCVEii2s7VK9r5uBLa93CyKfPEVE79f6Iy__i1VXpLPZuu1T3HVKVVNv1rvj7v0Ofm5NWBSVQPMiM03dAxRuK%26wd%3D%26eqid%3D8d9484c6000ce0f200000006609eb51a%22%5D; _pk_ses.100001.4cf6=*; __utma=30149280.1015094152.1570013182.1621689389.1621694762.168; __utmb=30149280.0.10.1621694762; __utma=223695111.1346728745.1570974096.1621689389.1621694762.92; __utmb=223695111.0.10.1621694762; _pk_id.100001.4cf6=e0c26d1c39ca78df.1570974095.90.1621695151.1621692454.")

	})

	// 绕过豆瓣防盗链机制
	imgC.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", "https://movie.douban.com/")
		fmt.Println("正在抓取：", r.URL)
	})

	imgC.OnResponse(func(response *colly.Response) {
		err := os.MkdirAll("./img/irene/", 0777)
		if err != nil {
			log.Fatal("make dir failed: ", err)
		}
		fileName := strings.Split(response.Request.URL.String(), "/")[7]
		f, _ := os.Create("./img/irene/" + fileName)
		defer f.Close()
		io.Copy(f, bytes.NewReader(response.Body))
		atomic.AddInt32(&downLoadNum, 1)
		fmt.Println("下载成功!", downLoadNum)
		if downLoadNum == 673 {
			fmt.Println("下载完成！")
			os.Exit(1)
		}
	})

	c.Visit("https://movie.douban.com/celebrity/1360542/photo/2643155707/#photo")
	c.Wait()
	imgC.Wait()
}

func main() {
	spiderWithCookie()
}
