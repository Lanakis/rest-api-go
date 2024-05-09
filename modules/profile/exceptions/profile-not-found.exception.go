package exceptions

type ProfileNotFoundError struct {
	Message string
}

func (e *ProfileNotFoundError) Error() string {
	return e.Message
}

func NewProfileNotFoundError(message string) error {
	return &ProfileNotFoundError{Message: message}
}
