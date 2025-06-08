package database

type ICrud[T any] interface {
	FindAll(limit int, offset int, orderBy string, order string) ([]T, error)
	FindById(id string) (*T, error)
	Create(item *T) (*T, error)
	Update(payload *T) (*T, error)
	Delete(id string) error
	BatchDelete(ids *[]string) error
	Count() *int64
}

type BatchDeleteRequest struct {
	Ids []string `json:"ids" validate:"required,min=1,dive,uuid"`
}
