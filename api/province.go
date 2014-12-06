package shopify

import (
  
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
  
}


func (api *API) Provinces() (*[]Province, error) {
  res, status, err := api.request("/admin/provinces.json", "GET", nil)

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

  return &result, nil
}




func (api *API) Province(id int64) (*Province, error) {
  endpoint := fmt.Sprintf("/admin/provinces/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

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

  return &result, nil
}





