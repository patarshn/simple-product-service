package repository

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"product-service/internal/models"
	"time"
)

type IProductRepository interface {
	GetProductByID(ctx context.Context, id int64) (models.Product, error)
	GetAllProduct(ctx context.Context, req models.GetAllProductRequest) ([]models.Product, error)
}

type ProductRepository struct {
	DB *sql.DB
}

var categoryName = []string{"Cat1", "Cat2", "Cat3", "Cat4", "Cat5"}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &ProductRepository{DB: db}
}

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateProducts(n int) (products []models.Product) {

	for i := 1; i <= n; i++ {
		catId := int64(rand.Intn(4) + 1)
		stock := int64(rand.Intn(2))
		isAvailable := false
		if stock > 0 {
			isAvailable = true
		}
		products = append(products, models.Product{
			ID:           int64(i),
			Name:         "Product " + generateRandomString(5),
			Price:        rand.Int63n(1000000) + 1,
			Description:  "Description for product " + generateRandomString(10),
			CategoryID:   catId,
			CategoryName: categoryName[catId],
			Stock:        stock,
			IsAvailable:  isAvailable,
			CreatedBy:    "user_" + generateRandomString(5),
			CreatedAt:    time.Now().Add(time.Duration(-rand.Intn(10000)) * time.Hour),
			UpdatedAt:    time.Now(),
		})
	}
	return products
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (models.Product, error) {

	products := generateProducts(5)

	product := models.Product{}
	for _, v := range products {
		if v.ID == id {
			product = v
			return product, nil
		}
	}

	return product, errors.New("Product Not Found")

}

func (r *ProductRepository) GetAllProduct(ctx context.Context, req models.GetAllProductRequest) ([]models.Product, error) {

	products := generateProducts(10)
	return products, nil
}
