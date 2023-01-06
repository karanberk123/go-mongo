package repository

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Insert(product entity.Product)
	SelectAll() (product []entity.Product)
	SelectById(productId string) entity.Product
	Update(productId string, product entity.Product) *mongo.UpdateResult
	Delete(productId string) *mongo.DeleteResult
}
