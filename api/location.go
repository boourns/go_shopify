package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Location struct {
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    LocationType string `json:"location_type"`
  
    Address1 string `json:"address1"`
  
    Address2 string `json:"address2"`
  
    Zip string `json:"zip"`
  
    City string `json:"city"`
  
    Province string `json:"province"`
  
    Country string `json:"country"`
  
    Phone string `json:"phone"`
  
    CreatedAt time.Time `json:"created_at"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}


func (api *API) Locations() (*[]Location, error) {
  res, status, err := api.request("/admin/locations.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Location{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) Location(id int64) (*Location, error) {
  endpoint := fmt.Sprintf("/admin/locations/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

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

  return &result, nil
}



