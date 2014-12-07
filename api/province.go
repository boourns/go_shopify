package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
)

type Province struct {
  
    Code string `json:"code"`
  
    CountryId int64 `json:"country_id"`
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    Tax float64 `json:"tax"`
  
    TaxName string `json:"tax_name"`
  
    TaxType string `json:"tax_type"`
  
    TaxPercentage float64 `json:"tax_percentage"`
  
  api *API
}


func (api *API) Provinces() (*[]Province, error) {
  res, status, err := api.request("/admin/provinces.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Province{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["province"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) Province(id int64) (*Province, error) {
  endpoint := fmt.Sprintf("/admin/provinces/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Province{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["province"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (obj *Province) Save() (error) {
  endpoint := fmt.Sprintf("/admin/provinces/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/provinces.json")
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



