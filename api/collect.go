package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type Collect struct {
  
    CollectionId int64 `json:collection_id`
  
    CreatedAt string `json:created_at`
  
    Featured string `json:featured`
  
    Id int64 `json:id`
  
    Position string `json:position`
  
    ProductId int64 `json:product_id`
  
    SortValue string `json:sort_value`
  
    UpdatedAt string `json:updated_at`
  
}


// TODO implement Collect.create

// TODO implement Collect.destroy

func (api *API) Collect_index() (*[]Collect, error) {
  res, status, err := api.request("/admin/collects.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Collect{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Collect.count

// TODO implement Collect.show


