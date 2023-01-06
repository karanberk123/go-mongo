package services

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/entity"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"github.com/zakariawahyu/go-fiber-mongo-clean/repository"
	validation2 "github.com/zakariawahyu/go-fiber-mongo-clean/validation"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServicesImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServices(productRepository *repository.ProductRepository) ProductServices {
	return &ProductServicesImpl{
		ProductRepository: *productRepository,
	}
}

func (services *ProductServicesImpl) Create(request model.ProductRequest) (response model.ProductResponse) {
	validation2.ValidateProduct(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	services.ProductRepository.Insert(product)

	response = model.ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (services *ProductServicesImpl) ListProduct() (response []model.ProductResponse) {
	products := services.ProductRepository.SelectAll()

	for _, product := range products {
		response = append(response, model.ProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return response
}

func (services *ProductServicesImpl) ProductById(productId string) (response model.ProductResponse) {
	product := services.ProductRepository.SelectById(productId)

	response = model.ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (services *ProductServicesImpl) UpdateProduct(productId string, request model.ProductRequest) *mongo.UpdateResult {
	validation2.ValidateProduct(request)
	product := entity.Product{
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	result := services.ProductRepository.Update(productId, product)
	return result
}

func (services *ProductServicesImpl) DeleteProduct(productId string) *mongo.DeleteResult {
	result := services.ProductRepository.Delete(productId)

	return result
}
