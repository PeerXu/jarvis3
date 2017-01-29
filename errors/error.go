package errors

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	ErrorInvalidRequest = "invalid_request"
	ErrorNotFound       = "not_found"
	ErrorServerError    = "server_error"
	ErrorAccessDenied   = "access_denied"
)

var NotImplementNowError = fmt.Errorf("not implement now")

type JarvisError struct {
	Module      string `json:"module"`
	Type        string `json:"error"`
	Description string `json:"error_description,omitempty"`
	Internal    error  `json:"-"`
}

func (e JarvisError) Error() string {
	return e.Type
}

func NewJarvisError(mdl string, typ string, dsc string, err error) JarvisError {
	return JarvisError{
		Module:      mdl,
		Type:        typ,
		Description: dsc,
		Internal:    err,
	}
}

func DecodeJarvisError(reader io.Reader) (JarvisError, bool, error) {
	var jerr JarvisError
	err := json.NewDecoder(reader).Decode(&jerr)
	if err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); !ok || ute.Value != "array" {
			return jerr, false, err
		}
		return jerr, false, nil
	}

	if jerr.Type != "" {
		return jerr, true, nil
	}

	return jerr, false, nil
}
