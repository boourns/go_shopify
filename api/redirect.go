package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
)

type Redirect struct {
  
    Id int64 `json:"id"`
  
    Path string `json:"path"`
  
    Target string `json:"target"`
  
  api *API
}


func (api *API) Redirects() ([]Redirect, error) {
  res, status, err := api.request("/admin/redirects.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Redirect{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["redirects"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Redirect(id int64) (*Redirect, error) {
  endpoint := fmt.Sprintf("/admin/redirects/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Redirect{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["redirect"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewRedirect() *Redirect {
  return &Redirect{api: api}
}


func (obj *Redirect) Save() (error) {
  endpoint := fmt.Sprintf("/admin/redirects/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/redirects.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Redirect{}
  body["redirect"] = obj

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

  r := map[string]Redirect{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["redirect"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





