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

	MovimentRepo    *repository.MovimentRepository
	MovimentService *services.MovimentService
	MovimentHandler *handlers.MovimentHandler
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

	//Moviments
	movimentRepo := repository.NewMovimentRepository(config.DB)
	movimentService := services.NewMovimentService(movimentRepo)
	movimentHandler := handlers.NewMovimentHandler(movimentService)

	return &Container{
		//Users
		UserRepo:    userRepo,
		UserService: userService,
		UserHandler: userHandler,

		//Products
		ProductRepo:    productRepo,
		ProductService: productService,
		ProductHandler: productHandler,

		//Moviments
		MovimentRepo:    movimentRepo,
		MovimentService: movimentService,
		MovimentHandler: movimentHandler,
	}

}
