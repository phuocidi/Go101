package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"ru-rocker/loremGrpc"
	"ru-rocker/loremGrpc/pb"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var (
		grpcAddr = flag.String("addr", "https://localhost:50051",
			"gRPC address")
	)
	flag.Parse()
	fmt.Println("\nClient is prepared to listening on %s\n", *grpcAddr)

	ctx := context.Background()

	selfCredential, errCert := credentials.NewClientTLSFromFile("server.pem", "")
	if errCert != nil {
		log.Fatalln("errCert:", errCert)
		return
	}
	conn, err := grpc.Dial(*grpcAddr,
		grpc.WithTransportCredentials(selfCredential),
		grpc.WithTimeout(1*time.Second))
	// conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
	// 	grpc.WithTimeout(1*time.Second))

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	//loremService := grpcClient.New(conn)
	loremService := pb.NewLoremClient(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	switch cmd {
	case "lorem":
		var requestType, minStr, maxStr string

		requestType, args = pop(args)
		minStr, args = pop(args)
		maxStr, args = pop(args)

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		lorem(ctx, loremService, requestType, min, max)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

// parse command line argument one by one
func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// call lorem service
func lorem(ctx context.Context, service loremGrpc.Service, requestType string, min int, max int) {
	fmt.Println("\n\nHIT 1")
	mesg, err := service.Lorem(ctx, requestType, min, max)
	if err != nil {
		fmt.Println("\n\nHIT error")
		log.Fatalln(err.Error())
	}
	fmt.Println("\n\nHIT 2")
	fmt.Println(mesg)
}
