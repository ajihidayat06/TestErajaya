package setup

import (
	"TestErajaya/router"
	"github.com/gofiber/fiber/v2"
)

type RegisterRouterStruct struct {
	ProductRouter router.ProductRouter
	AuthRouter    router.AuthRouter
}

func RegisterRouter(setupUseCase *RegisterUseCaseStruct) RegisterRouterStruct {
	return RegisterRouterStruct{
		ProductRouter: router.NewProductRouter(&setupUseCase.ProductUseCase),
		AuthRouter:    router.NewAuthRouter(),
	}
}

func SetupRouting(r RegisterRouterStruct, app *fiber.App) {
	r.AuthRouter.Route(app)
	r.ProductRouter.Route(app, router.AuthMiddleware)
}
