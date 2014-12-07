package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Theme struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    Role string `json:"role"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    Previewable string `json:"previewable"`
  
    Processing string `json:"processing"`
  
  api *API
}


func (api *API) Themes() (*[]Theme, error) {
  res, status, err := api.request("/admin/themes.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Theme{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["theme"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}


func (api *API) Theme(id int64) (*Theme, error) {
  endpoint := fmt.Sprintf("/admin/themes/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Theme{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["theme"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewTheme() *Theme {
  return &Theme{api: api}
}


func (obj *Theme) Save() (error) {
  endpoint := fmt.Sprintf("/admin/themes/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/themes.json")
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





