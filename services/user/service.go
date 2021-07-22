package user

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/requests"

	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Register(request *requests.RegisterRequest) error
}

type Service struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
