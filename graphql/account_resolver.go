package main

import "context"

type accountResolver struct {
	server *Server
}

func (ar *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {

}