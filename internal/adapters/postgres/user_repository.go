package postgres

import (
	"context"
	"database/sql"
	"table_link/internal/domain/entity"
	"table_link/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	query := `SELECT u.id, u.role_id, r.name, u.name, u.email, u.password, u.last_access FROM users u JOIN roles r ON u.role_id = r.id WHERE u.email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(user.ID, user.RoleID, user.RoleName, user.Name, user.Email, user.Password, user.LastAccess)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT u.id, u.role_id, r.name, u.name, u.email, u.last_access FROM users u JOIN roles r ON u.role_id = r.id`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var data []*entity.User
	for rows.Next() {
		user := new(entity.User)
		if err := rows.Scan(user.ID, user.RoleID, user.RoleName, user.Name, user.Email, user.LastAccess); err != nil {
			return nil, err
		}

		data = append(data, user)
	}

	return data, err
}

func (r *userRepository) CreateUsers(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (role_id, name, email, password, last_access) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, user.RoleID, user.Name, user.Email, user.Password, user.LastAccess)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUsers(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUsers(ctx context.Context, userID int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetRoleRightByRoleIDAndRoute(ctx context.Context, roleID, route, section string) (*entity.RoleRight, error) {
	roleRight := new(entity.RoleRight)
	query := `SELECT id, role_id, section, route, r_create, r_read, r_update, r_delete FROM role_rights WHERE role_id = $1 AND route = $2 AND section = $3`
	err := r.db.QueryRowContext(ctx, query, roleID, route, section).Scan(
		roleRight.ID, roleRight.RoleID, roleRight.Section, roleRight.Route, roleRight.RCreate, roleRight.RRead, roleRight.RUpdate, roleRight.RDelete,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return roleRight, nil
}
