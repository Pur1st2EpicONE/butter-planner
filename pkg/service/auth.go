package service

import (
	"os"
	"strconv"
	"time"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 12 * time.Hour

type AuthService struct {
	storer repository.Storer
}

func NewAuthService(storer repository.Storer) *AuthService {
	return &AuthService{storer: storer}
}

func (s *AuthService) CreateUser(user butterplanner.User) (int, error) {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	return s.storer.CreateUser(user)
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func (s *AuthService) GetUserId(user butterplanner.LoginPassword) (int, error) {
	return s.storer.GetUserId(user)
}

func (s *AuthService) CreateToken(userId int) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   strconv.Itoa(userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
