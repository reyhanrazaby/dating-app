package login

import (
	"github.com/reyhanrazaby/dating-app/datasource"
	"github.com/reyhanrazaby/dating-app/entity"
	"github.com/reyhanrazaby/dating-app/errors"
	"github.com/reyhanrazaby/dating-app/util"
)

type LoginService interface {
	Login(email, password string) (entity.UserProfile, error)
}

func NewService(repo datasource.Repo) LoginService {
	return &_service{repo}
}

type _service struct {
	Repo datasource.Repo
}

func (s *_service) Login(email, password string) (entity.UserProfile, error) {
	userAuth := s.Repo.GetUserAuthByEmail(email)
	if userAuth == nil {
		return entity.UserProfile{}, errors.AuthError{Reason: "Email is not registered"}
	}

	if !util.CheckPasswordHash(password+userAuth.Salt, userAuth.Password) {
		return entity.UserProfile{}, errors.AuthError{Reason: "Wrong password"}
	}

	return userAuth.User, nil
}
