package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type Country struct {
  
    CarrierShippingRateProviders []interface{} `json:"carrier_shipping_rate_providers"`
  
    Code string `json:"code"`
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    PriceBasedShippingRates []interface{} `json:"price_based_shipping_rates"`
  
    Provinces []interface{} `json:"provinces"`
  
    Tax float64 `json:"tax"`
  
    WeightBasedShippingRates []interface{} `json:"weight_based_shipping_rates"`
  
}


func (api *API) Countries() (*[]Country, error) {
  res, status, err := api.request("/admin/countries.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Country{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["country"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Country(id int64) (*Country, error) {
  endpoint := fmt.Sprintf("/admin/countries/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Country{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["country"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









