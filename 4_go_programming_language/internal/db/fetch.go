package db

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Fetch comics from xkcd json api
func Fetch(num string) (string, error) {
	slice := [3]string{"https://xkcd.com/", num, "/info.0.json"}
	url := strings.Join(slice[0:], "")

	resp, reqError := http.Get(url)

	if reqError != nil {
		return "", reqError
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if readErr != nil {
		return "", readErr
	}

	return string(body), nil
}
