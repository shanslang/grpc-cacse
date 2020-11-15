package main

import (
	"grpc-case/server/helper"
	"grpc-case/server/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	creds := helper.GetServerCreds()

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProductServiceServer(rpcServer, new(services.ProductService)) // 商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrderService))     // 订单服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))       // 用户服务

	// 方式一：tcp
	lis, _ := net.Listen("tcp", ":8081")
	go rpcServer.Serve(lis)

	// 方式二：Http
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Proto: ", r.Proto) // 使用协议
	// 	fmt.Println("Header:", r.Header)
	// 	fmt.Println("request:", r)
	// 	rpcServer.ServeHTTP(w, r)
	// })
	// httpServer := &http.Server{
	// 	Addr:    ":8081",
	// 	Handler: mux,
	// }
	// httpServer.ListenAndServeTLS("cert/server.crt", "cert/server.key")

	// http接口
	go httpServices()
	select {}
}
