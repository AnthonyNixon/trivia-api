package httperr

// New returns an error that formats as the given text.
func New(status int, description string) HttpErr {
	return &error{status, description}
}

type error struct {
	status      int
	description string
}

func (e *error) Error() string {
	return e.description
}

func (e *error) Description() string {
	return e.description
}

func (e *error) StatusCode() int {
	return e.status
}

type HttpErr interface {
	StatusCode() int
	Description() string
}
