package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/daily-summary-api/handler"
	proto "github.com/micro/services/portfolio/daily-summary-api/proto"
	auth "github.com/micro/services/portfolio/helpers/authentication"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-api-v1-daily-summary"),
		micro.Version("latest"),
	)
	service.Init()

	// Setup the auth packagee
	auth, err := auth.New(os.Getenv("JWT_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Could not initiate auth package: %v\n", err)
		os.Exit(2)
	}

	// Register to Service Discovery
	hander := handler.New(auth, service.Client())
	proto.RegisterDailySummaryHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
