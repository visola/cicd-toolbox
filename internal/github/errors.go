package github

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func generateErrorFrom(url string, response *http.Response) error {
	errorMessage := fmt.Sprintf("Got unexpected response code while fetching URL: %s", url)
	errorMessage = fmt.Sprintf("%s\nStatus: %s", errorMessage, response.Status)

	bodyData, _ := ioutil.ReadAll(response.Body)
	if len(bodyData) >= 0 {
		errorMessage = fmt.Sprintf("%s\nBody:\n%s\n", errorMessage, string(bodyData))
	}

	return errors.New(errorMessage)
}
