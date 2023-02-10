package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	adderpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/adder"
	subtractpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/subtractor"
	transformerpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/transformer"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const randMax = 1000

type Server struct {
	transformerpb.UnimplementedTransformerServiceServer
	adderServiceClient      adderpb.AdderServiceClient
	subtractorServiceClient subtractpb.SubtractorServiceClient
	podIP                   string
}

func (s *Server) TransformNumber(
	ctx context.Context,
	request *transformerpb.TransformNumberRequest,
) (*transformerpb.TransformNumberResponse, error) {
	randNum1, err := rand.Int(rand.Reader, big.NewInt(randMax))
	if err != nil {
		return nil, fmt.Errorf("generate random number: %w", err)
	}

	adderResp, err := s.adderServiceClient.AddNumbers(ctx, &adderpb.AddNumbersRequest{
		Numbers: []int64{request.GetNumber(), randNum1.Int64()},
	})
	if err != nil {
		return nil, fmt.Errorf("add numbers: %w", err)
	}

	if adderResp.GetMessage() != "" {
		log.Println(adderResp.GetMessage())
	}

	randNum2, err := rand.Int(rand.Reader, big.NewInt(randMax))
	if err != nil {
		return nil, fmt.Errorf("generate random number: %w", err)
	}

	subtractorResp, err := s.subtractorServiceClient.SubtractNumbers(ctx, &subtractpb.SubtractNumbersRequest{
		Numbers: []int64{adderResp.GetResult(), randNum2.Int64()},
	})
	if err != nil {
		return nil, fmt.Errorf("subtract numbers: %w", err)
	}

	if subtractorResp.GetMessage() != "" {
		log.Println(subtractorResp.GetMessage())
	}

	var msg string
	if s.podIP != "" {
		msg = fmt.Sprintf("Response form tranformer service (%s)", s.podIP)
	}

	return &transformerpb.TransformNumberResponse{
		Result:  subtractorResp.GetResult(),
		Message: msg,
	}, nil
}

func main() {
	v := viper.GetViper()
	v.AutomaticEnv()

	v.SetDefault("GRPC_PORT", 8083)
	port := v.GetInt("GRPC_PORT")

	v.SetDefault("ADDER_SERVICE_URL", "localhost:8081")
	adderServiceURL := v.GetString("ADDER_SERVICE_URL")

	v.SetDefault("SUBTRACT_SERVICE_URL", "localhost:8082")
	subtractorServiceURL := v.GetString("SUBTRACT_SERVICE_URL")

	podIP := v.GetString("MY_POD_IP")

	adderServiceConn, err := grpc.Dial(adderServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect to adder service: %v\n", err)
	}
	defer adderServiceConn.Close()

	adderServiceClient := adderpb.NewAdderServiceClient(adderServiceConn)

	subtractorServiceConn, err := grpc.Dial(subtractorServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect to subtractor service: %v\n", err)
	}
	defer subtractorServiceConn.Close()

	subtractorServiceClient := subtractpb.NewSubtractorServiceClient(subtractorServiceConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panicf("Failed to listen: %v\n", err)
	}

	grpcSrv := grpc.NewServer()
	transformerpb.RegisterTransformerServiceServer(grpcSrv, &Server{
		adderServiceClient:      adderServiceClient,
		subtractorServiceClient: subtractorServiceClient,
		podIP:                   podIP,
	})
	reflection.Register(grpcSrv)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Panicf("Failed to serve: %v\n", err)
	}
}
