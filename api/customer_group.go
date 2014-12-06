package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type CustomerGroup struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Name time.Time `json:"name"`
  
    Query time.Time `json:"query"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}


func (api *API) CustomerGroups() (*[]CustomerGroup, error) {
  res, status, err := api.request("/admin/customer_groups.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]CustomerGroup{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) CustomerGroup(id int64) (*CustomerGroup, error) {
  endpoint := fmt.Sprintf("/admin/customer_groups/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]CustomerGroup{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["customer_group"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}











