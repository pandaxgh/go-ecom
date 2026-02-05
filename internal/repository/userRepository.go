package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pandaxgh/ecom-backend/internal/model"
)

type userRepository struct {
	db *pgxpool.Pool
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {

	query := `
	INSERT INTO users (first_name,last_name,user_name,password,email,role)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(ctx, query, user.FirstName, user.LastName, user.UserName, user.Password, user.Email, user.Role).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)

}
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	query := `
	SELECT id, first_name,last_name,user_name,email,role, password, created_at,deleted_at,is_verified,updated_at,is_deleted
	FROM users
	WHERE email = $1 AND is_deleted = false
	`
	var user model.User

	err := r.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.DeletedAt, &user.IsVerified, &user.UpdatedAt, &user.IsDeleted)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {

	query := ` 
	SELECT id, first_name,last_name,user_name,email,role, password, created_at,deleted_at,is_verified,updated_at, is_deleted
	FROM USER 
	WHERE id = $1 AND is_deleted = false
	`
	var user *model.User

	err := r.db.QueryRow(ctx, query, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.DeletedAt, &user.IsVerified, &user.UpdatedAt, &user.IsDeleted)

	if err != nil {
		return nil, err
	}
	return user, nil

}
func (r *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {

	query := `
	SELECT id,first_name,last_name,user_name,email,role,password,created_at,deleted_at,is_verified,updated_at,is_deleted
	FROM USER 
	`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Role, &user.Password, &user.CreatedAt, &user.DeletedAt, &user.IsVerified, &user.UpdatedAt, &user.IsDeleted)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
