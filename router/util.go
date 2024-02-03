package router

import (
	"TestErajaya/model"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var Validator = validator.New()

func ReadRequest(req interface{}, ctx *fiber.Ctx) (id int64, err error) {

	if ctx.Method() != "GET" {
		err = ctx.BodyParser(req)
		if err != nil {
			return
		}
	}
	fmt.Println(ctx.Params("id"))

	if ctx.Params("id") != "" {
		idInt, errs := ReadRequestParams(ctx)
		if errs != nil {
			err = errs
			return
		}
		id = int64(idInt)
	}

	return
}

func SetResponseJson(ctx *fiber.Ctx, resp model.Response) error {
	response := model.Response{
		StatusCode: resp.StatusCode,
		Message:    resp.Message,
		Data:       resp.Data,
	}
	return ctx.Status(resp.StatusCode).JSON(response)
}

func ReadRequestParams(ctx *fiber.Ctx) (id int64, err error) {
	idInt, errs := strconv.Atoi(ctx.Params("id"))
	if errs != nil {
		err = errs
		return
	}
	fmt.Println(id)
	id = int64(idInt)

	return
}

func ReadRequestListData(req interface{}, ctx *fiber.Ctx) (userParam model.ListDataRequestStruct, err error) {
	if ctx.Query("order") != "" {
		userParam.Order = ReadRequestParamsOrder(ctx)
	}

	if ctx.Query("page") != "" {
		userParam.Page, err = ReadRequestParamsPage(ctx)
		if err != nil {
			return
		}
	}

	if ctx.Query("limit") != "" {
		userParam.Limit, err = ReadRequestParamsLimit(ctx)
		if err != nil {
			return
		}
	}

	return
}

func ReadRequestParamsOrder(ctx *fiber.Ctx) (orderBy string) {
	orderBy = ctx.Query("order")
	return
}

func ReadRequestParamsPage(ctx *fiber.Ctx) (page int64, err error) {
	pageInt, errs := strconv.Atoi(ctx.Query("page"))
	if errs != nil {
		err = errs
		return
	}
	page = int64(pageInt)

	return
}

func ReadRequestParamsLimit(ctx *fiber.Ctx) (limit int64, err error) {
	limitInt, errs := strconv.Atoi(ctx.Query("limit"))
	if errs != nil {
		err = errs
		return
	}
	limit = int64(limitInt)

	return
}
