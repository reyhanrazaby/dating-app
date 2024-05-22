package datasource

import (
	"github.com/reyhanrazaby/dating-app/entity"
)

type Repo interface {
	SaveUserAuth(entity.UserAuth) error
	GetUserAuthByEmail(string) *entity.UserAuth
}
