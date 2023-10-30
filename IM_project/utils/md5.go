package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	//这里的Write是一个类似写入缓冲区的操作
	h.Write([]byte(data))
	//Sum是将缓冲区的内容提取并计算,这里的Sum()需要的传参是一个字节数组,就是将新Sum的值追加到原来的数组里
	//这里直接填nil,自动新建一个就可以
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密,需要的分别是用户输入的密码,数据库中的盐,数据库中的加密密码
func ValidPassword(plainpwd, salt string, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}

func MakeToken() string {
	tmp := fmt.Sprintf("%d", time.Now().Unix())
	return MD5Encode(tmp)
}
