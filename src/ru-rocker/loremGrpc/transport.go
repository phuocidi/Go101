package loremGrpc

import (
	"ru-rocker/loremGrpc/pb"

	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	lorem grpctransport.Handler
}

// implement LoremServer Interface in lorem.pb.go
func (s *grpcServer) Lorem(ctx context.Context, r *pb.LoremRequest) (*pb.LoremResponse, error) {
	_, resp, err := s.lorem.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoremResponse), nil
}

// create new grpc server
func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.LoremServer {
	return &grpcServer{
		lorem: grpctransport.NewServer(
			endpoint.LoremEndpoint,
			DecodeGRPCLoremRequest,
			EncodeGRPCLoremResponse,
		),
	}
}
