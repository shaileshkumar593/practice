package main

import (
    "context"
    "log"
    "time"

    pb "grpc-go-example/proto"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.NewClient(
        "localhost:50051",
        grpc.WithTransportCredentials(
            insecure.NewCredentials(),
        ),
    )
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(
        context.Background(),
        5*time.Second,
    )
    defer cancel()

    resp, err := client.SayHello(
        ctx,
        &pb.HelloRequest{
            Name: "Shailesh",
        },
    )
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }

    log.Println("Server Response:", resp.Message)
}
