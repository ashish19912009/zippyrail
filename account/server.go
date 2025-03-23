package account

import (
	"context"
	"fmt"
	"net"

	"github.com/ashish19912009/zippyrail/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           s,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	a, err := s.service.PostAccount(ctx, r.MobileNo)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:       a.ID,
			MobileNo: a.MobileNo,
		},
	}, nil
}

func (s *grpcServer) UpdateAccount(ctx context.Context, r *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	a, err := s.service.UpdateAccount(ctx, r.MobileNo, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountResponse{
		Account: &pb.Account{
			Id:       a.ID,
			MobileNo: a.MobileNo,
			Name:     a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:       a.ID,
			MobileNo: a.MobileNo,
			Name:     a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(accounts,
			&pb.Account{
				Id:       p.ID,
				MobileNo: p.MobileNo,
				Name:     p.Name,
			})
	}
	return &pb.GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
