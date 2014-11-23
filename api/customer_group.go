package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type CustomerGroup struct {
  
    CreatedAt time.Time `json:created_at`
  
    Id string `json:id`
  
    Name time.Time `json:name`
  
    Query time.Time `json:query`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) CustomerGroup_index() (*[]CustomerGroup, error) {
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


// TODO implement CustomerGroup.count

// TODO implement CustomerGroup.show

// TODO implement CustomerGroup.other

// TODO implement CustomerGroup.create

// TODO implement CustomerGroup.update

// TODO implement CustomerGroup.destroy


