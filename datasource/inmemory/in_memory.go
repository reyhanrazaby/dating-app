package inmemory

import "github.com/reyhanrazaby/dating-app/entity"

var singleton *InMemory

type InMemory struct {
	userAuthsByEmail map[string]entity.UserAuth
}

func GetInstance() *InMemory {
	if singleton == nil {
		singleton = &InMemory{make(map[string]entity.UserAuth)}
	}
	return singleton
}

func (r *InMemory) SaveUserAuth(userAuth entity.UserAuth) error {
	r.userAuthsByEmail[userAuth.Email] = userAuth
	return nil
}

func (r *InMemory) GetUserAuthByEmail(email string) *entity.UserAuth {
	val, ok := r.userAuthsByEmail[email]
	if ok {
		return &val
	} else {
		return nil
	}
}
