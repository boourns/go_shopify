package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type CustomerSavedSearch struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Name time.Time `json:"name"`
  
    Query time.Time `json:"query"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
  api *API
}


func (api *API) CustomerSavedSearches() (*[]CustomerSavedSearch, error) {
  res, status, err := api.request("/admin/customer_saved_searches.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]CustomerSavedSearch{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["customer_saved_search"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) CustomerSavedSearch(id int64) (*CustomerSavedSearch, error) {
  endpoint := fmt.Sprintf("/admin/customer_saved_searches/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]CustomerSavedSearch{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["customer_saved_search"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}




func (api *API) NewCustomerSavedSearch() *CustomerSavedSearch {
  return &CustomerSavedSearch{api: api}
}


func (obj *CustomerSavedSearch) Save() (error) {
  endpoint := fmt.Sprintf("/admin/customer_saved_searches/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/customer_saved_searches.json")
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





