package user

import (
	"context"
	"go.uber.org/zap"
	"poc/internal/domain"
	"poc/internal/repository"
	"poc/internal/tracking"
)

type FindAllUsers interface {
	Execute(ctx context.Context) ([]domain.User, error)
}

type findAllUsers struct {
	userRepository repository.UserRepository
}

func NewFindAllUsers(r repository.UserRepository) FindAllUsers {
	return findAllUsers{
		userRepository: r,
	}
}

func (u findAllUsers) Execute(ctx context.Context) ([]domain.User, error) {
	logger := ctx.Value(tracking.LoggerFlag).(*zap.SugaredLogger)
	logger.Info("Listing all users...")
	users, err := u.userRepository.FindAll()
	if err != nil {
		return make([]domain.User, 0), tracking.WithOperation(err, "findAllUsers.Execute")
	}

	return users, nil
}
