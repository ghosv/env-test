package main

import (
	"github.com/ghosv/env-test/handler"
	envTest "github.com/ghosv/env-test/proto/envtest"
	"github.com/ghosv/env-test/subscriber"
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

const proName = "ghostv"
const srvName = "env-test"

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name(srvName),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	envTest.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(proName+".srv."+srvName, service.Server(), new(subscriber.Test))
	// Register Function as Subscriber
	micro.RegisterSubscriber(proName+".srv."+srvName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
