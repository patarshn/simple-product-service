package service

import (
	"context"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type IProductService interface {
	GetProductByID(ctx context.Context, id int64) (models.Product, error)
	GetAllProduct(ctx context.Context, req models.GetAllProductRequest) (models.Paginate, error)
}

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (models.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) GetAllProduct(ctx context.Context, req models.GetAllProductRequest) (models.Paginate, error) {
	var paginateRes models.Paginate

	req.Offset = (req.Page * req.Limit) - req.Limit

	products, err := s.repo.GetAllProduct(ctx, req)
	if err != nil {
		return models.Paginate{}, err
	}

	paginateRes.Next = false
	loopLimit := len(products)
	if loopLimit > req.Limit {
		loopLimit = req.Limit
		paginateRes.Next = true
	}

	paginateRes.Prev = true
	if req.Offset == 0 {
		paginateRes.Prev = false
	}

	paginateRes.Data = []any{}
	for i := 0; i < loopLimit; i++ {
		paginateRes.Data = append(paginateRes.Data, products[i])
	}

	paginateRes.From = req.Offset + 1
	paginateRes.To = req.Offset + req.Limit
	paginateRes.Page = int64(req.Page)

	return paginateRes, nil
}
