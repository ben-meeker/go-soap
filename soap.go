package soap

import (
	"bytes"
	"fmt"
	"net/http"
)

// Send a SOAP http request
func SoapCall(url string, contentType string, template string, values ...any) (*http.Response, error) {
	request := fmt.Sprintf(template, values...)

	res, err := http.Post(url, contentType, bytes.NewBuffer([]byte(request)))
	if err != nil {
		return res, err
	}

	return res, nil
}
