package main

import (
	"context"
	"log"
	"time"
)

type accountResolver struct {
	server *Server
}

func (ar *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	orderList, err := ar.server.orderClient.GetOrdersForAccount(ctx, obj.ID)
	if err != nil {
		log.Println(err);
		return nil, err
	}

	orders := []*Order{};
	for _, o := range orderList {
		products := []*OrderedProduct{};
		for _, p := range o.Products {
			products = append(products, &OrderedProduct{
				ID: p.ID,
				Name: p.Name,
				Description: p.Description,
				Quantity: int(p.Quantity),
				Price: p.Price,
			})
		}

		orders = append(orders, &Order{
			ID: o.ID,
			CreatedAt: o.CreatedAt,
			TotalPrice: o.TotalPrice,
			Products: products,
		})
	}

	return orders, nil
}