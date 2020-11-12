package main

import (
	"context"
	"fmt"
	"grpc-case/client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("keys/example.com.cert", "www.example.com")
	if err != nil {
		log.Fatal(err)
	}

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
