package main

import (
	"context"
	"fmt"
	"log"
	"net"

	adderpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/adder"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	adderpb.UnimplementedAdderServiceServer
	podIP string
}

func (s *Server) AddNumbers(
	ctx context.Context,
	request *adderpb.AddNumbersRequest,
) (*adderpb.AddNumbersResponse, error) {
	var msg string
	if s.podIP != "" {
		msg = fmt.Sprintf("Response form adder service (%s)", s.podIP)
	}

	return &adderpb.AddNumbersResponse{
		Result:  lo.Sum(request.Numbers),
		Message: msg,
	}, nil
}

func main() {
	v := viper.GetViper()
	v.AutomaticEnv()

	v.SetDefault("GRPC_PORT", 8081)
	port := v.GetInt("GRPC_PORT")

	podIP := v.GetString("MY_POD_IP")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcSrv := grpc.NewServer()
	adderpb.RegisterAdderServiceServer(grpcSrv, &Server{
		podIP: podIP,
	})
	reflection.Register(grpcSrv)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
