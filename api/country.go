package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Country struct {
  
    CarrierShippingRateProviders string `json:carrier_shipping_rate_providers`
  
    Code string `json:code`
  
    Id int64 `json:id`
  
    Name string `json:name`
  
    PriceBasedShippingRates string `json:price_based_shipping_rates`
  
    Provinces string `json:provinces`
  
    Tax time.Time `json:tax`
  
    WeightBasedShippingRates string `json:weight_based_shipping_rates`
  
}


func (api *API) Country_index() (*[]Country, error) {
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

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Country.count

// TODO implement Country.show

// TODO implement Country.create

// TODO implement Country.update

// TODO implement Country.destroy


