package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	Err(err, "配置文件读取失败")
	fmt.Println("配置文件读取成功...")
}
func InitMysql() {
	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Err(err, "mysql连接失败...")
	fmt.Println("mysql连接成功...")
}
func Err(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
