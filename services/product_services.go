package services

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServices interface {
	Create(request model.ProductRequest) (response model.ProductResponse)
	ListProduct() (response []model.ProductResponse)
	ProductById(productId string) (response model.ProductResponse)
	UpdateProduct(productId string, request model.ProductRequest) *mongo.UpdateResult
	DeleteProduct(productId string) *mongo.DeleteResult
}
