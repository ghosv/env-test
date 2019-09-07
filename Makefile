pro_name:=ghostv
srv_name:=env-test
#GOPATH:=$(shell go env GOPATH)
GOPATH:=${shell echo $$ORIGIN_GOPATH}
consul_ip:=consul.dev.ghost.gh
consul_port:=80

docker_name:=${pro_name}/${srv_name}

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/entry-point main.go plugin.go

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/**/*.proto

.PHONY: dev
dev:
	go run . --registry consul --registry_address ${consul_ip}:${consul_port}

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ${docker_name}:latest
