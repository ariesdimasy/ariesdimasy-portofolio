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
	GetUserID(id uint) (*models.User, error)
	GetUserEmail(email string) (*models.User, error)
	GetUserUsername(username string) (*models.User, error)
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

func (ur userRepository) GetUserID(id uint) (*models.User, error) {
	var user models.User
	return &user, ur.db.First(&user, id).Error
}

func (ur userRepository) GetUserEmail(email string) (*models.User, error) {
	var user models.User
	return &user, ur.db.Where("email = ?", email).First(&user).Error
}

func (ur userRepository) GetUserUsername(username string) (*models.User, error) {
	var user models.User
	return &user, ur.db.Where("username = ?", username).First(&user).Error
}

func (ur userRepository) GetAllUsers(query models.UserQuery) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	errCount := ur.db.Model(&models.User{})

	if query.Name != "" {
		errCount = errCount.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.LoginBy != "" {
		errCount = errCount.Where("login_by = ?", query.LoginBy)
	}

	errCount.Count(&total)

	if errCount.Error != nil {
		return nil, 0, errCount.Error
	}

	offset := (query.Page - 1) * query.Limit
	dbData := ur.db.Model(&models.User{}).Order("created_at desc").Offset(offset).Limit(query.Limit)

	if query.Name != "" {
		dbData = dbData.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.LoginBy != "" {
		dbData = dbData.Where("login_by = ?", query.LoginBy)
	}

	if err := dbData.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
