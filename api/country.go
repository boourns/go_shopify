package shopify


import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
)


type Country struct {
  
    Code string `json:"code"`
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    Tax float64 `json:"tax"`
  
    TaxName string `json:"tax_name"`
  
    Provinces []interface{} `json:"provinces"`
  
    WeightBasedShippingRates []WeightBasedShippingRate `json:"weight_based_shipping_rates"`
  
    PriceBasedShippingRates []interface{} `json:"price_based_shipping_rates"`
  
    CarrierShippingRateProviders []interface{} `json:"carrier_shipping_rate_providers"`
  

  
    api *API
  
}


func (api *API) Countries() ([]Country, error) {
  res, status, err := api.request("/admin/countries.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Country{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["countries"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Country(id int64) (*Country, error) {
  endpoint := fmt.Sprintf("/admin/countries/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

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

  result.api = api

  return &result, nil
}


func (api *API) NewCountry() *Country {
  return &Country{api: api}
}


func (obj *Country) Save() (error) {
  endpoint := fmt.Sprintf("/admin/countries/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/countries.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Country{}
  body["country"] = obj

  buf := &bytes.Buffer{}
  err := json.NewEncoder(buf).Encode(body)

  if err != nil {
    return err
  }

  res, status, err := obj.api.request(endpoint, method, nil, buf)

  if err != nil {
    return err
  }

  if status != expectedStatus {
    r := errorResponse{}
    err = json.NewDecoder(res).Decode(&r)
    if err == nil {
      return fmt.Errorf("Status %d: %v", status, r.Errors)
    } else {
      return fmt.Errorf("Status %d, and error parsing body: %s", status, err)
    }
  }

  r := map[string]Country{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["country"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





