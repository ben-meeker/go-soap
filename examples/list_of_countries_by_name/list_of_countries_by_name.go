package main

import (
	"fmt"

	"github.com/ben-meeker/soap" // Will be go-soap on non-local imports
)

// Create struct to fill request body with
// Fields must be exported or available in package, or you will get an error
// populating them into the template
type ListOfContinentsByName struct {
	Continents []Continent
}

type Continent struct {
	Code string
	Name string
}

func main() {
	// Set endpoint URL
	const url string = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

	// Retrieve farenheit_to_celsius.xml template
	template, err := soap.GetTemplate("list_of_countries_by_name.xml")
	if err != nil {
		fmt.Println(err)
	}

	// Set values based on template requirements
	// In this case, no values are required
	listOfContinents := ListOfContinentsByName{}

	// Set headers
	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	// Make SOAP call with standard content/type
	res, err := soap.SoapCall(url, headers, template, listOfContinents)
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

	// Get continents from XML response                // Reference Children          // Key
	continents := xmlRes.Body.Children["soap:Envelope"].Children["soap:Body"].Children["m:ListOfContinentsByNameResponse"].Children["m:ListOfContinentsByNameResult"].Children

	// Get values for each continent
	for key, _ := range continents {
		continentData := Continent{
			Code: continents[key].Children["m:sCode"].Value,
			Name: continents[key].Children["m:sName"].Value,
		}
		// Append to struct
		listOfContinents.Continents = append(listOfContinents.Continents, continentData)
	}

	// Output struct
	fmt.Println(listOfContinents)
}
