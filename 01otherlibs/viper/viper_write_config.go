package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("configio")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./01otherlibs/viper") // todo:注意路径，直接运行可以，在命令行不行的

	viper.Set("yaml", "this is a example of yaml")

	viper.Set("redis.port", 4405)
	viper.Set("redis.host", "127.0.0.1")

	viper.Set("mysql.port", 3306)
	viper.Set("mysql.host", "192.168.1.0")
	viper.Set("mysql.username", "root123")
	viper.Set("mysql.password", "root123")

	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)
	}
}
