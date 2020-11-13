package services

import (
	context "context"
)

type ProductService struct {
}

func (this *ProductService) GetProductStock(ctx context.Context, in *ProductRequest) (*ProductRespones, error) {
	return &ProductRespones{ProductStock: 20}, nil
}
