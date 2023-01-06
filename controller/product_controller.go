package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zakariawahyu/go-fiber-mongo-clean/exception"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"github.com/zakariawahyu/go-fiber-mongo-clean/services"
	"net/http"
)

type ProductController struct {
	ProductServices services.ProductServices
}

func NewProductController(productServices *services.ProductServices) ProductController {
	return ProductController{ProductServices: *productServices}
}

func (controller *ProductController) Routes(app *fiber.App) {
	app.Post("/product", controller.CreateProduct)
	app.Get("/product", controller.GetAllProduct)
	app.Get("/product/:id", controller.GetProductById)
	app.Put("/product/:id", controller.UpdateProductById)
	app.Delete("/product/:id", controller.DeleteProductById)
}

func (controller *ProductController) CreateProduct(c *fiber.Ctx) error {
	var request model.ProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.ProductServices.Create(request)
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) GetAllProduct(c *fiber.Ctx) error {
	response := controller.ProductServices.ListProduct()
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := controller.ProductServices.ProductById(id)
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) UpdateProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var request model.ProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	_ = controller.ProductServices.ProductById(id)

	response := controller.ProductServices.UpdateProduct(id, request)
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) DeleteProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	
	_ = controller.ProductServices.ProductById(id)

	response := controller.ProductServices.DeleteProduct(id)
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
