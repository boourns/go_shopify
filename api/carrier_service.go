package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type CarrierService struct {
  
    Active string `json:"active"`
  
    CallbackUrl string `json:"callback_url"`
  
    CarrierServiceType string `json:"carrier_service_type"`
  
    Name string `json:"name"`
  
    ServiceDiscovery string `json:"service_discovery"`
  
}






func (api *API) CarrierServices() (*[]CarrierService, error) {
  res, status, err := api.request("/admin/carrier_services.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]CarrierService{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) CarrierService(id int64) (*CarrierService, error) {
  endpoint := fmt.Sprintf("/admin/carrier_services/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]CarrierService{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["carrier_service"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}





