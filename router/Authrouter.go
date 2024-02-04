package router

import (
	"TestErajaya/config"
	"TestErajaya/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type AuthRouter struct {
}

func NewAuthRouter() AuthRouter {
	return AuthRouter{}
}

func (r *AuthRouter) Route(app *fiber.App) {
	app.Post("/api/login", r.Login)

	app.Use(AuthMiddleware)
	app.Get("/api/logout", r.Logout)
}

func (r *AuthRouter) Login(ctx *fiber.Ctx) (err error) {
	var authModel model.LoginModel
	_, err = ReadRequest(&authModel, ctx)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if authModel.Username == "admin" && authModel.Password == "admin123" {
		token := fmt.Sprintf("%s:%d", authModel.Username, time.Now().UnixNano())

		err = config.Client.Set(config.Client.Context(), token, authModel.Username, time.Hour*1).Err()
		if err != nil {
			return SetResponseJson(ctx, model.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf(`%s, %s`, "Failed to create session", err.Error()),
			})
		}

		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusOK,
			Message:    "success login",
			Data:       token,
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusUnauthorized,
		Message:    fmt.Sprintf(`%s`, "Invalid username or password"),
	})
}

func (r *AuthRouter) Logout(ctx *fiber.Ctx) (err error) {
	token := ctx.Get("Authorization")

	err = config.Client.Del(ctx.Context(), token).Err()
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success logout",
		Data:       nil,
	})

	return
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token == "" {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized user",
		})
	}

	username, err := config.Client.Get(ctx.Context(), token).Result()
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid Token",
		})
	}

	ctx.Locals("username", username)
	return ctx.Next()
}
