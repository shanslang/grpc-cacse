package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (this *OrderService) NewOrder(ctx context.Context, orderMain *OrderMain) (*OrderResponse, error) {
	fmt.Println(orderMain)
	return &OrderResponse{
		Status:  "OK",
		Message: "success",
	}, nil
}

func (this *OrderService) mustEmbedUnimplementedOrderServiceServer() {

}
