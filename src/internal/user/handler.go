package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

type HandlerDeps struct {
	*configs.Config
	UserServices *Service
}

type Handler struct {
	*configs.Config
	UserServices *Service
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	router := app.Group("/user")

	handler := Handler{
		Config:       deps.Config,
		UserServices: deps.UserServices,
	}

	router.Get("/", func(ctx *fiber.Ctx) error {
		users, err := handler.UserServices.findAll()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   users,
			Status: api.StatusSuccess,
		})
	})
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user, err := handler.UserServices.findById(id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   user,
			Status: api.StatusSuccess,
		})
	})
	router.Get("/by_email/:email", func(ctx *fiber.Ctx) error {
		email := ctx.Params("email")

		user, err := handler.UserServices.findByEmail(email)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   user,
			Status: api.StatusSuccess,
		})
	})
	router.Post("/", func(ctx *fiber.Ctx) error {
		payload, err := req.DecodeBody[CreateRequest](ctx)
		if err != nil {
			return err
		}
		validateError, valid := req.ValidateBody(payload)
		if !valid {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error:  validateError,
				Status: api.StatusError,
			})
		}

		user := User{
			Name:          payload.Name,
			Email:         payload.Email,
			VerifiedEmail: payload.VerifiedEmail,
			Picture:       payload.Picture,
			Password:      payload.Password,
		}

		created, err := handler.UserServices.create(user)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(api.Response{
			Data:   created,
			Status: api.StatusSuccess,
		})
	})
	router.Patch("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		payload, err := req.DecodeBody[UpdateRequest](ctx)
		if err != nil {
			return err
		}

		validateError, valid := req.ValidateBody(payload)
		if !valid {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error:  validateError,
				Status: api.StatusError,
			})
		}

		updatedUser, err := handler.UserServices.update(id, payload)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   updatedUser,
			Status: api.StatusSuccess,
		})
	})
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		err := handler.UserServices.delete(id)
		if err != nil {
			statusCode := api.GetErrorCode(err)
			return ctx.Status(statusCode).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   fmt.Sprintf("User with id %s deleted", id),
			Status: api.StatusSuccess,
		})
	})
}
