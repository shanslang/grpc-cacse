package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc-case/client/services"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if err != nil {
		log.Fatal("LoadX509KeyPair: ", err)
	}
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		// InsecureSkipVerify: true,                    // 出现错误rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate signed by unknown authority"  暂时设置为true来取消对服务端证书的校验
		Certificates: []tls.Certificate{cert}, // 客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds)) // 连接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	productClient := services.NewProductServiceClient(conn)
	productRes, err := productClient.GetProductStock(context.Background(),
		&services.ProductRequest{ProductId: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(productRes.ProductStock)
}
