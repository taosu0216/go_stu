package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 这里DB必须要大写,这样好像是全局变量,在包外也可以使用
var (
	CTX = context.Background()
	DB  *gorm.DB
	Red *redis.Client
)

const (
	PublishKey = "websocket"
)

func InitConfig() {
	//设置搜索路径,即在config目录下查找app文件(这里默认是当前路径,这里准确来说应该是./config)
	viper.AddConfigPath("config")
	//这里是获取文件名(因为什么配置文件都能获取,所以这里只有文件名,不加后缀,扩展名)
	viper.SetConfigName("app")

	//判断是否能够成功读取文件内容
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("config app inited ... ")
	//这里不打印也可以,只是测试用
	// fmt.Println("config app info : ", viper.Get("app"))
	// fmt.Println("config mysql info : ", viper.Get("mysql"))
}
func InitMySQL() {
	newLoger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	//dns := viper.GetString("mysql.dns")
	username := viper.GetString("mysql.username")
	passwd := viper.GetString("mysql.passwd")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	db := viper.GetString("mysql.db")
	////获取配置中的子配置项
	options := viper.Sub("mysql.options")
	charset := options.GetString("charset")
	//这里是getbool,因为这里定义的是bool值
	parseTime := options.GetBool("parseTime")
	loc := options.GetString("loc")
	//拼接数据库连接语句
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", username, passwd, host, port, db, charset, parseTime, loc)
	//连接语句,完成初始化
	DB, _ = gorm.Open(mysql.Open(databaseURL), &gorm.Config{Logger: newLoger})

	fmt.Println("mysql inited ... ")
}

func InitRedis() {
	redis_data := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host"),
		Password:     viper.GetString("redis.passwd"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := redis_data.Ping(CTX).Result()
	if err != nil {
		fmt.Println("init redis err : ", err)
	} else {
		fmt.Println("Redis inited... ", pong)
	}

}

// Publish
// 这里的channel不是go特有的那个管道的意思,而是要订阅/发布消息的频道
// 所以这里的channel是string类型,这里就是要发布消息的频道的名字
// 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = Red.Publish(ctx, channel, msg).Err()
	fmt.Println("Publish : ", msg)
	return err
}

//Subscribe
/*
区别在于订阅的方式和范围。

Subscribe方法用于订阅一个或多个指定的频道。当有消息发布到被订阅的频道时，客户端会接收到这些消息。
这种订阅方式只能订阅普通的频道，不能订阅模式。

PSubscribe方法则用于订阅一个或多个指定的模式。当有消息匹配被订阅的模式时，客户端会接收到这些消息。
模式订阅可以使用通配符来匹配多个频道，例如使用*匹配一个单词，使用?匹配一个字符。
这种订阅方式可以订阅模式，但不能订阅普通的频道。
*/
// 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe : ", msg.Payload)
	return msg.Payload, err
}
