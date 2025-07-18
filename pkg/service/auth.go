package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 12 * time.Hour

type AuthService struct {
	storer repository.Storer
}

func NewAuthService(storer repository.Storer) *AuthService {
	return &AuthService{storer: storer}
}

func (a *AuthService) CreateUser(user butterplanner.User) (int, error) {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	return a.storer.CreateUser(user)
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func (a *AuthService) GetUserId(user butterplanner.LoginPassword) (int, error) {
	return a.storer.GetUserId(user)
}

func (a *AuthService) CreateToken(userId int) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func (a *AuthService) ParseToken(tokenString string) (int, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, errors.New("invalid user id in token subject")
	}

	return userId, nil
}
