package main

import (
	"fmt"
	"grpc-case/server/services"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("keys/example.com.cert", "keys/example.com.key")
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProductServiceServer(rpcServer, new(services.ProductService))

	// 方式一：tcp
	// lis, _ := net.Listen("tcp", ":8081")
	// rpcServer.Serve(lis)

	// 方式二：Http
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Proto: ", r.Proto) // 使用协议
		fmt.Println("Header:", r.Header)
		fmt.Println("request:", r)
		rpcServer.ServeHTTP(w, r)
	})
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	// httpServer.ListenAndServe() // 无证书方式 // http://localhost:8081/
	httpServer.ListenAndServeTLS("keys/example.com.cert", "keys/example.com.key") // 证书验证 // https://localhost:8081/
}
