package repository

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/config"
	"github.com/zakariawahyu/go-fiber-mongo-clean/entity"
	"github.com/zakariawahyu/go-fiber-mongo-clean/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &ProductRepositoryImpl{
		Collection: database.Collection("products"),
	}
}

func (repository *ProductRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"id":       product.Id,
		"name":     product.Name,
		"price":    product.Price,
		"quantity": product.Quantity})
	exception.PanicIfNeeded(err)
}

func (repository *ProductRepositoryImpl) SelectAll() (product []entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		product = append(product, entity.Product{
			Id:       document["id"].(string),
			Name:     document["name"].(string),
			Price:    document["price"].(int64),
			Quantity: document["quantity"].(int32),
		})
	}

	return product
}

func (repository *ProductRepositoryImpl) SelectById(productId string) entity.Product {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	var product entity.Product

	err := repository.Collection.FindOne(ctx, bson.M{"id": productId}).Decode(&product)
	exception.PanicIfNeeded(err)

	return product
}

func (repository *ProductRepositoryImpl) Update(productId string, product entity.Product) *mongo.UpdateResult {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	productUpdate := bson.M{
		"name":     product.Name,
		"price":    product.Price,
		"quantity": product.Quantity,
	}

	result, err := repository.Collection.UpdateOne(ctx, bson.M{"id": productId}, bson.M{"$set": productUpdate})
	exception.PanicIfNeeded(err)

	return result
}

func (repository *ProductRepositoryImpl) Delete(productId string) *mongo.DeleteResult {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	result, err := repository.Collection.DeleteOne(ctx, bson.M{"id": productId})
	exception.PanicIfNeeded(err)

	return result
}
