package service

import (
	"errors"
	"future-path/entity"
	"future-path/internal/repository"
)

type IOAuthService interface {
	HandleGoogleLogin(email, name, token string) (*entity.User, error)
}

type OAuthService struct {
	UserRepository repository.IUserRepository
}

func NewOAuthService(repo repository.IUserRepository) IOAuthService {
	return &OAuthService{repo}
}

func (o *OAuthService) HandleGoogleLogin(email, name, token string) (*entity.User, error) {
	user, err := o.UserRepository.FindUserByEmail(email)
	if err != nil {
		user = &entity.User{
			Nama_User:  name,
			Email_User: email,
			Token:      token,
			RoleID:     2,
		}
	} else {
		user.Token = token
	}

	err = o.UserRepository.CreateOrUpdateUser(user)
	if err != nil {
		return nil, errors.New("failed to save user")
	}

	return user, nil
}
