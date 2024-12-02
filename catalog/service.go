package catalog

import (
	"context"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostProduct(ctx context.Context, name string, description string, price float64) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type catalogService struct {
	repository Repository
}

func NewCatalogService(repository Repository) *catalogService {
	return &catalogService{repository}
}

func (cs *catalogService) PostProduct(ctx context.Context, name string, description string, price float64) (*Product, error) {
	p := &Product{
		ID: ksuid.New().String(),
		Name: name,
		Description: description,
		Price: price,
	}

	if err := cs.repository.PutProduct(ctx, *p); err != nil {
		return nil, err
	}

	return p, nil;
}

func (cs *catalogService) GetProduct(ctx context.Context, id string) (*Product, error) {
	return cs.repository.GetProductByID(ctx, id);
}

func (cs *catalogService) GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return cs.repository.ListProducts(ctx, skip, take);
}

func (cs *catalogService) GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error) {
	return cs.repository.ListProductsWithIDs(ctx, ids);
}

func (cs *catalogService) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return cs.repository.SearchProducts(ctx, query, skip, take);
}