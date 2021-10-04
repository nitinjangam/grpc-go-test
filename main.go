package main

import (
	"context"
	"grpc-go-test/greetpb"
	"grpc-go-test/healthpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
	healthpb.UnimplementedHealthServer
}

func (*server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function called with: %v", in)
	req := in.GetGreetReq()
	log.Printf("req from client: %v", req)
	resp := greetpb.GreetResponse{
		GreetResp: "Hello " + req + "!",
	}
	return &resp, nil
}

func (*server) Watch(in *healthpb.HealthCheckRequest, res healthpb.Health_WatchServer) error {
	log.Printf("Got health-check request: %v", in)
	resp := healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}
	if err := res.Send(&resp); err != nil {
		return err
	}
	return nil
}

func (*server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	log.Printf("Got Check request: %v", in)
	resp := healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Started listening on port 50055")
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	healthpb.RegisterHealthServer(s, &server{})
	// RegisterHealthServer(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
