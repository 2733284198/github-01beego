package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func tpflag() {

	pflag.Int("redis.port", 3302, "redis port")

	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() //根据上面配置加载文件
	if err != nil {
		fmt.Println(err)
		return
	}

	host := viper.Get("mysql.host")
	username := viper.GetString("mysql.username")
	port := viper.GetInt("mysql.port")
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetInt("redis.port")

	fmt.Println("mysql - host: ", host, ", username: ", username, ", port: ", port)
	fmt.Println("redis - host: ", redisHost, ", port: ", redisPort)

}

func tconfigtoml() {
	viper.SetConfigName("config")
	//viper.SetConfigName("config1")

	viper.SetConfigType("json")
	//viper.SetConfigType("toml")

	//viper.AddConfigPath("./01otherlibs")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if viper.IsSet("redis") {
		host := viper.GetString("redis")

		fmt.Println("redis:", host)
	} else {
		fmt.Println("key redis: no exist")
	}

	mysqlp := viper.GetString("mysql.host")
	fmt.Println("mysqlp", mysqlp)

	mysqlports := viper.GetIntSlice("mysql.host")
	fmt.Println("mysqlp", mysqlports)

	mysqlmapmetric := viper.GetStringMap("mysql.metric")
	fmt.Println("mysqlmapmetric:", mysqlmapmetric)

	mysqlmapmetric1 := viper.GetStringMap("mysql.metric")
	fmt.Println("mysqlmapmetric1:", mysqlmapmetric1["host"])
}

func main() {
	//tpflag()
	main2()
	//tconfigtoml()
}

func main2() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./01otherlibs/viper")
	err := viper.ReadInConfig() //根据上面配置加载文件
	if err != nil {
		fmt.Println(err)
		return
	}

	host := viper.Get("mysql.host")
	username := viper.GetString("mysql.username")
	port := viper.GetInt("mysql.port")
	portsSlice := viper.GetIntSlice("mysql.ports")

	metricPort := viper.GetInt("mysql.metric.port")
	redis := viper.Get("redis")

	mysqlMap := viper.GetStringMapString("mysql")

	if viper.IsSet("mysql.host") {
		fmt.Println("[IsSet()]mysql.host is set")
	} else {
		fmt.Println("[IsSet()]mysql.host is not set")
	}
	fmt.Println("mysql - host: ", host, ", username: ", username, ", port: ", port)
	fmt.Println("mysql ports :", portsSlice)
	fmt.Println("metric port: ", metricPort)
	fmt.Println("redis - ", redis)

	fmt.Println("mysqlmap - ", mysqlMap, ", username: ", mysqlMap["username"])
}
