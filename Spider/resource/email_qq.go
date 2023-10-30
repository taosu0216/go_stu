package resource

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func Getemail(file *os.File) {
	//计算爬取时间
	start := time.Now()
	//对网站进行访问及爬取
	resp, err := http.Get("https://www.xyfinance.org/hot/516352")
	HandleErr(err, "http.Get url")
	defer resp.Body.Close()
	//读取网站的全部响应内容
	pagebody, err := io.ReadAll(resp.Body)
	HandleErr(err, "io.ReadAll")
	//pagebody本来是字节切片,这里转换成字符串格式方便操作
	pageStr := string(pagebody)
	/*
		这行代码首先使用 regexp.MustCompile 函数来将正则表达式模式 reQQemail 编译为一个可重复使用的正则表达式对象 re。
		这将创建一个用于匹配 reQQemail 模式的正则表达式。
	*/
	re := regexp.MustCompile(reQQemail)
	/*
		使用之前创建的正则表达式 re 在字符串 pageStr 中查找所有匹配 reQQemail 模式的子字符串，并将结果存储在 results 变量中。
		-1 表示查找所有匹配项。如果是正整数则表示匹配次数,而-1则是有多少匹配多少
	*/
	results := re.FindAllStringSubmatch(pageStr, -1)
	//这里的results是一个二维切片,result是一个一维切片,reslut[0]相当于results[0][0]
	for index, result := range results {
		email := result[0]
		qq := result[1]
		fmt.Println("email:", email)
		fmt.Println("QQ:", qq)
		fmt.Fprintf(file, "%d. \nemail:%s\n", index+1, email)
		fmt.Fprintf(file, "QQ:%s\n", qq)
	}
	end := time.Now()
	fmt.Println("线程用时:", end.Sub(start))
}
