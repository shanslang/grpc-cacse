package services

import (
	context "context"
)

type ProductService struct {
}

func (this *ProductService) GetProductStock(ctx context.Context, in *ProductRequest) (*ProductRespones, error) {
	var stock int32 = 0
	if in.ProductArea == ProdctAreas_A {
		stock = 30
	} else if in.ProductArea == ProdctAreas_B {
		stock = 31
	} else {
		stock = 32
	}
	return &ProductRespones{ProductStock: stock}, nil
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
