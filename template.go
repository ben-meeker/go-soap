package soap

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

// Retrieve a .xml file template from the specified path
func GetTemplate(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// Ensure that all parameter placeholders have been filled with values
func VerifyParameters(requestXML string) error {
	// Get unfilled parameters
	reg := regexp.MustCompile("{.+}")
	matches := reg.FindAllString(requestXML, -1)
	var errString string
	if len(matches) > 0 {
		for _, val := range matches {
			errString = errString + val + " "
		}
		return errors.New("The following placeholders are missing values: " + errString)
	}

	return nil
}

// Fill a template with fields from given struct
func FillTemplate(template string, parameters any) string {
	values := reflect.ValueOf(parameters)
	types := values.Type()
	requestBody := template
	for i := 0; i < values.NumField(); i++ {
		requestBody = strings.ReplaceAll(requestBody, "{"+types.Field(i).Name+"}", fmt.Sprint(values.Field(i).Interface()))
	}
	return requestBody
}
