package main

import (
	"net"
	"os"
	"time"

	"github.com/d0riven/HelloApiSchema/server/golang/proto/pkg/pb"
	"github.com/d0riven/HelloApiSchema/server/golang/proto/pkg/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// @ref https://qiita.com/iamtakeruXO/items/948316f29d8f7901f722

func main() {
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
		),
	)
	pb.RegisterUserServiceServer(server, service.NewUserService())
	logger.Info("run server")
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
