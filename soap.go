package soap

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type XMLResponse struct {
	HTTPResponse *http.Response
	Body         *XMLObject

	// Pretty print string of structure for help getting values
	Structure string
}

type XMLObject struct {
	Contents map[string]any
	Parent   *XMLObject
	Children []*XMLObject

	// Times nested
	Depth int
}

// Send a POST http request with SOAP payload
func SoapCall(url string, headers map[string]string, template string, values []any) (*http.Response, error) {
	// Verify parameter count matches template
	err := VerifyParameterList(template, values)
	if err != nil {
		return nil, err
	}

	// Start http client
	client := &http.Client{}

	// Create request body from template and values
	requestBody := fmt.Sprintf(template, values...)

	// Create request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return nil, err
	}

	// Add headers
	for label, value := range headers {
		request.Header.Add(label, value)
	}

	// Make request
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ParseXMLResponse(response *http.Response) (*XMLResponse, error) {
	// Initalize root object
	xmlResponse := XMLResponse{HTTPResponse: response}
	rootObject := XMLObject{Contents: make(map[string]any)}
	currentObject := &rootObject
	depth := 0

	// Get response body
	readResponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Add newlines after end carrot '>'
	newResponseBody := strings.ReplaceAll(string(readResponseBody), ">", ">\n")
	newResponseBody = strings.ReplaceAll(newResponseBody, "<", "\n<")

	// Scan each line
	reader := strings.NewReader(newResponseBody)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip XML instruction lines
		if strings.Contains(line, `<?`) {
			continue
			// Skip empty lines
		} else if line == "" {
			continue
			// Skip end tokens
		} else if strings.Contains(line, "</") {
			currentObject = currentObject.Parent
			depth--
		} else if strings.Contains(line, "<") {
			// Get indentation level
			depth++
			var tab string
			for range depth {
				tab = tab + "  "
			}
			// Get name of section
			reg := regexp.MustCompile(`<(\S+)[\s>]`)
			match := reg.FindAllStringSubmatch(line, 1)
			// Create XML Object record
			child := XMLObject{Contents: make(map[string]any), Parent: currentObject, Depth: depth}
			// Append values to parent object
			currentObject.Children = append(currentObject.Children, &child)
			// Verify key doesn't already exist, if it does increment number to avoid conflict
			_, doesExist := currentObject.Contents[match[0][1]]
			if !doesExist {
				currentObject.Contents[match[0][1]] = child
				// Print for format reference
				xmlResponse.Structure = xmlResponse.Structure + tab + match[0][1] + "\n"
			} else {
				currentObject.Contents[match[0][1]+"_"+fmt.Sprint(len(currentObject.Children))] = child
				// Print for format reference
				xmlResponse.Structure = xmlResponse.Structure + tab + match[0][1] + "_" + fmt.Sprint(len(currentObject.Children)) + "\n"
			}
			// Make this the current object
			currentObject = &child
		} else {
			// Assign value
			currentObject.Contents["value"] = line
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	xmlResponse.Body = &rootObject
	return &xmlResponse, nil
}
