package services

import (
	models "github.com/ladarat/ClinicVue/go/models"
	repository "github.com/ladarat/ClinicVue/go/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService is service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (ps *userService) GetAll() ([]models.User, error) {

	return nil, nil
}
