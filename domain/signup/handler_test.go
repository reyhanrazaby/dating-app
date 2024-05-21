package signup

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	expected := request{
		FullName: "Rey Raz",
		Gender:   "M",
		Email:    "reyraz@mail.id",
		Password: "1234",
	}
	srvMock.On("SignUp", expected).Return(nil)

	w := httptest.NewRecorder()

	body := []byte(`
		{
			"full_name": "Rey Raz",
			"gender": "M",
			"email": "reyraz@mail.id",
			"password": "1234"
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_EmptyBody(t *testing.T) {
	w := httptest.NewRecorder()

	body := []byte(`
		{
			"full_name": "",
			"gender": "",
			"email": "",
			"password": ""
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_SignUpError(t *testing.T) {
	w := httptest.NewRecorder()

	srvMock.On("SignUp", mock.Anything).Return(errors.SignUpError{})

	body := []byte(`
		{
			"full_name": "Rara",
			"gender": "F",
			"email": "rara@mail.id",
			"password": "123"
		}
	`)
	req, _ := http.NewRequest("POST", Path, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

type serviceMock struct{ mock.Mock }

func (s *serviceMock) SignUp(req request) error {
	args := s.Called(req)
	var err error = nil

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
