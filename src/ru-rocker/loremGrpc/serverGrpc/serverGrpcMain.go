package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"ru-rocker/loremGrpc"
	"ru-rocker/loremGrpc/pb"
	"syscall"

	"context"

	"google.golang.org/grpc"
)

func main() {

	var (
		gRPCAddr = flag.String("grpc", ":50051",
			"gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// init lorem service
	var svc loremGrpc.Service
	svc = loremGrpc.LoremService{}
	errChan := make(chan error)

	// creating Endpoints struct
	endpoints := loremGrpc.Endpoints{
		LoremEndpoint: loremGrpc.MakeLoremEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := loremGrpc.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterLoremServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
