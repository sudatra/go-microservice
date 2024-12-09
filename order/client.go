package order

import (
	"context"
	"log"
	"time"

	"github.com/sudatra/go-microservice/order/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.OrderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure());
	if err != nil {
		return nil, err;
	}

	c := pb.NewOrderServiceClient(conn);
	return &Client{
		conn: conn,
		service: c,
	}, nil
}

func (c *Client) CLose() {
	c.conn.Close();
}

func (c *Client) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{};
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity: p.Quantity,
		})
	}

	r, err := c.service.PostOrder(
		ctx,
		&pb.PostOrderRequest{
			AccountId: accountID,
			Products: protoProducts,
		},
	);
	if err != nil {
		return nil, err
	}

	newOrder := r.Order;
	newOrderCreatedAt := time.Time{};
	newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedAt);

	return &Order{
		ID: newOrder.Id,
		AccountID: newOrder.AccountId,
		Products: products,
		TotalPrice: newOrder.TotalPrice,
		CreatedAt: newOrderCreatedAt,
	}, nil
}

func (c *Client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	r, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{AccountId: accountID});
	if err != nil {
		log.Println(err);
		return nil, err;
	}

	orders := []Order{};
	for _, orderProto := range r.Orders {
		newOrder := Order{
			ID: orderProto.Id,
			AccountID: orderProto.AccountId,
			TotalPrice: orderProto.TotalPrice,
		}
		
		newOrder.CreatedAt = time.Time{};
		newOrder.CreatedAt.UnmarshalBinary(orderProto.CreatedAt);

		products := []OrderedProduct{};
		for _, p := range orderProto.Products {
			products = append(products, OrderedProduct{
				ID: p.Id,
				Name: p.Name,
				Description: p.Description,
				Price: p.Price,
				Quantity: p.Quantity,
			})
		}

		newOrder.Products = products
		orders = append(orders, newOrder);
	}

	return orders, nil;
}