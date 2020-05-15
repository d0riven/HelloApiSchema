package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/d0riven/HelloApiSchema/server/golang/proto/pkg/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:5555", "gRPC server endpoint")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serveMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, serveMux, *grpcServerEndpoint, opts); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		panic(err)
	}
}
