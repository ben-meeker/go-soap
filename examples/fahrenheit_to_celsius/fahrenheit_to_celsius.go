package main

import (
	"fmt"
	"strconv"

	"github.com/ben-meeker/soap" // Will be go-soap on non-local imports
)

// Create struct to fill request body with
// Fields must be exported or available in package, or you will get an error
// populating them into the template
type FahrenheitToCelsius struct {
	Fahrenheit float64
	Celsius    float64
}

func main() {
	// Set endpoint URL
	const url string = "https://www.w3schools.com/xml/tempconvert.asmx"

	// Retrieve farenheit_to_celsius.xml template
	template, err := soap.GetTemplate("fahrenheit_to_celsius.xml")
	if err != nil {
		fmt.Println(err)
	}

	// Set values based on template requirements
	// In this case, the only required value is fahrenheit
	values := FahrenheitToCelsius{
		Fahrenheit: 75,
	}

	// Set headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	// Make SOAP call with standard content/type
	res, err := soap.SoapCall(url, headers, template, values)
	if err != nil {
		// Handle error
		panic(err)
	}

	// Parse XML response
	xmlRes, err := soap.ParseXMLResponse(res)
	if err != nil {
		fmt.Println(err)
	}

	// View structure of XML response
	fmt.Println(xmlRes.Structure)

	// Get value from XML response                         // Reference Children          // Key                                    // Value will always be a string
	celsiusString := xmlRes.Body.Children["soap:Envelope"].Children["soap:Body"].Children["FahrenheitToCelsiusResponse"].Children["FahrenheitToCelsiusResult"].Value
	// Convert value type to match struct
	values.Celsius, err = strconv.ParseFloat(celsiusString, 64)

	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(values)
}
