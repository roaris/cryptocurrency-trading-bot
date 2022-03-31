package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DoHttpRequest(method, endpoint string, header, query map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
