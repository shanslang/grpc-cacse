package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (this *OrderService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderRequest.OrderMain)
	return &OrderResponse{
		Status:  "OK",
		Message: "success",
	}, nil
}

func (this *OrderService) mustEmbedUnimplementedOrderServiceServer() {

}
