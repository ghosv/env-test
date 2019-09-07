package main

import (
	"fmt"
	"os"

	"github.com/ghosv/env-test/handler"
	envTest "github.com/ghosv/env-test/proto/envtest"
	"github.com/ghosv/env-test/subscriber"
	"github.com/kr/pretty"
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"

	"github.com/micro/go-micro/config"
	consulConfig "github.com/micro/go-micro/config/source/consul"
)

const proName = "ghostv"
const srvName = "env-test"
const fullName = proName + ".srv." + srvName

func main() {
	consulAddr := os.Getenv("MICRO_REGISTRY_ADDRESS")
	source := consulConfig.NewSource(
		consulConfig.WithAddress(consulAddr),
		consulConfig.WithPrefix("/demo"),
		consulConfig.StripPrefix(true),
	)
	conf := config.NewConfig()
	if err := conf.Load(source); err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(conf.Get("test"))

	// New Service
	service := grpc.NewService(
		micro.Name(fullName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	envTest.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(fullName, service.Server(), new(subscriber.Test))
	// Register Function as Subscriber
	micro.RegisterSubscriber(fullName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
