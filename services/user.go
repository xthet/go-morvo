package services

import (
	"fmt"
	"os"

	"github.com/xthet/go-morvo/models"
	"github.com/xthet/go-morvo/services/auth"
	"github.com/xthet/go-morvo/types"
)

type UserService struct {
	user_collection *models.UserCollection
}

func User(user_collection *models.UserCollection) *UserService {
	return &UserService{
		user_collection: user_collection,
	}
}

func (s UserService) RegisterUser(payload types.RegisterUserPayload) (*types.RegisterUserResponse, error) {
	// check if email exists
	_, err := s.user_collection.GetUserByEmail(payload.Email)
	if err == nil {
		// means user already exists
		return nil, fmt.Errorf("user with email: %s already exists", payload.Email)
	}

	hashed_password, err := auth.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	err = s.user_collection.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashed_password,
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterUserResponse{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}, nil
}

func (s UserService) LoginUser(user *types.LoginUserPayload) (string, error) {
	u, err := s.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		return "", fmt.Errorf("invalid email or password")
	}

	secret := os.Getenv("JWT_SECRET")
	token, err := auth.CreateJWT([]byte(secret), u.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s UserService) GetUserByEmail(email string) (*types.User, error) {
	user, err := s.user_collection.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}
