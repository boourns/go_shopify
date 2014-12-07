package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Fulfillment struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    LineItems int64 `json:"line_items"`
  
    OrderId int64 `json:"order_id"`
  
    Receipt string `json:"receipt"`
  
    Status string `json:"status"`
  
    TrackingCompany string `json:"tracking_company"`
  
    TrackingNumbers []string `json:"tracking_numbers"`
  
    TrackingUrls []string `json:"tracking_urls"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    VariantInventoryManagement string `json:"variant_inventory_management"`
  
  api *API
}


func (api *API) Fulfillments() (*[]Fulfillment, error) {
  res, status, err := api.request("/admin/fulfillments.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Fulfillment{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["fulfillment"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) Fulfillment(id int64) (*Fulfillment, error) {
  endpoint := fmt.Sprintf("/admin/fulfillments/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Fulfillment{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["fulfillment"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewFulfillment() *Fulfillment {
  return &Fulfillment{api: api}
}


func (obj *Fulfillment) Save() (error) {
  endpoint := fmt.Sprintf("/admin/fulfillments/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/fulfillments.json")
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







