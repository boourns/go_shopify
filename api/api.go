package shopify

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type API struct {
	URI    string
	Token  string
	Secret string
	client *http.Client
}

type errorResponse struct {
	Errors map[string]string `json:"errors"`
}

func (api *API) request(endpoint string, method string, params map[string]interface{}, body io.Reader) (result *bytes.Buffer, status int, err error) {
	if api.client == nil {
		api.client = &http.Client{}
	}

	uri := fmt.Sprintf("%s/%s", api.URI, endpoint)
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return
	}

	req.SetBasicAuth(api.Token, api.Secret)
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	fmt.Printf("resp %v err %v", resp, err)
	if err != nil {
		return
	}

	status = resp.StatusCode

	result = &bytes.Buffer{}
	if _, err = io.Copy(result, resp.Body); err != nil {
		return
	}

	return
}
