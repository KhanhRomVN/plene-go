package services

import (
	"errors"
	"pleno-go/internal/models"
	"pleno-go/internal/repository"
	"pleno-go/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*models.User, string, string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, "", "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", "", errors.New("invalid credentials")
	}

	accessToken, err := generateAccessToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := generateRefreshToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}

	err = s.userRepo.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func generateAccessToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("your_access_token_secret"))
}

func generateRefreshToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte("your_refresh_token_secret"))
}
