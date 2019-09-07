package subscriber

import (
	"context"

	envTest "github.com/ghosv/env-test/proto/envtest"
	"github.com/go-log/log"
)

type Test struct{}

func (a *Test) Handle(ctx context.Context, msg *envTest.Message) error {
	log.Log("Handler Received message: ", msg)
	return nil
}

func Handler(ctx context.Context, msg *envTest.Message) error {
	log.Log("Function Received message: ", msg)
	return nil
}
