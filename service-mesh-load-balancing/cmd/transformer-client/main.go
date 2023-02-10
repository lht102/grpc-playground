package main

import (
	"context"
	"log"
	"time"

	transformerpb "github.com/lht102/grpc-playground/service-mesh-load-balancing/proto/transformer"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	v := viper.GetViper()
	v.AutomaticEnv()

	v.SetDefault("TRANSFORMER_SERVICE_URL", "localhost:8083")
	transformerServiceURL := v.GetString("TRANSFORMER_SERVICE_URL")

	v.SetDefault("NUMBER_OF_CALLS", 100)
	numOfCalls := v.GetInt64("NUMBER_OF_CALLS")

	conn, err := grpc.Dial(transformerServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect to transformer service: %v\n", err)
	}
	defer conn.Close()

	client := transformerpb.NewTransformerServiceClient(conn)

	startTime := time.Now()

	for i := int64(1); i <= numOfCalls; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Printf("Transform number input: %d\n", i)

		resp, err := client.TransformNumber(ctx, &transformerpb.TransformNumberRequest{
			Number: i,
		})
		if err != nil {
			log.Printf("Failed to transform number: %v\n", err)
		}

		log.Printf("Transform number output: %d\n", resp.GetResult())

		if resp.GetMessage() != "" {
			log.Println(resp.GetMessage())
		}
	}

	log.Printf("Time elapsed for making %d calls: %s", numOfCalls, time.Since(startTime))
}
