package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadAndAssignResponseBody(res *http.Response) (io.Reader, error) {
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if err != io.EOF || len(buf) != 0 {
			return nil, err
		}
	}

	res.Body = ioutil.NopCloser(bytes.NewReader(buf))
	return bytes.NewReader(buf), nil
}
