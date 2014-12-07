package shopify

import (
	"bytes"

	"encoding/json"

	"fmt"

	"time"
)

type ScriptTag struct {
	CreatedAt time.Time `json:"created_at"`

	Event string `json:"event"`

	Id int64 `json:"id"`

	Src string `json:"src"`

	UpdatedAt time.Time `json:"updated_at"`

	api *API
}

func (api *API) ScriptTags() (*[]ScriptTag, error) {
	res, status, err := api.request("/admin/script_tags.json", "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := &map[string][]ScriptTag{}
	err = json.NewDecoder(res).Decode(r)

	fmt.Printf("things are: %v\n\n", *r)

	result := (*r)["script_tag"]

	if err != nil {
		return nil, err
	}

	for _, v := range result {
		v.api = api
	}

	return &result, nil
}

func (api *API) ScriptTag(id int64) (*ScriptTag, error) {
	endpoint := fmt.Sprintf("/admin/script_tags/%d.json", id)

	res, status, err := api.request(endpoint, "GET", nil, nil)

	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("Status returned: %d", status)
	}

	r := map[string]ScriptTag{}
	err = json.NewDecoder(res).Decode(&r)

	fmt.Printf("things are: %v\n\n", r)

	result := r["script_tag"]

	if err != nil {
		return nil, err
	}

	result.api = api

	return &result, nil
}

func (api *API) NewScriptTag() *ScriptTag {
	return &ScriptTag{api: api}
}

func (obj *ScriptTag) Save() error {
	endpoint := fmt.Sprintf("/admin/script_tags/%d.json", obj.Id)
	method := "PUT"
	expectedStatus := 201

	if obj.Id == 0 {
		endpoint = fmt.Sprintf("/admin/script_tags.json")
		method = "POST"
		expectedStatus = 201
	}

	body := map[string]*ScriptTag{}
	body["script_tag"] = obj

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(body)

	if err != nil {
		return err
	}

	res, status, err := obj.api.request(endpoint, method, nil, buf)

	if err != nil {
		return err
	}

	if status != expectedStatus {
		r := errorResponse{}
		err = json.NewDecoder(res).Decode(&r)
		if err == nil {
			return fmt.Errorf("Status %d: %v", status, r.Errors)
		} else {
			return fmt.Errorf("Status %d, and error parsing body: %s", status, err)
		}
	}

	fmt.Printf("things are: %v\n\n", res)

	return nil
}
