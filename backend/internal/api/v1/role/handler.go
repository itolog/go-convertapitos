package role

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/itolog/go-convertapitos/backend/middleware"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/req"
	"math"
	"strconv"
)

type HandlerDeps struct {
	*configs.Config
	RoleServices IRoleService
}

type Handler struct {
	*configs.Config
	RoleServices IRoleService
}

func NewHandler(app fiber.Router, deps HandlerDeps) {
	router := app.Group("/role", middleware.Protected())

	handler := Handler{
		Config:       deps.Config,
		RoleServices: deps.RoleServices,
	}

	router.Get("/", handler.GetAllRoles)
	router.Get("/options", handler.GetForOptions)
	router.Get("/:id", handler.GetRoleById)
	router.Post("/", handler.CreateRole)
}

// GetAllRoles godoc
//
//	@Summary		Get all roles
//	@Description	Returns a list of all users with pagination and sorting options
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int											false	"Number of records per page"	default(10)
//	@Param			page		query		int											false	"Page number"					default(1)	minimum(1)
//	@Param			order_by	query		string										false	"Field to order by"				default(updated_at)
//	@Param			desc		query		boolean										false	"Sort in descending order"		default(false)
//	@Success		200			{object}	api.ResponseData{data=[]Role,meta=api.Meta}	"Successful response with list of roles and metadata"
//	@Failure		400			{object}	api.ResponseError{error=string}				"Bad request error"
//	@Router			/role [get]
func (h *Handler) GetAllRoles(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid limit: %s", api.ErrMustBeANumber))
	}
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid page: %s", api.ErrMustBeANumber))
	}
	if page < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "page must be greater than 0")
	}

	desc := ctx.QueryBool("desc", true)
	orderBy := ctx.Query("order_by", "updated_at")

	offset := (page - 1) * limit

	roles, err := h.RoleServices.FindAll(limit, offset, orderBy, desc)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data: roles.Items,
		Meta: &api.Meta{
			Items: *roles.Count,
			Pages: int(math.Ceil(float64(*roles.Count) / float64(limit))),
		},
		Status: api.StatusSuccess,
	})
}

// GetForOptions godoc
//
//	@Summary		Get all roles
//	@Description	Returns a list of all users with pagination and sorting options
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int											false	"Number of records per page"	default(10)
//	@Param			page		query		int											false	"Page number"					default(1)	minimum(1)
//	@Param			order_by	query		string										false	"Field to order by"				default(updated_at)
//	@Param			desc		query		boolean										false	"Sort in descending order"		default(false)
//	@Success		200			{object}	api.ResponseData{data=[]OptionsResponse,meta=api.Meta}	"Successful response with list of roles and metadata"
//	@Failure		400			{object}	api.ResponseError{error=string}				"Bad request error"
//	@Router			/role/options [get]
func (h *Handler) GetForOptions(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid limit: %s", api.ErrMustBeANumber))
	}
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid page: %s", api.ErrMustBeANumber))
	}
	if page < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "page must be greater than 0")
	}

	desc := ctx.QueryBool("desc", true)
	orderBy := ctx.Query("order_by", "updated_at")

	offset := (page - 1) * limit

	roles, err := h.RoleServices.GetForOptions(limit, offset, orderBy, desc)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data: roles.Items,
		Meta: &api.Meta{
			Items: *roles.Count,
			Pages: int(math.Ceil(float64(*roles.Count) / float64(limit))),
		},
		Status: api.StatusSuccess,
	})
}

// GetRoleById godoc
//
//	@Summary		Get user by ID
//	@Description	Returns user data by ID
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Role ID"
//	@Success		200	{object}	api.ResponseData{data=Role}	"Successful response with role data"
//	@Failure		400	{object}	api.ResponseError			"Bad request error"
//	@Router			/role/{id} [get]
func (h *Handler) GetRoleById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	role, err := h.RoleServices.FindById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   role,
		Status: api.StatusSuccess,
	})
}

// CreateRole godoc
//
//	@Summary		Create new user
//	@Description	Creates a new user with provided data
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			role	body		CreateRequest				true	"Role data"
//	@Success		201		{object}	api.ResponseData{data=Role}	"Successfully created user"
//	@Failure		400		{object}	api.ResponseError			"Bad request error"
//	@Router			/role [post]
func (h *Handler) CreateRole(ctx *fiber.Ctx) error {
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

	pl := PermissionsList(payload.Permissions)
	user := Role{
		Name:        payload.Name,
		Permissions: pl,
	}

	created, err := h.RoleServices.Create(user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(api.Response{
		Data:   created,
		Status: api.StatusSuccess,
	})
}
