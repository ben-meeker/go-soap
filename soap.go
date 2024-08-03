package soap

import (
	"bytes"
	"fmt"
	"net/http"
)

// Send a SOAP http request
func SoapCall(url string, headers map[string]string, template string, values []any) (*http.Response, error) {
	// Start http client
	client := &http.Client{}

	// Create request body from template and values
	requestBody := fmt.Sprintf(template, values...)

	// Create request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return nil, err
	}

	// Add headers
	for label, value := range headers {
		request.Header.Add(label, value)
	}

	// Make request
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
