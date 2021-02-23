package main

import (
	"github.com/WiFeng/go-sky"
	"github.com/WiFeng/go-sky-example/pkg/endpoint"
	"github.com/WiFeng/go-sky-example/pkg/service"
	"github.com/WiFeng/go-sky-example/pkg/transport/http"
)

func main() {

	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = http.NewHandler(endpoints)
	)

	sky.Run(httpHandler)
}
