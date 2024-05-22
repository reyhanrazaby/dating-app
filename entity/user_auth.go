package entity

type UserAuth struct {
	User     UserProfile
	Email    string
	Password string
	Salt     string
}
