package auth

import (
	"context"
	"fmt"
	"table_link/internal/domain/entity"
	"table_link/internal/domain/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthCache interface {
	StoreUserSession(token string, user *entity.User, expiration time.Duration) error
	DeleteUserSession(token string) error
}

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, token string) error
}

type service struct {
	userRepo  repository.UserRepository
	authCache AuthCache
}

func NewService(userRepo repository.UserRepository, authCache AuthCache) Service {
	return &service{
		userRepo:  userRepo,
		authCache: authCache,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token := uuid.New().String()

	if err := s.authCache.StoreUserSession(token, user, 30*time.Minute); err != nil {
		return "", err
	}

	user.LastAccess = time.Now()

	return token, nil
}

func (s *service) Logout(ctx context.Context, token string) error {
	if err := s.authCache.DeleteUserSession(token); err != nil {
		return err
	}

	return nil
}
