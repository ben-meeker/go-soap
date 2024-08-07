# Go-Soap

Go-Soap is an open-source library for handling SOAP requests and responses in Go. 

## Features

- [x] Easy to understand functions for making SOAP requests
- [x] Ability to add headers
- [x] Able to parse responses into dynamic objects for easy data retrieval

## Installation

You can install Soap using Go modules:

```bash
go get github.com/ben-meeker/go-soap
```

## Getting Started

Please see the `examples` folder for working builds that use this library

1. In your Go code, import the Soap package:
    ```go
    import (
        "github.com/ben-meeker/go-soap"
    )
    ```

2. Create a struct to house relevant data
   ```go
   type FahrenheitToCelsius struct {
	   Fahrenheit float64
	   Celsius    float64
   }
   ```

3. Create a new XML file template using the struct field name inside curly braces as variable placeholders
    ```xml
   <?xml version="1.0" encoding="utf-8"?>
   <soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
    <soap12:Body>
      <FahrenheitToCelsius xmlns="https://www.w3schools.com/xml/">
         <Fahrenheit>{Fahrenheit}</Fahrenheit>
      </FahrenheitToCelsius>
    </soap12:Body>
   </soap12:Envelope>
    ```

4. Set your endpoint URL
   ```go
   const url string = "https://www.w3schools.com/xml/tempconvert.asmx"
   ```

5. Use the `GetTemplate` function to load your template into Go
   ```go
   template, err := soap.GetTemplate("fahrenheit_to_celsius.xml")
   if err != nil {
	   fmt.Println(err)
	}
   ```
   This will return an error if unable to fetch the template provided


6. Create headers in a type `map[string]string` 
   ```go
   headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"
   ```

7. Use the `SoapCall` function to make the request
   ```go
   res, err := soap.SoapCall(url, headers, template, values)
	if err != nil {
		panic(err)
	}
   ```
   This will look at the number of paramters provided, and compare it to the number of parameters detected in your template designated by the `%v` placeholder. If they do not match, or there is an error making the request, the function will return an error.

8. Use the `ParseXMLResponse` function to create a readable object from the XML response
   ```go
   xmlRes, err := soap.ParseXMLResponse(res)
	if err != nil {
		fmt.Println(err)
	}
   ```
   This will parse the data into a series of `XMLObject` that contain a `Contents` field of type `map[string]any` . This may contain another `XMLObject` if nesting has occured, or will have a value designated by the key `value`

9. To view the structure of the response, print the `Structure` field of your `XMLResponse` object
   ```go
    fmt.Println(xmlRes.Structure)
   ```

10. To access a value from the object, use the syntax below to drill down to the location, and access it using `.Value` this will always be a `string`
   ```go
   celsiusString := xmlRes.Body.Children["soap:Envelope"].Children["soap:Body"].Children["FahrenheitToCelsiusResponse"].Children["FahrenheitToCelsiusResult"].Value
   ```

11. Use the `strconv` package to convert your value to a type that matches your struct and assign it to a field
   ```go
   values.Celsius, err = strconv.ParseFloat(celsiusString, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(values)
   ```

Full code:
   ```go
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

   ```

## Contribute
If you want to contribute to this project, feel free to open a PR or an issue.
