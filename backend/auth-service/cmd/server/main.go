package main

import (
	"log"
	"net"

	"github.com/Lemper29/Jarcy/auth-service/internal/config"
	"github.com/Lemper29/Jarcy/auth-service/internal/database"
	"github.com/Lemper29/Jarcy/auth-service/internal/repository"
	"github.com/Lemper29/Jarcy/auth-service/internal/service"
	pb "github.com/Lemper29/Jarcy/gen/go/auth"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Envs

	databasePostgres, err := database.NewPostgresDatabase(cfg.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer databasePostgres.Pool.Close()

	repo := repository.NewRepo(databasePostgres)
	authService := service.NewService(repo)

	lis, err := net.Listen("tcp", ":"+cfg.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, authService)

	log.Printf("Server running on port %s", cfg.PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
