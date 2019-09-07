package handler

import (
	"context"

	"github.com/micro/go-log"

	envTest "github.com/ghosv/env-test/proto/envtest"
)

type Test struct{}

func (e *Test) Env(ctx context.Context, req *envTest.Request, rsp *envTest.Response) error {
	log.Log("Received Test.Env request")
	rsp.Env = map[string]string{
		"ID": req.Id,
	}
	return nil
}
