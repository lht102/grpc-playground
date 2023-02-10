package main

import (
	"context"
	"fmt"
	"log"
	"net"

	subtractpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/subtractor"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	subtractpb.UnimplementedSubtractorServiceServer
	podIP string
}

func (s *Server) SubtractNumbers(
	ctx context.Context,
	request *subtractpb.SubtractNumbersRequest,
) (*subtractpb.SubtractNumbersResponse, error) {
	if len(request.GetNumbers()) == 0 {
		return &subtractpb.SubtractNumbersResponse{
			Result: 0,
		}, nil
	}

	res := request.GetNumbers()[0]
	for i := 1; i < len(request.GetNumbers()); i++ {
		res -= request.GetNumbers()[i]
	}

	var msg string
	if s.podIP != "" {
		msg = fmt.Sprintf("Response form subtractor service (%s)", s.podIP)
	}

	return &subtractpb.SubtractNumbersResponse{
		Result:  res,
		Message: msg,
	}, nil
}

func main() {
	v := viper.GetViper()
	v.AutomaticEnv()

	v.SetDefault("GRPC_PORT", 8082)
	port := v.GetInt("GRPC_PORT")

	podIP := v.GetString("MY_POD_IP")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcSrv := grpc.NewServer()
	subtractpb.RegisterSubtractorServiceServer(grpcSrv, &Server{
		podIP: podIP,
	})
	reflection.Register(grpcSrv)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
