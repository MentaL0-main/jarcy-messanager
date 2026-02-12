package service

import (
	"github.com/Lemper29/Jarcy/auth-service/internal/repository"
	pb "github.com/Lemper29/Jarcy/gen/go/auth"
)

type Service struct {
	pb.UnimplementedAuthServer
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
