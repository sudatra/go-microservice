package main

import (
	"context"
	"log"
	"time"
)

type queryResolver struct {
	server *Server
}

func (qr *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	if id != nil {
		r, err := qr.server.accountClient.GetAccount(ctx, *id);
		if err != nil {
			log.Println(err);
			return nil, err
		}

		return []*Account{{
			ID: r.ID,
			Name: r.Name,
		}}, nil
	}

	skip, take := uint64(0), uint64(0);
	if pagination != nil {
		skip, take = pagination.bounds();
	}

	accountList, err := qr.server.accountClient.GetAccounts(ctx, skip, take);
	if err != nil {
		log.Println(err);
		return nil, err;
	}

	accounts := []*Account{};
	for _, a := range accountList {
		account := &Account{
			ID: a.ID,
			Name: a.Name,
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (qr *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second);
	defer cancel();

	if id != nil {
		r, err := qr.server.catalogCLient.GetProduct(ctx, *id);
		if err != nil {
			log.Println(err);
			return nil, err
		}

		return []*Product{{
			ID: r.ID,
			Name: r.Name,
			Description: r.Description,
			Price: r.Price,
		}}, nil
	}

	skip, take := uint64(0), uint64(0);
	if pagination != nil {
		skip, take = pagination.bounds();
	}

	q := "";
	if query != nil {
		q = *query;
	}

	productList, err := qr.server.catalogCLient.GetProducts(ctx, skip, take, nil, q);
	if err != nil {
		log.Println(err);
		return nil, err;
	}

	products := []*Product{};
	for _, p := range productList {
		product := &Product{
			ID: p.ID,
			Name: p.Name,
			Description: p.Description,
			Price: p.Price,
		}

		products = append(products, product);
	}

	return products, nil
}

func (p PaginationInput) bounds() (uint64, uint64) {
	skipValue := uint64(0);
	takeValue := uint64(0);

	if p.Skip != nil {
		skipValue = uint64(*p.Skip);
	}

	if p.Take != nil {
		takeValue = uint64(*p.Take);
	}

	return skipValue, takeValue
}