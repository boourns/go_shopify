package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Customer struct {
  
    AcceptsMarketing string `json:accepts_marketing`
  
    Addresses string `json:addresses`
  
    CreatedAt time.Time `json:created_at`
  
    DefaultAddress string `json:default_address`
  
    Email string `json:email`
  
    FirstName string `json:first_name`
  
    Id int64 `json:id`
  
    Metafield string `json:metafield`
  
    MultipassIdentifier string `json:multipass_identifier`
  
    LastName string `json:last_name`
  
    LastOrderId string `json:last_order_id`
  
    LastOrderName string `json:last_order_name`
  
    Note string `json:note`
  
    OrdersCount int64 `json:orders_count`
  
    State string `json:state`
  
    Tags string `json:tags`
  
    TotalSpent float64 `json:total_spent`
  
    UpdatedAt time.Time `json:updated_at`
  
    VerifiedEmail string `json:verified_email`
  
}


func (api *API) Customer_index() (*[]Customer, error) {
  res, status, err := api.request("/admin/customers.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Customer{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Customer.search

// TODO implement Customer.show

// TODO implement Customer.create

// TODO implement Customer.update

// TODO implement Customer.destroy

// TODO implement Customer.count

// TODO implement Customer.orders


