package service

import "redirect_api/internal/repository"

// collect every "service" interface
type Service struct {
	repo *repository.Repository
}

// link interface and implementation
func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
