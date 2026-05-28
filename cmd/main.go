package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HeadBangZ/pulse/api"

	"google.golang.org/grpc"
)

type ingestorService struct {
	pb.UnimplementedPulseServiceServer
}

func (s *ingestorService) SendLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Log received:\n\tProject: %s\n\tService: %s\n\tMessage: %s", req.ProjectId, req.ServiceName, req.RequestMsg)

	return &pb.LogResponse{
		Success:     true,
		ResponseMsg: "Log securely queued by Pulse Ingestor",
	}, nil
}

func main() {
	grpcPort := ":50051"
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to bind TCP listener on %s : %v", grpcPort, err)
	}

	grcpServer := grpc.NewServer()

	log.Printf("Pulse Ingestor accepting gRPC logs on port %s", grpcPort)
	if err := grcpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to maintain gRPC server context: %v", err)
	}
}
