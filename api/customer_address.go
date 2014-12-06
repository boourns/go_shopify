package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type CustomerAddress struct {
  
    Address1 string `json:"address1"`
  
    Address2 string `json:"address2"`
  
    City string `json:"city"`
  
    Company string `json:"company"`
  
    FirstName string `json:"first_name"`
  
    LastName string `json:"last_name"`
  
    Phone string `json:"phone"`
  
    Province string `json:"province"`
  
    Country string `json:"country"`
  
    Zip string `json:"zip"`
  
    Name string `json:"name"`
  
    ProvinceCode string `json:"province_code"`
  
    CountryCode string `json:"country_code"`
  
    CountryName string `json:"country_name"`
  
}


func (api *API) CustomerAddresses() (*[]CustomerAddress, error) {
  res, status, err := api.request("/admin/customer_addresses.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]CustomerAddress{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) CustomerAddress(id int64) (*CustomerAddress, error) {
  endpoint := fmt.Sprintf("/admin/customer_addresses/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]CustomerAddress{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["customer_address"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}













