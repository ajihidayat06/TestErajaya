package setup

import (
	"TestErajaya/router"
	"github.com/gofiber/fiber/v2"
)

type RegisterRouterStruct struct {
	ProductRouter router.ProductRouter
}

func RegisterRouter(setupUseCase *RegisterUseCaseStruct) RegisterRouterStruct {
	return RegisterRouterStruct{
		ProductRouter: router.NewProductRouter(&setupUseCase.ProductUseCase),
	}
}

func SetupRouting(r RegisterRouterStruct, app *fiber.App) {
	r.ProductRouter.Route(app)
}
