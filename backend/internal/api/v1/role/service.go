package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/common"
)

type IRoleService interface {
	FindAll(limit int, offset int, orderBy string, desc bool) (*common.FindAllResponse[Role], error)
	FindById(id string) (*Role, error)
	Create(user Role) (*Role, error)
	//Update(id string, payload *Role) (*Role, error)
	//Delete(id string) error
	//BatchDelete(ids *[]string) error
}

type Service struct {
	RoleRepository IRepository
}

func NewService(repository *Repository) *Service {
	return &Service{
		RoleRepository: repository,
	}
}

func (service *Service) FindAll(limit int, offset int, orderBy string, desc bool) (*common.FindAllResponse[Role], error) {
	count := service.RoleRepository.Count()

	order := "asc"
	if desc {
		order = "desc"
	}
	roles, err := service.RoleRepository.FindAll(limit, offset, orderBy, order)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return &common.FindAllResponse[Role]{
		Items: roles,
		Count: count,
	}, nil
}

func (service *Service) FindById(id string) (*Role, error) {
	user, err := service.RoleRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return user, nil
}

func (service *Service) Create(role Role) (*Role, error) {
	created, err := service.RoleRepository.Create(&role)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return created, nil
}
