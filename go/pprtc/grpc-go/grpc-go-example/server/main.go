package main

import (
    "context"
    "log"
    "net"

    pb "grpc-go-example/proto"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(
    ctx context.Context,
    req *pb.HelloRequest,
) (*pb.HelloReply, error) {

    log.Printf("Received request from: %s", req.Name)

    return &pb.HelloReply{
        Message: "Hello " + req.Name + " from gRPC server!",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()

    pb.RegisterGreeterServer(grpcServer, &server{})

    log.Println("gRPC server running on :50051")

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
