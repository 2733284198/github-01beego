package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-micro/web"
	//"github.com/micro/go-plugins/registry/consul/v2"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/go-micro/web"
)

func main() {
	microRun()
	//ginRun()
}

func ginRun() {
	//default := gin.

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8002")
}

func microRun() {

	/*
	 consul
	//*/
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	ginRouter := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	// 路由
	ginRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World gin rpc!")
	})

	ginRouter.GET("/product", func(c *gin.Context) {
		//c.String(http.StatusOK, "hello World gin rpc!")
		c.JSON(200, gin.H{
			"status": 200,
			"data":   "product1",
		})
	})

	// 路由group
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("GET", "/index", func(context *gin.Context) {
			context.String(200, "v1/index")
		})

		v1Group.Handle("GET", "/users", func(context *gin.Context) {
			context.String(200, "gin user api")
		})

		v1Group.GET("/news", func(c *gin.Context) {
			//c.String(http.StatusOK, "hello World gin rpc!")
			c.JSON(200, gin.H{
				"status": 200,
				"data":   "news",
			})
		})
	}

	//server := web.NewService(web.Address(":8001"))
	server := web.NewService(
		web.Name("com.manlan.consul.reg"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	//server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, _ = writer.Write([]byte("Hello world"))
	//})

	_ = server.Run()
}
