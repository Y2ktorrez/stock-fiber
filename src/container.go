package src

import (
	"github.com/mutinho/config"
	"github.com/mutinho/src/handlers"
	"github.com/mutinho/src/repository"
	"github.com/mutinho/src/services"
)

type Container struct {
	UserRepo    *repository.UserRepository
	UserService *services.UserService
	UserHandler *handlers.UserHandler

	ProductRepo    *repository.ProductRepository
	ProductService *services.ProductService
	ProductHandler *handlers.ProductHandler
}

func SetupContainer() *Container {

	//Users
	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//Products
	productRepo := repository.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	return &Container{
		//Users
		UserRepo:    userRepo,
		UserService: userService,
		UserHandler: userHandler,

		//Products
		ProductRepo:    productRepo,
		ProductService: productService,
		ProductHandler: productHandler,
	}

}
