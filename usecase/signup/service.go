package signup

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reyhanrazaby/dating-app/datasource"
	"github.com/reyhanrazaby/dating-app/entity"
	"github.com/reyhanrazaby/dating-app/errors"
	"github.com/reyhanrazaby/dating-app/util"
)

type SignUpService interface {
	SignUp(request) error
}

func NewService(repo datasource.Repo) SignUpService {
	return &_service{repo}
}

type _service struct {
	Repo datasource.Repo
}

func (s *_service) SignUp(req request) error {
	if req.Gender != "M" && req.Gender != "F" {
		return errors.SignUpError{Reason: fmt.Sprintf("Gender %s is not valid", req.Gender)}
	}

	dateLayout := "02-01-2006"
	dateBirth, err := time.Parse(dateLayout, req.DateBirth)
	if err != nil {
		return errors.SignUpError{Reason: fmt.Sprintf("Date Birth %s is not valid", req.DateBirth)}
	}

	userProfile := entity.UserProfile{
		Id:          uuid.New(),
		FullName:    req.FullName,
		Gender:      rune(req.Gender[0]),
		Bio:         req.Bio,
		DateBirth:   dateBirth,
		LocationLat: req.LocationLat,
		LocationLng: req.LocationLng,
	}

	salt := uuid.New().String()
	hashedPassword, err := util.HashPassword(req.Password + salt)
	if err != nil {
		return err
	}
	userAuth := entity.UserAuth{
		User:     userProfile,
		Email:    req.Email,
		Password: hashedPassword,
		Salt:     salt,
	}

	return s.Repo.SaveUserAuth(userAuth)
}
