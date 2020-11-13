package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc-case/server/services"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal("LoadX509KeyPair: ", err)
	}
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, // 双向验证
		ClientCAs:    certPool,
	})

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
	httpServer.ListenAndServeTLS("cert/server.crt", "cert/server.key")
}
