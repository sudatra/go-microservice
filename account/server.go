package account

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	server Service
}

func ListenGRPC(s Service, port int) error {
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port));
  if err != nil {
    return err
  }

  serv := grpc.NewServer();
  reflection.Register(serv);

  return serv.Serve(lis) 
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
  a, err := s.server.PostAccount(ctx, r.Name);
  if err != nil {
    return nil, err;
  }

  return &pb.PostAccountResponse{Account: &pb.Account{
    Id: a.ID,
    Name: a.Name,
  }}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
  a, err := s.server.GetAccount(ctx, r.Id);
  if err != nil {
    return nil, err;
  }

  return &pb.GetAccountResponse{Account: &pb.Account{
    Id: a.ID,
    Name: a.Name,
  }}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
  res, err := s.server.GetAccounts(ctx, r.skip, r.take);
  if err != nil {
    return nil, err;
  }

  accounts := []*pb.Account{};
  for _, a := range res {
    accounts = append(accounts, &pb.Account{
      Id: a.ID,
      Name: a.Name,
    })
  }

  return &pb.GetAccountsResponse{Accounts: accounts}, nil
}