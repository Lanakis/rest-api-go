package exceptions

type UserExistsError struct {
	Message string
}

func (e *UserExistsError) Error() string {
	return e.Message
}

func NewUserExistsError(message string) error {
	return &UserExistsError{Message: message}
}
