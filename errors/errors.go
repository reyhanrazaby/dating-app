package errors

type AuthError struct {
	Reason string
}

func (e AuthError) Error() string {
	return e.Reason
}

type SignUpError struct {
	Reason string
}

func (e SignUpError) Error() string {
	return e.Reason
}
