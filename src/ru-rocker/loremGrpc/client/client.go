package client

import (
	"ru-rocker/loremGrpc"
	"ru-rocker/loremGrpc/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// Return new loremGrpc service
func New(conn *grpc.ClientConn) loremGrpc.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "Lorem", "Lorem",
		loremGrpc.EncodeGRPCLoremRequest,
		loremGrpc.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return loremGrpc.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
