package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Collect struct {
  
    CollectionId int64 `json:"collection_id"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Featured string `json:"featured"`
  
    Id int64 `json:"id"`
  
    Position int64 `json:"position"`
  
    ProductId int64 `json:"product_id"`
  
    SortValue string `json:"sort_value"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}






func (api *API) Collects() (*[]Collect, error) {
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

  result := (*r)["collect"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Collect(id int64) (*Collect, error) {
  endpoint := fmt.Sprintf("/admin/collects/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Collect{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["collect"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}



