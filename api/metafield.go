package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Metafield struct {
  
    CreatedAt time.Time `json:created_at`
  
    Description string `json:description`
  
    Id int64 `json:id`
  
    Key string `json:key`
  
    Namespace string `json:namespace`
  
    OwnerId int64 `json:owner_id`
  
    OwnerResource string `json:owner_resource`
  
    Value time.Time `json:value`
  
    ValueType string `json:value_type`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) Metafield_index() (*[]Metafield, error) {
  res, status, err := api.request("/admin/metafields.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Metafield{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) Metafield_index() (*[]Metafield, error) {
  res, status, err := api.request("/admin/metafields.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Metafield{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) Metafield_index() (*[]Metafield, error) {
  res, status, err := api.request("/admin/metafields.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Metafield{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Metafield.count

// TODO implement Metafield.count

// TODO implement Metafield.show

// TODO implement Metafield.show

// TODO implement Metafield.create

// TODO implement Metafield.create

// TODO implement Metafield.update

// TODO implement Metafield.update

// TODO implement Metafield.destroy

// TODO implement Metafield.destroy


