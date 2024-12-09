package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/sudatra/go-microservice/order"
)

var (
	ErrorInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (mr *mutationResolver) CreateAccount(ctx context.Context, inp AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	a, err := mr.server.accountClient.PostAccount(ctx, inp.Name);
	if err != nil {
		log.Println(err);
		return nil, err
	}

	return &Account{
		ID: a.ID,
		Name: a.Name,
	}, nil
}

func (mr *mutationResolver) CreateProduct(ctx context.Context, inp ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	p, err := mr.server.catalogCLient.PostProduct(ctx, inp.Name, inp.Description, inp.Price);
	if err != nil {
		log.Println(err);
		return nil, err
	}

	return &Product{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
	}, nil
}

func (mr *mutationResolver) CreateOrder(ctx context.Context, inp OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	products := []order.OrderedProduct{};
	for _, p := range inp.Products {
		if p.Quantity <= 0 {
			return nil, ErrorInvalidParameter
		}

		products = append(products, order.OrderedProduct{
			ID: p.ID,
			Quantity: uint32(p.Quantity),
		})
	}

	o, err := mr.server.orderClient.PostOrder(ctx, inp.AccountID, products);
	if err != nil {
		log.Println(err);
		return nil, err
	}

	return &Order{
		ID: o.ID,
		CreatedAt: o.CreatedAt,
		TotalPrice: o.TotalPrice,
	}, nil
}