package main

import (
	"fmt"

	"github.com/ben-meeker/soap" // Will be go-soap on non-local imports
)

func main() {
	// Set endpoint URL
	const url string = "https://www.w3schools.com/xml/tempconvert.asmx"

	// Retrieve farenheit_to_celsius.xml template
	template, err := soap.GetTemplate("fahrenheit_to_celsius.xml")
	if err != nil {
		fmt.Println(err)
	}

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

	// Parse XML response
	xmlRes, err := soap.ParseXMLResponse(res)
	if err != nil {
		fmt.Println(err)
	}

	// View structure of XML response
	fmt.Println(xmlRes.Structure)

	// Get value from XML response   // Key        // Type assertion           // To the moon! --------------------------------------------------------------------------------------------------------------------------------> You made it! :)
	fmt.Println(xmlRes.Body.Contents["soap:Envelope"].(soap.XMLObject).Contents["soap:Body"].(soap.XMLObject).Contents["FahrenheitToCelsiusResponse"].(soap.XMLObject).Contents["FahrenheitToCelsiusResult"].(soap.XMLObject).Contents["value"])
}
