# Go-Soap

Go-Soap is an open-source library for handling SOAP requests and responses in Go. 

The parsing of responses into generic objects is useful for handling dynamic XML responses, or for testing phases where creating structs may not be worth the time. It may not be ideal for production use cases.

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

2. Create a new XML file template using `%v` as variable placeholders
    ```xml
        <?xml version="1.0" encoding="utf-8"?>
          <soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
            <soap12:Body>
              <FahrenheitToCelsius xmlns="https://www.w3schools.com/xml/">
                <Fahrenheit>%v</Fahrenheit>
                </FahrenheitToCelsius>
            </soap12:Body>
          </soap12:Envelope>
    ```

3. Use the `GetTemplate` function to load your template into Go
   ```go
    template, err := soap.GetTemplate("fahrenheit_to_celsius.xml")
    if err != nil {
		fmt.Println(err)
	}
   ```
   This will return an error if unable to fetch the template provided

4. Set parameters as a `[]any` array. Ensure parameters are in the order they appear in your template
   ```go
    values := []any{75}
   ```

5. Create headers in a type `map[string]string` 
   ```go
    headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"
   ```

6. Use the `SoapCall` function to make the request
   ```go
   	res, err := soap.SoapCall(url, headers, template, values)
	if err != nil {
		// Handle error
		panic(err)
	}
   ```
   This will look at the number of paramters provided, and compare it to the number of parameters detected in your template designated by the `%v` placeholder. If they do not match, or there is an error making the request, the function will return an error.

7. Use the `ParseXMLResponse` function to create a readable object from the XML response
   ```go
    xmlRes, err := soap.ParseXMLResponse(res)
	if err != nil {
		fmt.Println(err)
	}
   ```
   This will parse the data into a series of `XMLObject` that contain a `Contents` field of type `map[string]any` . This may contain another `XMLObject` if nesting has occured, or will have a value designated by the key `value`

8. To view the structure of the response, print the `Structure` field of your `XMLResponse` object
   ```go
    fmt.Println(xmlRes.Structure)
   ```

9. To access a value from the object, use the syntax below to drill down to the location, and access it using the `value` key
   ```go
	fmt.Println(xmlRes.Body.Contents["soap:Envelope"].(soap.XMLObject).Contents["soap:Body"].(soap.XMLObject).Contents["FahrenheitToCelsiusResponse"].(soap.XMLObject).Contents["FahrenheitToCelsiusResult"].(soap.XMLObject).Contents["value"])
   ```
   This is lengthy syntax, but allows for dynamic structures to be generated

Full code:
   ```go
	const url string = "https://www.w3schools.com/xml/tempconvert.asmx"

	template := soap.GetTemplate("fahrenheit_to_celsius.xml")

	values := []any{75}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	res, err := soap.SoapCall(url, headers, template, values)
	if err != nil {
		// Handle error
		panic(err)
	}

	xmlRes, err := soap.ParseXMLResponse(res)
	if err != nil {
        // Handle error
		fmt.Println(err)
	}

	fmt.Println(xmlRes.Structure)

	fmt.Println(xmlRes.Body.Contents["soap:Envelope"].(soap.XMLObject).Contents["soap:Body"].(soap.XMLObject).Contents["FahrenheitToCelsiusResponse"].(soap.XMLObject).Contents["FahrenheitToCelsiusResult"].(soap.XMLObject).Contents["value"])
   ```

## Contribute
If you want to contribute to this project, feel free to open a PR or an issue.
