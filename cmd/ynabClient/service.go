package ynabClient

type Service[T any] interface {
	List() ([]T, error)
	// Get(id string) ([]T, error)
}
