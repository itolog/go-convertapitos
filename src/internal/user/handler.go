package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/req"
	"strconv"
)

// HandlerDeps contains dependencies for the user handler
type HandlerDeps struct {
	*configs.Config
	UserServices *Service
}

// Handler for user requests
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

	router.Get("/", handler.GetAllUsers)
	router.Get("/:id", handler.GetUserById)
	router.Get("/by_email/:email", handler.GetUserByEmail)
	router.Post("/", handler.CreateUser)
	router.Patch("/:id", handler.UpdateUser)
	router.Delete("/:id", handler.DeleteUser)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Returns a list of all users with pagination
// @Tags User
// @Accept json
// @Produce json
// @Param limit query int false "Number of records per page" default(10)
// @Param offset query int false "Pagination offset" default(0)
// @Success 200 {object} api.ResponseData{data=[]User,meta=api.Meta} "Successful response with list of users"
// @Failure 400 {object} api.ResponseError{error=string} "Bad request error"
// @Router /user [get]
func (h *Handler) GetAllUsers(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid limit: %s", api.ErrMustBeANumber))
	}
	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid offset: %s", api.ErrMustBeANumber))
	}

	users, err := h.UserServices.FindAll(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data: users.Users,
		Meta: &api.Meta{
			Count: *users.Count,
		},
		Status: api.StatusSuccess,
	})
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Returns user data by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} api.ResponseData{data=User} "Successful response with user data"
// @Failure 400 {object} api.ResponseError "Bad request error"
// @Router /user/{id} [get]
func (h *Handler) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.UserServices.FindById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Returns user data by email
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} api.ResponseData{data=User} "Successful response with user data"
// @Failure 400 {object} api.ResponseError "Bad request error"
// @Router /user/by_email/{email} [get]
func (h *Handler) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	user, err := h.UserServices.FindByEmail(email)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}

// CreateUser godoc
// @Summary Create new user
// @Description Creates a new user with provided data
// @Tags User
// @Accept json
// @Produce json
// @Param user body CreateRequest true "User data"
// @Success 201 {object} api.ResponseData{data=User} "Successfully created user"
// @Failure 400 {object} api.ResponseError "Bad request error"
// @Router /user [post]
func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
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

	created, err := h.UserServices.Create(user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(api.Response{
		Data:   created,
		Status: api.StatusSuccess,
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Updates existing user data
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateRequest true "Update data"
// @Success 200 {object} api.ResponseData{data=User} "Successfully updated user"
// @Failure 400 {object} api.ResponseError "Bad request error"
// @Router /user/{id} [patch]
func (h *Handler) UpdateUser(ctx *fiber.Ctx) error {
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

	updatedUser, err := h.UserServices.Update(id, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   updatedUser,
		Status: api.StatusSuccess,
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Deletes a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} api.ResponseData{data=string} "Success message"
// @Failure 400 {object} api.ResponseError "Bad request error"
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := h.UserServices.Delete(id)
	if err != nil {
		statusCode := api.GetErrorCode(err)
		return fiber.NewError(statusCode, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   fmt.Sprintf("User with id %s deleted", id),
		Status: api.StatusSuccess,
	})
}
