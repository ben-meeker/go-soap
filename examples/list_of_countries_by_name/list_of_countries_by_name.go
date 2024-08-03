package main

import (
	"fmt"

	"github.com/ben-meeker/soap"
)

func main() {
	// Set endpoint URL
	const url string = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

	// Retrieve farenheit_to_celsius.xml template
	template := soap.GetTemplate("list_of_countries_by_name.xml")

	// Set values based on template requirements
	// In this case, no values are required
	values := []any{}

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

	// Get values from XML response   // Key        // Type assertion           // To the moon! ------------------------------------------------------------------------------------------------------------------------------------------------Third country in list------------------------------------------------------> You made it! :)
	fmt.Println(xmlRes.Body.Contents["soap:Envelope"].(soap.XMLObject).Contents["soap:Body"].(soap.XMLObject).Contents["m:ListOfContinentsByNameResponse"].(soap.XMLObject).Contents["m:ListOfContinentsByNameResult"].(soap.XMLObject).Contents["m:tContinent_3"].(soap.XMLObject).Contents["m:sCode"].(soap.XMLObject).Contents["value"])
	fmt.Println(xmlRes.Body.Contents["soap:Envelope"].(soap.XMLObject).Contents["soap:Body"].(soap.XMLObject).Contents["m:ListOfContinentsByNameResponse"].(soap.XMLObject).Contents["m:ListOfContinentsByNameResult"].(soap.XMLObject).Contents["m:tContinent_3"].(soap.XMLObject).Contents["m:sName"].(soap.XMLObject).Contents["value"])
}
