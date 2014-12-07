package shopify

import (
	"bytes"

	"encoding/json"

	"fmt"

	"time"
)

type Webhook struct {
	Address string `json:"address"`

	CreatedAt time.Time `json:"created_at"`

	Fields []string `json:"fields"`

	Format string `json:"format"`

	Id int64 `json:"id"`

	MetafieldNamespaces []string `json:"metafield_namespaces"`

	Topic string `json:"topic"`

	UpdatedAt time.Time `json:"updated_at"`

	api *API
}

func (api *API) Webhooks() (*[]Webhook, error) {
	res, status, err := api.request("/admin/webhooks.json", "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := &map[string][]Webhook{}
	err = json.NewDecoder(res).Decode(r)

	fmt.Printf("things are: %v\n\n", *r)

	result := (*r)["webhook"]

	if err != nil {
		return nil, err
	}

	for _, v := range result {
		v.api = api
	}

	return &result, nil
}

func (api *API) Webhook(id int64) (*Webhook, error) {
	endpoint := fmt.Sprintf("/admin/webhooks/%d.json", id)

	res, status, err := api.request(endpoint, "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := map[string]Webhook{}
	err = json.NewDecoder(res).Decode(&r)

	fmt.Printf("things are: %v\n\n", r)

	result := r["webhook"]

	if err != nil {
		return nil, err
	}

	result.api = api

	return &result, nil
}

func (api *API) NewWebhook() *Webhook {
	return &Webhook{api: api}
}

func (obj *Webhook) Save() error {
	endpoint := fmt.Sprintf("/admin/webhooks/%d.json", obj.Id)
	method := "PUT"
	expectedStatus := 201

	if obj.Id == 0 {
		endpoint = fmt.Sprintf("/admin/webhooks.json")
		method = "POST"
		expectedStatus = 201
	}

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(obj)

	if err != nil {
		return err
	}

	res, status, err := obj.api.request(endpoint, method, nil, buf)

	if err != nil {
		return err
	}

	if status != expectedStatus {
		return fmt.Errorf("Status returned: %d", status)
	}

	fmt.Printf("things are: %v\n\n", res)

	return nil
}
