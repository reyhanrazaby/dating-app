package login

import "github.com/reyhanrazaby/dating-app/errors"

type LoginService interface {
	Login(string, string) error
}

func NewService() LoginService {
	return &_service{}
}

type _service struct {
}

func (s *_service) Login(email, password string) error {
	if email == "reyhan@mail.com" && password == "1234" {
		return nil
	}
	return errors.AuthError{Reason: "Wrong email or password"}
}
