package service

import (
	"errors"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
	"future-path/pkg/bcrypt"
	"future-path/pkg/jwt"
)

type IUserService interface {
	Register(param model.UserRegister) error
	Login(param model.UserLogin) (model.UserLoginResponses, error)
	GetUser(param model.UserParam) (entity.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserService {
	return &UserService{
		bcrypt:         bcrypt,
		UserRepository: userRepository,
		jwtAuth:        jwt,
	}
}

func (us *UserService) Register(param model.UserRegister) error {
	checkUser, err := us.UserRepository.FindUserByEmail(param.Email_User)
	if err == nil && checkUser != nil {
		return errors.New("email already registered")
	}

	hashPassword, err := us.bcrypt.GenerateFromPassword(param.Password_User)
	if err != nil {
		return err
	}

	param.Password_User = hashPassword

	user := entity.User{
		Nama_User:     param.Nama_User,
		Email_User:    param.Email_User,
		Password_User: hashPassword,
		RoleID:        2,
	}

	_, err = us.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}

func (us *UserService) Login(param model.UserLogin) (model.UserLoginResponses, error) {
	var result model.UserLoginResponses

	user, err := us.UserRepository.GetUser(model.UserParam{
		Email_User: param.Email_User,
	})
	if err != nil {
		return result, err
	}

	err = us.bcrypt.CompareAndHashPassword(user.Password_User, param.Password_User)
	if err != nil {
		return result, err
	}

	token, err := us.jwtAuth.CreateToken(user.ID_User)
	if err != nil {
		return result, errors.New("failed to create jwt")
	}

	result.Token = token
	result.RoleID = user.RoleID

	return result, nil
}

func (us *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return us.UserRepository.GetUser(param)
}
