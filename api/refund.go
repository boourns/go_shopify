package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Refund struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Note string `json:"note"`
  
    RefundLineItems int64 `json:"refund_line_items"`
  
    Restock string `json:"restock"`
  
    Transactions string `json:"transactions"`
  
    UserId int64 `json:"user_id"`
  
  api *API
}


func (api *API) Refund(id int64) (*Refund, error) {
  endpoint := fmt.Sprintf("/admin/refunds/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Refund{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["refund"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}



