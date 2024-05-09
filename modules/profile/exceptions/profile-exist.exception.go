package exceptions

type ProfileExistsError struct {
	Message string
}

func (e *ProfileExistsError) Error() string {
	return e.Message
}

func NewProfileExistsError(message string) error {
	return &ProfileExistsError{Message: message}
}
