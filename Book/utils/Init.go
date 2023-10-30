package utils

import (
	"Book/handler"
	"fmt"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	isErr := handler.Handler(err, "reading config error")
	if isErr {
		return
	} else {
		fmt.Println("init config success...")
	}
}

func InitMySQL() {
	var err error
	host := viper.GetString("mysql.host")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	db := viper.GetString("mysql.db")
	options := viper.Sub("mysql.options")
	charset := options.GetString("charset")
	parseTime := options.GetBool("parseTime")
	loc := options.GetString("loc")
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s", user, password, host, db, charset, parseTime, loc)
	DB, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{})
	isErr := handler.Handler(err, "open mysql error")
	if isErr {
		return
	}
	fmt.Println("init mysql success...")
}
