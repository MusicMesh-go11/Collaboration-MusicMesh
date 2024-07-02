package main

import (
	"log"
	"net"
	
	"github.com/MusicMesh-go11/Collabration_Service/storage/service"
	pb "github.com/MusicMesh-go11/Collabration_Service/genproto/collabration"
	"github.com/MusicMesh-go11/Collabration_Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {

    database, err := postgres.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer database.Close()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    collaborationService := service.NewCollaborationService(database)
    pb.RegisterCollaborationServiceServer(s, collaborationService)

    log.Println("Starting gRPC server on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}