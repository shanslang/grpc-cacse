package services

import (
	context "context"
)

type ProductService struct {
}

func (this *ProductService) GetProductStock(ctx context.Context, in *ProductRequest) (*ProductRespones, error) {
	return &ProductRespones{ProductStock: 20}, nil
}

func (this *ProductService) GetProductStocks(ctx context.Context, size *QuerySize) (*ProductResponesList, error) {
	products := []*ProductRespones{
		&ProductRespones{ProductStock: 20},
		&ProductRespones{ProductStock: 21},
		&ProductRespones{ProductStock: 22},
	}
	return &ProductResponesList{Products: products}, nil
}

func (this *ProductService) mustEmbedUnimplementedProductServiceServer() {

}
