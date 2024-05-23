package datasource

import (
	"github.com/reyhanrazaby/dating-app/entity"
)

type Repo interface {
	SaveUserAuth(userAuth entity.UserAuth) error
	GetUserAuthByEmail(email string) *entity.UserAuth
}
