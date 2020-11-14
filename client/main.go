package main

import (
	"context"
	"fmt"
	"grpc-case/client/helper"
	. "grpc-case/client/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCreds())) // 连接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	productClient := NewProductServiceClient(conn)
	ctx := context.Background()

	// productRes, err := productClient.GetProductStock(ctx,
	// 	&ProductRequest{ProductId: 1, ProductArea: ProdctAreas_C})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(productRes.ProductStock)

	// products, err := productClient.GetProductStocks(ctx, &services.QuerySize{Size: 10})
	// if err != nil {
	// 	log.Fatal("GetProductStocks: ", err)
	// }
	// fmt.Println(products.Products)

	productModel, err := productClient.GetProductInfo(ctx, &ProductRequest{ProductId: 12})
	if err != nil {
		log.Fatal("GetProductStocks: ", err)
	}
	fmt.Println(productModel) // product_id:101 product_name:"书本" product_price:23.3

}
