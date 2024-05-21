package signup

import (
	"fmt"

	"github.com/reyhanrazaby/dating-app/errors"
)

type SignUpService interface {
	SignUp(request) error
}

func NewService() SignUpService {
	return &_service{}
}

type _service struct {
}

func (s *_service) SignUp(req request) error {
	if req.Gender != "M" && req.Gender != "F" {
		return errors.SignUpError{Reason: fmt.Sprintf("Gender %s is not valid", req.Gender)}
	}
	return nil
}
