package order

import (
	"context"
	"time"
	"github.com/segmentio/ksuid"
)

type Service interface {
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
	o := &Order{
		ID: ksuid.New().String(),
		CreatedAt: ksuid.New().Time().UTC(),
		AccountID: accountID,
		Products: products,
	}

	o.TotalPrice = 0.0;
	for _, p := range products {
		o.TotalPrice += p.Price * float64(p.Quantity);
	}

	err := os.repository.PutOrder(ctx, *o);
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (os *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	return os.repository.GetOrdersForAccount(ctx, accountID);
}
