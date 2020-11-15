package main

import (
	"context"
	"fmt"
	"grpc-case/client/helper"
	. "grpc-case/client/services"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
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

	// 调用订单服务方法
	orderClient := NewOrderServiceClient(conn)
	t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	newOrder, err := orderClient.NewOrder(ctx, &OrderRequest{OrderMain: &OrderMain{OrderId: 1, OrderNo: "ddh", OrderMoney: 33.3, OrderTime: &t}})
	if err != nil {
		log.Fatal("NewOrder: ", err)
	}
	fmt.Println(newOrder) // status:"OK" message:"success"

	// 调用用户服务
	userClient := NewUserServiceClient(conn)
	users := make([]*UserInfo, 0)
	for i := 1; i < 6; i++ {
		users = append(users, &UserInfo{
			UserId: int32(i),
		})
	}
	// userScore, err := userClient.GetUserScore(ctx, &UserScoreRequest{Users: users})
	// if err != nil {
	// 	log.Fatal("GetUserScore", err)
	// }
	// fmt.Println(userScore.Users)

	stream, err := userClient.GetUserScoreByServerStream(ctx, &UserScoreRequest{Users: users})
	if err != nil {
		log.Fatal("GetUserScore", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Recv:", err)
		}
		fmt.Println(res.Users)
	}

	// 客户端流模式发送请求
	streatClient, err := userClient.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal("GetUserScoreByClientStream:", err)
	}
	for j := 1; j < 4; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: int32(i)})
		}
		err := streatClient.Send(&req) // 只是向服务端发送数据
		if err != nil {
			fmt.Println(err)
		}
	}
	res, err := streatClient.CloseAndRecv() // 接收服务端响应数据
	if err != nil {
		log.Fatal("CloseAndRecv", err)
	}
	fmt.Println(res.Users)

	// 双向流模式
	clients, err := userClient.GetUserScoreByCSStream(ctx)
	if err != nil {
		log.Fatal("GetUserScoreByCSStream:", err)
	}
	for j := 1; j < 4; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i := 1; i < 6; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: int32(i)})
		}
		err := clients.Send(&req) // 只是向服务端发送数据
		if err != nil {
			fmt.Println(err)
		}
		res, err := clients.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println("suan:", res.Users)
	}
}
