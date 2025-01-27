package repository

import (
	"future-path/entity"
	"future-path/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateOrUpdateUser(user *entity.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("email_user = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) CreateOrUpdateUser(user *entity.User) error {
	err := u.db.Debug().Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
