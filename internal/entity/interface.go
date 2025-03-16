package entity

type OrderRepositoryInterface interface {
	FindAll() ([]Order, error)
}
