package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) error
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(query models.UserQuery) ([]models.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur userRepository) Register(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur userRepository) Login(user *models.User) error {
	return ur.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
}

func (ur userRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur userRepository) UpdateUser(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur userRepository) DeleteUser(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	return &user, ur.db.First(&user, id).Error
}

func (ur userRepository) GetAllUsers(query models.UserQuery) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	errCount := ur.db.Model(&models.User{})

	if query.Name != "" {
		errCount = errCount.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Email != "" {
		errCount = errCount.Where("email LIKE ?", "%"+query.Email+"%")
	}

	if query.LoginBy != "" {
		errCount = errCount.Where("login_by = ?", query.LoginBy)
	}

	errCount.Count(&total)

	if errCount.Error != nil {
		return nil, 0, errCount.Error
	}

	offset := (query.Page - 1) * query.Limit
	errData := ur.db.Order("created_at desc").Offset(offset).Limit(query.Limit).Find(&users)

	if query.Name != "" {
		errData = errData.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Email != "" {
		errData = errData.Where("email LIKE ?", "%"+query.Email+"%")
	}

	if query.LoginBy != "" {
		errData = errData.Where("login_by = ?", query.LoginBy)
	}

	errData.Find(&users)

	if errData.Error != nil {
		return nil, 0, errData.Error
	}

	return users, total, nil
}
