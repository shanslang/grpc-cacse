package main

import (
	"context"
	"grpc-case/server/helper"
	"grpc-case/server/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func httpServices() {
	gwmux := runtime.NewServeMux()
	gRpcEndPoint := "localhost:8081"
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProductServiceHandlerFromEndpoint(context.Background(), gwmux, gRpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	// 订单
	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, gRpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("httpServices err: ", err)
	}
}
