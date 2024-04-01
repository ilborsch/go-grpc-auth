package main

import (
	"fmt"
	"go-grpc-auth/internal/config"
)

func main() {
	// init config object
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger object

	// TODO: init app

	// TODO: start gRPC server

}
