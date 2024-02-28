package ynabClient

type Service interface {
	List() (interface{}, error)
	Get(id string) (interface{}, error)
}
