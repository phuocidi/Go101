package loremGrpc

import (
	"context"
	"ru-rocker/loremGrpc/pb"
)

// EncodeGRPCLoremRequest ...Encode Lorem Request
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(LoremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// DecodeGRPCLoremRequest ... Decode Lorem Request
func DecodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)
	return LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// EncodeGRPCLoremResponse ...Encode Lorem Response
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(LoremResponse)
	return &pb.LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}

// DecodeGRPCLoremResponse ...Decode Lorem Response
func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)
	return LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}
