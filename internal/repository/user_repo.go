package repository

import (
	models "vibex-api/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByID(id int64) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) FindUserByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Status").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Status").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Status").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) DeleteUser(user *models.User) error {
	return r.db.Delete(user).Error
}
