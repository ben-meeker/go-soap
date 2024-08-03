package soap

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Retrieve a .xml file template from the specified path
func GetTemplate(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(file)
}

// Ensure that the number of parameters required matches that of the template
func VerifyParameterList(template string, parameters []any) error {
	// Read length of parameters
	params := len(parameters)
	// Get number of parameters in template
	reg := regexp.MustCompile(">%v<")
	matches := reg.FindAllStringIndex(template, -1)
	templateParams := len(matches)

	if params > templateParams {
		return errors.New("Too many values provided to template")
	} else if templateParams > params {
		return errors.New("Not enough values passed to template")
	}

	return nil
}
