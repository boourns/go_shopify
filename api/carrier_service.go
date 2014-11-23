package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type CarrierService struct {
  
    Active string `json:active`
  
    CallbackUrl string `json:callback_url`
  
    CarrierServiceType string `json:carrier_service_type`
  
    Name string `json:name`
  
    ServiceDiscovery string `json:service_discovery`
  
}


// TODO implement CarrierService.create

// TODO implement CarrierService.update

func (api *API) CarrierService_index() (*[]CarrierService, error) {
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


// TODO implement CarrierService.show

// TODO implement CarrierService.destroy


