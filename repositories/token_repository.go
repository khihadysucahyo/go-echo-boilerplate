package repositories

import (
	"time"

	"github.com/khihadysucahyo/go-echo-boilerplate/config"

	"github.com/khihadysucahyo/go-echo-boilerplate/models"

	"github.com/dgrijalva/jwt-go"
)

const ExpireCount = 2
const ExpireRefreshCount = 168

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ServiceWrapper interface {
	CreateAccessToken(user *models.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *models.User) (t string, err error)
}

type Service struct {
	config *config.Config
}

func NewTokenRepository(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
	}
}

func (tokenService *Service) CreateAccessToken(user *models.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Hour * ExpireCount).Unix()
	claims := &JwtCustomClaims{
		user.Name,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(tokenService.config.Auth.AccessSecret))
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}

func (tokenService *Service) CreateRefreshToken(user *models.User) (t string, err error) {
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * ExpireRefreshCount).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := refreshToken.SignedString([]byte(tokenService.config.Auth.RefreshSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}
