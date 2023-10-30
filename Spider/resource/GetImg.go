package resource

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	ImgUrl chan string
	Task   chan string
)

// 类似main的启动程序
func GetImg() {
	//1.初始化channel
	ImgUrl = make(chan string, 100000)
	Task = make(chan string, 26)
	//2.爬虫携程
	for i := 1; i < 27; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://www.bizhizu.cn/shouji/tag-可爱/%d.html", i)
		go GetAllBody(url)
	}
	//3.检测是否完成爬取任务
	wg.Add(1)
	go Check()
	//4.下载图片
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Download()
	}
	wg.Wait()
}
func GetAllBody(url string) {
	//从整个网页的响应中获取图片
	urls := GetImgs(url)
	for _, url := range urls {
		ImgUrl <- url
	}
	Task <- url
	wg.Done()
}
func GetImgs(url string) (urls []string) {
	resp, err := http.Get(url)
	HandleErr(err, "http.Get url")
	defer resp.Body.Close()
	AllPage, err := io.ReadAll(resp.Body)
	HandleErr(err, "io.ReadAll")
	AllPageStr := string(AllPage)
	reImg := regexp.MustCompile(reimg)
	results := reImg.FindAllStringSubmatch(AllPageStr, -1)
	fmt.Println("共找到", len(results), "个结果")
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}
func Check() {
	var count int
	for {
		url := <-Task
		fmt.Printf("%s完成了任务\n", url)
		count++
		if count == 26 {
			close(ImgUrl)
			break
		}
	}
	wg.Done()
}
func Download() {
	for url := range ImgUrl {
		filename := Getfilename(url)
		ok := download(url, filename)
		if ok {
			fmt.Printf("%s下载成功\n", filename)
		} else {
			fmt.Printf("%s下载失败\n", filename)
		}
	}
}
func Getfilename(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	time := strconv.Itoa(int(time.Now().UnixNano()))
	filename = time + filename
	return
}
func download(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleErr(err, "http.Get url")
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	HandleErr(err, "io.ReadAll")
	filename = "tmp/" + filename
	err = os.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
