package repository

import (
	"context"
	"ebook-api/db"
	"ebook-api/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	query := `INSERT INTO "user" (name, email) VALUES ($1, $2) RETURNING user_id`
	err := db.Conn.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.UserID)
	return user, err
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := db.Conn.Query(ctx, `SELECT user_id, name, email FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
