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
  
    Fields []interface{} `json:"fields"`
  
    Format string `json:"format"`
  
    Id int64 `json:"id"`
  
    MetafieldNamespaces []interface{} `json:"metafield_namespaces"`
  
    Topic string `json:"topic"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
  api *API
}


func (api *API) Webhooks() ([]Webhook, error) {
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

  result := (*r)["webhooks"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
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


func (obj *Webhook) Save() (error) {
  endpoint := fmt.Sprintf("/admin/webhooks/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/webhooks.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Webhook{}
  body["webhook"] = obj

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

  r := map[string]Webhook{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["webhook"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





