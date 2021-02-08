package main

import (
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

type ProdModel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)

	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prodName"+strconv.Itoa(100+i)))
	}

	return ret
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	ginRouter := gin.Default()

	// 路由group
	/*v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("GET", "/prod", func(context *gin.Context) {
			//context.String(200, "gin user api")
			context.JSON(200, NewProdList(5))
		})

	}*/

	ginRouter.Handle("GET", "/prod", func(context *gin.Context) {
		//context.String(200, "gin user api")
		//context.JSON(200, NewProdList(5))

		//context.JSON(200,  )

		context.JSON(200, gin.H{
			"status": 200,
			"data1":  NewProdList(3),
			"data2":  NewProdList(2),
		})
	})

	//server := web.NewService(web.Address(":8001"))
	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	//server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, _ = writer.Write([]byte("Hello world"))
	//})

	_ = server.Run()
}
