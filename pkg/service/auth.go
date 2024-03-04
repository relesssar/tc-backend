package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	tc "tc_kaztranscom_backend_go"
	"tc_kaztranscom_backend_go/pkg/repository"
	"time"
)

const (
	salt       = "jfufhTGre43m!jHBP276cEWKt"
	signingKey = "jfnvhYJUjfy649!jerbglkN#BWl2w"

	// Время жизни Токена
	tokenTTL = 48 * time.Hour

	//определяет приложение, из которого отправляется токен.
	issuer = "Total Control +77773785631,7773785631@mail.ru"
)

type tokenClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user tc.User) (string, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	logrus.Info("GenerateToken: start")
	user, err := s.repo.GetUser(email, s.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	logrus.Infof("user: %v", user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		user.Id,
		jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type ")
	}
	return claims.UserId, nil
}
func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
