package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/order-service/internal/service"
	"github.com/example/order-service/internal/store"
	proto "github.com/example/order-service/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := flag.String("grpc-addr", ":50051", "gRPC listen address")
	httpAddr := flag.String("http-addr", ":8080", "HTTP gateway listen address")
	flag.Parse()

	// Create store and service
	st := store.NewInMemoryStore()
	svc := service.NewOrderService(st)

	// gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, svc)

	// listen grpc
	l, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		panic(err)
	}

	// start grpc in goroutine
	go func() {
		fmt.Println("gRPC server listening on", *grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			panic(err)
		}
	}()

	// HTTP gateway
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// register
	if err := proto.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, *grpcAddr, opts); err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    *httpAddr,
		Handler: mux,
	}

	go func() {
		fmt.Println("HTTP gateway listening on", *httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("shutting down...")

	grpcServer.GracefulStop()

	ctxShut, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	httpServer.Shutdown(ctxShut)
}
