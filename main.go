package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/zakariawahyu/go-fiber-mongo-clean/config"
	"github.com/zakariawahyu/go-fiber-mongo-clean/controller"
	"github.com/zakariawahyu/go-fiber-mongo-clean/repository"
	"github.com/zakariawahyu/go-fiber-mongo-clean/services"
	"net/http"
)

func main() {
	// Run Database
	database := config.NewMongoDatabase()

	// Setup Repository
	productRepository := repository.NewProductRepository(database)

	// Setup Services
	productServices := services.NewProductServices(&productRepository)

	// Setup Controller
	productController := controller.NewProductController(&productServices)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})
	productController.Routes(app)
	app.Listen(":8081")
}
