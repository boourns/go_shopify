package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type FulfillmentService struct {
  
    CallbackUrl string `json:"callback_url"`
  
    Credential1 string `json:"credential1"`
  
    Credential2Exists string `json:"credential2_exists"`
  
    Format string `json:"format"`
  
    Handle string `json:"handle"`
  
    InventoryManagement string `json:"inventory_management"`
  
    Name string `json:"name"`
  
    ProviderId string `json:"provider_id"`
  
    RequiresShippingMethod string `json:"requires_shipping_method"`
  
    TrackingSupport string `json:"tracking_support"`
  
}


func (api *API) FulfillmentServices() (*[]FulfillmentService, error) {
  res, status, err := api.request("/admin/fulfillment_services.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]FulfillmentService{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["fulfillment_service"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) FulfillmentService(id int64) (*FulfillmentService, error) {
  endpoint := fmt.Sprintf("/admin/fulfillment_services/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]FulfillmentService{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["fulfillment_service"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}







