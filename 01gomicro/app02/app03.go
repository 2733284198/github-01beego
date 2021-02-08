package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/go-plugins/registry/consul/v2"
	//_ "github.com/micro/go-plugins/transport/rabbitmq"
	// kafka broker
	//_ "github.com/micro/go-plugins/broker/kafka"
	//_ "github.com/micro/go-plugins/registry/etcdv3"
	//_ "github.com/micro/go-plugins/registry/kubernetes"
	//_ "github.com/micro/go-plugins/transport/nats"
	//_ "github.com/micro/cli"
	//_ "github.com/micro/micro/plugin"
)

type Greeter struct{}

//func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
//	rsp.Greeting = "Hello " + req.Name
//	return nil
//}

func main() {
	//consulconfig := consul.Config()

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.01:8500",
		}
	})

	service := micro.NewService(
		//micro.Name("greeter"),
		micro.Name("my.consule"),
		micro.Version("lastest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	//pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
