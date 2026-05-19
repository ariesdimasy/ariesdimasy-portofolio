package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(query models.UserQuery) ([]models.User, int64, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{repo: userRepository}
}

func (us userService) Register(user *models.User) error {
	return us.repo.Register(user)
}

func (us userService) Login(email, password string) (*models.User, error) {
	return us.repo.Login(email, password)
}

func (us userService) CreateUser(user *models.User) error {
	return us.repo.CreateUser(user)
}

func (us userService) UpdateUser(user *models.User) error {
	return us.repo.UpdateUser(user)
}

func (us userService) DeleteUser(user *models.User) error {
	return us.repo.DeleteUser(user)
}

func (us userService) GetUserByID(id uint) (*models.User, error) {
	return us.repo.GetUserID(id)
}

func (us userService) GetAllUsers(query models.UserQuery) ([]models.User, int64, error) {
	return us.repo.GetAllUsers(query)
}
