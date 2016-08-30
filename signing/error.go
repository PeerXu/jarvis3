package signing

const (
	errorInvalidRequest = "invalid_request"
	errorNotFound       = "not_found"
	errorServerError    = "server_error"
	errorAccessDenied   = "access_denied"
)

type signError struct {
	Type        string `json:"error"`
	Description string `json:"error_description,omitempty"`
	Internal    error  `json:"-"`
}

func (e *signError) Error() string {
	return e.Type
}

func newSignError(typ string, desc string, err error) *signError {
	return &signError{Type: typ, Description: desc, Internal: err}
}

func convertSignError(x interface{}) (*signError, bool) {
	err, ok := x.(*signError)
	return err, ok
}
