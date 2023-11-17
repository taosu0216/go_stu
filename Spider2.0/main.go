package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	DB    *gorm.DB
	count int
)

type Movie struct {
	gorm.Model
	Title  string `json:"title" gorm:"type:varchar(255);not null;"`
	Img    string `json:"img" gorm:"type:varchar(256);not null;"`
	Rank   string `json:"rank" gorm:"type:varchar(256);not null;"`
	Desc   string `json:"desc" gorm:"type:varchar(256);not null;"`
	Tags   string `json:"tags"`
	Author string `json:"author"`
	Actor  string `json:"actor"`
	Time   string `json:"time"`
}

func main() {
	InitDB()
	start := time.Now()
	for i := 0; i < 10; i++ {
		num := fmt.Sprintf("%d", i*25)
		Spider(num)
	}
	end := time.Since(start)
	st := time.Now()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		num := fmt.Sprintf("%d", i*25)
		go Spider2(num, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	ed := time.Since(st)

	fmt.Println("正常爬虫耗时：", end)  // 3.8124319s
	fmt.Println("go协程爬虫耗时：", ed) // 468.1153ms
	fmt.Println("爬取速率比: ", end/ed)
}
func Err(err error, str string) {
	if err != nil {
		log.Fatalln(str, "is: ", err)
	}
}
func Spider(page string) {
	client := http.Client{}
	url := `https://movie.douban.com/top250?start=`
	req, err := http.NewRequest("GET", url+page, nil)
	Err(err, "创建连接失败")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	Err(err, "请求连接失败")
	//解析网页
	docs, err := goquery.NewDocumentFromReader(resp.Body)

	defer resp.Body.Close()

	Err(err, "解析网页失败")

	//获取节点信息
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	docs.Find("#content > div > div.article > ol > li").
		Each(func(i int, s *goquery.Selection) {
			title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			imgTag := s.Find("div > div.pic > a > img")
			img, ok := imgTag.Attr("src")
			info := s.Find("div > div.info > div.bd > p:nth-child(1)").Text()
			rank := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
			desc := s.Find("div > div.info > div.bd > p.quote > span").Text()
			if ok {
				count++
				author, actor, time, tags := InfoSplit(info)
				if title == "" {
					title = "none"
				}
				if img == "" {
					img = "none"
				}
				if author == "" {
					author = "none"
				}
				if actor == "" {
					actor = "none"
				}
				if time == "" {
					time = "none"
				}
				if tags == "" {
					tags = "none"
				}
				if rank == "" {
					rank = "none"
				}
				if desc == "" {
					desc = "none"
				}
				data := Movie{
					Title:  title,
					Img:    img,
					Author: author,
					Actor:  actor,
					Time:   time,
					Tags:   tags,
					Rank:   rank,
					Desc:   desc,
				}
				InsertDB(&data)
				fmt.Println(data)
			}
		})
}
func Spider2(page string, ch chan bool) {
	client := http.Client{}
	url := `https://movie.douban.com/top250?start=`
	req, err := http.NewRequest("GET", url+page, nil)
	Err(err, "创建连接失败")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	Err(err, "请求连接失败")
	//解析网页
	docs, err := goquery.NewDocumentFromReader(resp.Body)

	defer resp.Body.Close()

	Err(err, "解析网页失败")

	//获取节点信息
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	docs.Find("#content > div > div.article > ol > li").
		Each(func(i int, s *goquery.Selection) {
			title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			imgTag := s.Find("div > div.pic > a > img")
			img, ok := imgTag.Attr("src")
			info := s.Find("div > div.info > div.bd > p:nth-child(1)").Text()
			rank := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
			desc := s.Find("div > div.info > div.bd > p.quote > span").Text()
			if ok {
				count++
				author, actor, time, tags := InfoSplit(info)
				if title == "" {
					title = "none"
				}
				if img == "" {
					img = "none"
				}
				if author == "" {
					author = "none"
				}
				if actor == "" {
					actor = "none"
				}
				if time == "" {
					time = "none"
				}
				if tags == "" {
					tags = "none"
				}
				if rank == "" {
					rank = "none"
				}
				if desc == "" {
					desc = "none"
				}
				data := Movie{
					Title:  title,
					Img:    img,
					Author: author,
					Actor:  actor,
					Time:   time,
					Tags:   tags,
					Rank:   rank,
					Desc:   desc,
				}
				InsertDB(&data)
				fmt.Println(data)
				if ch != nil {
					ch <- true
				}
			}
		})
}
func InfoSplit(info string) (author, actor, time, tags string) {
	//电影导演
	authorRe, err := regexp.Compile(`导演:.*   `)
	Err(err, "电影导演错误")
	author = string(authorRe.Find([]byte(info)))
	author = strings.TrimPrefix(author, "导演:")
	author = strings.TrimSpace(author)

	//电影演员
	actorRe, err := regexp.Compile(`主演:(.*)`)
	Err(err, "电影演员错误")
	actor = string(actorRe.Find([]byte(info)))
	actor = strings.TrimPrefix(actor, "主演:")
	parts := strings.Split(actor, "/")
	actor = parts[0]
	actor = strings.TrimSpace(actor)

	//电影时间
	timeRe, err := regexp.Compile(`(\d+)`)
	Err(err, "电影时间错误")
	time = string(timeRe.Find([]byte(info)))
	time = strings.TrimSpace(time)

	//电影标签
	tagsRe, err := regexp.Compile(`/([^\/]+)$`)
	Err(err, "电影标签错误")
	tags = string(tagsRe.Find([]byte(info)))
	tags = strings.TrimSpace(tags)

	return
}
func InitDB() {
	path := "root:root@tcp(127.0.0.1:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(path), &gorm.Config{})
	Err(err, "数据库连接失败")
	DB.AutoMigrate(&Movie{})
	fmt.Println("数据库连接成功")
}
func InsertDB(data *Movie) {
	result := DB.Create(data)
	Err(result.Error, "insert failed")
}

func (table *Movie) TableName() string {
	return "movies"
}
