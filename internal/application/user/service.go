package user

import (
	"context"
	"table_link/internal/domain/entity"
	"table_link/internal/domain/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUsers(ctx context.Context, user *entity.User) error
	UpdateUsers(ctx context.Context, user *entity.User) error
	DeleteUsers(ctx context.Context, userID int) error
	GetRoleRightByRoleIDAndRoute(ctx context.Context, roleID, route, section string) (*entity.RoleRight, error)
}

type service struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) Service {
	return &service{userRepo: userRepo}
}

func (s *service) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *service) CreateUsers(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.LastAccess = time.Now()

	return s.userRepo.CreateUsers(ctx, user)
}

func (s *service) UpdateUsers(ctx context.Context, user *entity.User) error {
	return s.userRepo.UpdateUsers(ctx, user)
}

func (s *service) DeleteUsers(ctx context.Context, userID int) error {
	return s.userRepo.DeleteUsers(ctx, userID)
}

func (s *service) GetRoleRightByRoleIDAndRoute(ctx context.Context, roleID, route, section string) (*entity.RoleRight, error) {
	return s.userRepo.GetRoleRightByRoleIDAndRoute(ctx, roleID, route, section)
}
