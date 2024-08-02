package soap

import (
	"fmt"
	"os"
)

// Retrieve a .xml file template from the specified path
func GetTemplate(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("File contents:", string(file))
	return string(file)
}
