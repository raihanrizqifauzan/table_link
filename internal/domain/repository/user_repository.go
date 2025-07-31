package repository

import (
	"context"
	"table_link/internal/domain/entity"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUsers(ctx context.Context, user *entity.User) error
	UpdateUsers(ctx context.Context, user *entity.User) error
	DeleteUsers(ctx context.Context, userID int) error
	GetRoleRightByRoleIDAndRoute(ctx context.Context, roleID, route, section string) (*entity.RoleRight, error)
}
