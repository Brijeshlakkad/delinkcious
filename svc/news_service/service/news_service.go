package service

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Brijeshlakkad/delinkcious/pb/news_service/pb"
	nm "github.com/Brijeshlakkad/delinkcious/pkg/news_manager"
	"google.golang.org/grpc"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6060"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	natsHostname := os.Getenv("NATS_CLUSTER_SERVICE_HOST")
	natsPort := os.Getenv("NATS_CLUSTER_SERVICE_PORT")
	svc, err := nm.NewNewsManager(natsHostname, natsPort)
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterNewsServer(gRPCServer, newNewsServer(svc))

	fmt.Printf("News service is listening on port %s...\n", port)
	err = gRPCServer.Serve(listener)
	fmt.Println("Serve() failed", err)
}
