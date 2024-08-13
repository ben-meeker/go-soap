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
	requestBody := template
	reg := regexp.MustCompile("{.+}")
	matches := reg.FindAllString(requestBody, -1)
	for _, placeholder := range matches {
		var value reflect.Value
		values := reflect.ValueOf(parameters)
		if strings.Contains(placeholder, ".") {
			noBrackets := strings.Trim(placeholder, "{")
			noBrackets = strings.Trim(noBrackets, "}")
			split := strings.Split(noBrackets, ".")
			for i := range split {
				if i == len(split)-1 {
					requestBody = strings.ReplaceAll(requestBody, placeholder, fmt.Sprint(value.FieldByName(split[i])))
				} else {
					value = values.FieldByName(split[i])
				}
			}
		} else {
			noBrackets := strings.Trim(placeholder, "{")
			noBrackets = strings.Trim(noBrackets, "}")
			requestBody = strings.ReplaceAll(requestBody, placeholder, fmt.Sprint(values.FieldByName(noBrackets)))
		}
	}

	return requestBody
}
