package main

import (
	"context"
	"grpc-go-test/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Client dialing on port 50055")
	cc, err := grpc.Dial("0.0.0.0:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dial: %v", err)
	}
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := greetpb.GreetRequest{
		GreetReq: "nitin",
	}
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res.GetGreetResp())
}
