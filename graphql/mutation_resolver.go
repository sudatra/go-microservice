package main

import "context"

type mutationResolver struct {
	server *Server
}

func (mr *mutationResolver) CreateAccount(ctx context.Context, inp AccountInput) (*Account, error) {

}

func (mr *mutationResolver) CreateProduct(ctx context.Context, inp ProductInput) (*Product, error) {
	
}

func (mr *mutationResolver) CreateOrder(ctx context.Context, inp OrderInput) (*Order, error) {
	
}