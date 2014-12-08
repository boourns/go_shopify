package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Transaction struct {
  
    Amount time.Time `json:"amount"`
  
    Authorization string `json:"authorization"`
  
    CreatedAt time.Time `json:"created_at"`
  
    DeviceId string `json:"device_id"`
  
    Gateway string `json:"gateway"`
  
    SourceName string `json:"source_name"`
  
    PaymentDetails string `json:"payment_details"`
  
    Id string `json:"id"`
  
    Kind string `json:"kind"`
  
    OrderId int64 `json:"order_id"`
  
    Receipt string `json:"receipt"`
  
    Status string `json:"status"`
  
    Test string `json:"test"`
  
    UserId string `json:"user_id"`
  
    Currency string `json:"currency"`
  
  api *API
}


func (api *API) Transactions() ([]Transaction, error) {
  res, status, err := api.request("/admin/transactions.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Transaction{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["transactions"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Transaction(id int64) (*Transaction, error) {
  endpoint := fmt.Sprintf("/admin/transactions/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Transaction{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["transaction"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewTransaction() *Transaction {
  return &Transaction{api: api}
}



