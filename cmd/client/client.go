package main

import (
	"context"
	pb "github.com/cgqaq/grpc-go-hello-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:6668", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := pb.NewHelloClient(conn)
	helloResponse, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "grpc client"})
	if err != nil {
		log.Printf("failed to call SayHello: %v", err)
		return
	}
	log.Printf("response: %s", helloResponse.Message)

	clockStream, err := client.ClockStream(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("failed to call ClockStream: %v", err)
		return
	}

	for {
		var resp pb.ClockResponse
		err := clockStream.RecvMsg(&resp)
		if err != nil {
			log.Printf("failed to receive ClockResponse: %v", err)
			break
		}

		log.Printf("ClockResponse: %s", resp.Datetime)
	}
}
