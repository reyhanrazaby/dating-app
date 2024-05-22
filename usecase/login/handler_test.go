package login

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/reyhanrazaby/dating-app/entity"
	"github.com/reyhanrazaby/dating-app/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var router *gin.Engine
var srvMock = new(serviceMock)

func TestMain(m *testing.M) {
	router = gin.Default()
	router.POST(Path, Handler())
	service = srvMock
	m.Run()
}

func Test_StatusOK(t *testing.T) {
	srvMock.On("Login", "reyraz@mail.id", "1234").Return(entity.UserProfile{}, nil)

	w := httptest.NewRecorder()

	body := []byte(`
		{
			"email": "reyraz@mail.id",
			"password": "1234"
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_EmptyEmail(t *testing.T) {
	w := httptest.NewRecorder()

	body := []byte(`
		{
			"email": "",
			"password": "1234"
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_EmptyPassword(t *testing.T) {
	w := httptest.NewRecorder()

	body := []byte(`
		{
			"email": "reyraz@mail.com",
			"password": ""
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_AuthError(t *testing.T) {
	w := httptest.NewRecorder()

	srvMock.On("Login", "rara@mail.id", "123").Return(entity.UserProfile{}, errors.AuthError{})

	body := []byte(`
		{
			"email": "rara@mail.id",
			"password": "123"
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

type serviceMock struct{ mock.Mock }

func (s *serviceMock) Login(email, password string) (entity.UserProfile, error) {
	args := s.Called(email, password)
	var err error = nil
	userProfile := entity.UserProfile{}

	if args.Get(0) != nil {
		userProfile = args.Get(0).(entity.UserProfile)
	}
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return userProfile, err
}
