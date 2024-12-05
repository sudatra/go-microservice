package order

import (
	"context"
	"time"
)

type service interface {
	PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice float64
	AccountID  string
	Products []OrderedProduct
}

type OrderedProduct struct {
	ID string
	Name string
	Description string
	Price float64
	Quantity uint32
}

type orderService struct {
	repository Repository
}

func NewOrderService(repository Repository) *orderService {
	return &orderService{repository}
}

func (os *orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {

}

func (os *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	
}
