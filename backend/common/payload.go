package common

type FindAllResponse[T any] struct {
	Items []T
	Count *int64
}
