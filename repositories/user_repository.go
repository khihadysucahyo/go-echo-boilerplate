package repositories

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/models"
	"github.com/khihadysucahyo/go-echo-boilerplate/requests"
	"github.com/khihadysucahyo/go-echo-boilerplate/server/builders"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

type UserRepositoryQ interface {
	GetUserByEmail(user *models.User, email string)
	Register(request *requests.RegisterRequest) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) GetUserByEmail(user *models.User, email string) {
	userRepository.DB.Where("email = ?", email).Find(user)
}

func (userRepository *UserRepository) Register(request *requests.RegisterRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := builders.NewUserBuilder().SetEmail(request.Email).
		SetName(request.Name).
		SetPassword(string(encryptedPassword)).
		Build()

	return userRepository.DB.Create(&user).Error
}
