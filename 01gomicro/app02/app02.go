package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/labstack/echo/v4"
	//"github.com/douyu/jupiter/pkg/govern"
)

var app jupiter.Application

func main() {

	_ = app.Startup()
	_ = app.Serve(startHTTPServer())
	//app.Serve(startGRPCServer())
	//app.Schedule(startWorker())
	//
	//app.SetGovernor("127.0.0.1:801")
	//app.SetGovernor("127.0.0.1:9990") // 默认为127.0.0.1:9990

	_ = app.Run()
}

func startHTTPServer() server.Server {

	server1 := xecho.DefaultConfig().Build()

	server1.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, "Gopher Wuhan")
	})

	return server1
}
