package computing

const (
	errorInvalidRequest = "invalid_request"
	errorNotFound       = "not_found"
	errorServerError    = "server_error"
)

type computeError struct {
	Type        string `json:"error"`
	Description string `json:"error_description,omitempty"`
	Internal    error  `json:"-"`
}

func (e *computeError) Error() string {
	return e.Type
}

func newComputeError(typ string, desc string, err error) *computeError {
	return &computeError{Type: typ, Description: desc, Internal: err}
}
