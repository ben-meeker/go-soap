package main

import (
	"fmt"

	"github.com/ben-meeker/soap"
)

func main() {
	// Set endpoint URL
	const url string = "https://www.w3schools.com/xml/tempconvert.asmx"

	// Retrieve farenheit_to_celsius.xml template
	template := soap.GetTemplate("fahrenheit_to_celsius.xml")

	// Set values based on template requirements
	// In this case, the first and only required value is temperature
	values := []any{75}

	// Set headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	// Make SOAP call with standard content/type
	res, err := soap.SoapCall(url, headers, template, values)
	if err != nil {
		// Handle error
		panic(err)
	}

	xmlRes, err := ParseXMLResponse(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(xmlRes)
}
