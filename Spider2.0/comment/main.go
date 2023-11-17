package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Err(err error, str string) {
	if err != nil {
		log.Fatalln(str, "is: ", err)
	}
}

type Comment struct {
	gorm.Model
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	InitDB()
	urls := []string{
		"https://api.bilibili.com/x/v2/reply/wbi/main?oid=40559171&type=1&mode=3&pagination_str=%7B%22offset%22:%22%7B%5C%22type%5C%22:1,%5C%22direction%5C%22:1,%5C%22session_id%5C%22:%5C%221740848790700856%5C%22,%5C%22data%5C%22:%7B%7D%7D%22%7D&plat=1&web_location=1315875&w_rid=7e2b3f24bb98af6f95615b6497da021c&wts=1700047667",
		"https://api.bilibili.com/x/v2/reply/wbi/main?oid=40559171&type=1&mode=3&pagination_str=%7B%22offset%22:%22%7B%5C%22type%5C%22:1,%5C%22direction%5C%22:1,%5C%22session_id%5C%22:%5C%221740848790700856%5C%22,%5C%22data%5C%22:%7B%7D%7D%22%7D&plat=1&web_location=1315875&w_rid=359566361b21a402fb41e3d8ee81c872&wts=1700047670",
	}
	for index, url := range urls {
		Spider(url, index+1)
	}
}
func Spider(url string, index int) {
	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	Err(err, "创建连接失败")

	req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Microsoft Edge\";v=\"120\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	//req.Header.Set("Sec-Fetch-Dest", "document")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7")
	req.Header.Set("Cookie", "**********************************")

	resp, err := client.Do(req)
	Err(err, "发起连接失败")
	defer resp.Body.Close()
	fmt.Println("正在爬取第", index, "页的评论")
	docBytes, err := io.ReadAll(resp.Body)
	Err(err, "读取网页失败")
	docs := string(docBytes)
	messages, messagesLen := messageSplit(docs)
	names, namesLen := nameSplit(docs)
	if messagesLen <= namesLen {
		for i := 0; i < messagesLen; i++ {
			name := strings.TrimPrefix(names[i], "\"uname\":")
			message := strings.TrimPrefix(messages[i], "\"message\":")
			comment := Comment{
				Name:    name,
				Message: message,
			}
			DB.Create(&comment)
		}
	} else {
		for i := 0; i < namesLen; i++ {
			name := strings.TrimPrefix(names[i], "\"uname\":")
			message := strings.TrimPrefix(messages[i], "\"message\":")
			comment := Comment{
				Name:    name,
				Message: message,
			}
			DB.Create(&comment)
		}
	}

}
func messageSplit(docs string) ([]string, int) {
	count := 0
	docRe, err := regexp.Compile(`"message":".*?"`)
	matches := docRe.FindAllString(docs, -1)
	Err(err, "编译正则表达式失败")
	for _ = range matches {
		count++
	}
	return matches, count
}
func InitDB() {
	path := "root:root@tcp(127.0.0.1:3306)/HuiYe?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(path), &gorm.Config{})
	Err(err, "数据库连接失败")
	DB.AutoMigrate(&Comment{})
	fmt.Println("数据库连接成功")
}
func nameSplit(docs string) ([]string, int) {
	count := 0
	docRe, err := regexp.Compile(`"uname":".*?"`)
	matches := docRe.FindAllString(docs, -1)
	Err(err, "编译正则表达式失败")
	for _ = range matches {
		count++
	}
	return matches, count
}
