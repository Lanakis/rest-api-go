package exceptions

type UserNotFoundError struct {
	Message string
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}

func NewUserNotFoundError(message string) error {
	return &UserNotFoundError{Message: message}
}
