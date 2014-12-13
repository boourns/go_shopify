package shopify


import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Location struct {
  
    Address1 string `json:"address1"`
  
    Address2 string `json:"address2"`
  
    City string `json:"city"`
  
    Country string `json:"country"`
  
    CreatedAt time.Time `json:"created_at"`
  
    DeletedAt string `json:"deleted_at"`
  
    Id int64 `json:"id"`
  
    LocationType string `json:"location_type"`
  
    Name string `json:"name"`
  
    Phone string `json:"phone"`
  
    Province string `json:"province"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    Zip string `json:"zip"`
  
    CountryCode string `json:"country_code"`
  
    CountryName string `json:"country_name"`
  
    ProvinceCode string `json:"province_code"`
  

  
    api *API
  
}


func (api *API) Locations() ([]Location, error) {
  res, status, err := api.request("/admin/locations.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Location{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["locations"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}


func (api *API) Location(id int64) (*Location, error) {
  endpoint := fmt.Sprintf("/admin/locations/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Location{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["location"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}



