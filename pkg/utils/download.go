package utils

import (
	"io/ioutil"
	"net/http"
)

func DownloadFile(url string, dst string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(dst, content, 0644)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
