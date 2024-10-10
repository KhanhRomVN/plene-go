package repository

import (
	"database/sql"
	"pleno-go/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password, full_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, user.Username, user.Email, user.Password, user.FullName).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, full_name, created_at, updated_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateRefreshToken(userID int64, refreshToken string) error {
	query := `UPDATE users SET refresh_token = $1 WHERE id = $2`
	_, err := r.db.Exec(query, refreshToken, userID)
	return err
}
