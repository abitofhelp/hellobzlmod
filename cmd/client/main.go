package main

import (
	"context"
	v1 "github.com/abitofhelp/hellobzlmod/bazel-bin/proto/abitofhelp/api/helloworld/v1/abitofhelp_api_helloworld_v1_go_proto_/github.com/abitofhelp/hellobzlmod/proto/abitofhelp/api/helloworld/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {

	now := time.Now().UTC()
	timestamp := timestamppb.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	req := v1.HelloRequest{
		Name:    "Mikie",
		SentUtc: &timestamp,
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
