package service

import (
	"context"

	"gin-crud/model"
	"gin-crud/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {

	return &UserService{
		repo: repo,
	}
}

// Create user:

func (s *UserService) Create(
	ctx context.Context,
	user *model.User,
) error {

	// Business logic can go here.

	return s.repo.Create(ctx, user)
}

// Get users:

func (s *UserService) GetAll(
	ctx context.Context,
) ([]model.User, error) {

	return s.repo.GetAll(ctx)
}

// Get one:
func (s *UserService) GetByID(
	ctx context.Context,
	id int,
) (*model.User, error) {

	return s.repo.GetByID(ctx, id)
}

// Update User:
func (s *UserService) Update(
	ctx context.Context,
	user *model.User,
) error {

	return s.repo.Update(ctx, user)
}

// Delete User:
func (s *UserService) Delete(
	ctx context.Context,
	id int,
) error {

	return s.repo.Delete(ctx, id)
}
