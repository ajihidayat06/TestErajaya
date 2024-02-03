package router

import (
	"TestErajaya/model"
	"TestErajaya/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductRouter struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductRouter(productUseCase *usecase.ProductUseCase) ProductRouter {
	return ProductRouter{ProductUseCase: *productUseCase}
}

func (r *ProductRouter) Route(app *fiber.App) {
	app.Post("/api/product", r.InsertProduct)
	app.Get("/api/product", r.GetAll)
}

func (r *ProductRouter) InsertProduct(ctx *fiber.Ctx) error {
	var product model.Product
	_, err := ReadRequest(&product, ctx)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// validation
	err = Validator.Struct(product)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// save data
	response, err := r.ProductUseCase.Create(product)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success insert product",
		Data:       response,
	})
}

func (r *ProductRouter) GetAll(ctx *fiber.Ctx) (err error) {
	var userParam model.ListDataRequestStruct
	userParam, err = ReadRequestListData(&userParam, ctx)
	if err != nil {
		return
	}

	responses := r.ProductUseCase.List(userParam)
	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success get list products",
		Data:       responses,
	})
}
