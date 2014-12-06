package shopify

import (
  
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
  
}


func (api *API) Fulfillments() (*[]Fulfillment, error) {
  res, status, err := api.request("/admin/fulfillments.json", "GET", nil)

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

  return &result, nil
}




func (api *API) Fulfillment(id int64) (*Fulfillment, error) {
  endpoint := fmt.Sprintf("/admin/fulfillments/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

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

  return &result, nil
}











