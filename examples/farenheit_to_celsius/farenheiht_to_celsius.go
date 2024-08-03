package main

import (
	"fmt"
	"io"

	"github.com/ben-meeker/soap"
)

func main() {
	// Set endpoint URL
	const url string = "https://www.w3schools.com/xml/tempconvert.asmx"

	// Retrieve farenheit_to_celsius.xml template
	template := soap.GetTemplate("farenheit_to_celsius.xml")

	// Make SOAP call with standard content/type
	res, err := soap.SoapCall(url, "application/soap+xml; charset=utf-8", template)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	// Parse response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Do something with response
	fmt.Println(body)
}
