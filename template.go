package soap

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Retrieve a .xml file template from the specified path
func GetTemplate(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// Ensure that the number of parameters required matches that of the template
func VerifyParameterList(template string, parameters []any) error {
	// Read length of parameters
	params := len(parameters)
	// Get number of parameters in template
	reg := regexp.MustCompile(">%v<")
	matches := reg.FindAllStringIndex(template, -1)
	templateParams := len(matches)

	if params != templateParams {
		return errors.New("Invalid parameter count. Template wanted " + fmt.Sprint(templateParams) + ", got " + fmt.Sprint(params))
	}

	return nil
}
