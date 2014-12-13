package shopify


import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Customer struct {
  
    AcceptsMarketing bool `json:"accepts_marketing"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Email string `json:"email"`
  
    FirstName string `json:"first_name"`
  
    Id int64 `json:"id"`
  
    LastName string `json:"last_name"`
  
    LastOrderId string `json:"last_order_id"`
  
    MultipassIdentifier string `json:"multipass_identifier"`
  
    Note string `json:"note"`
  
    OrdersCount int64 `json:"orders_count"`
  
    State string `json:"state"`
  
    TotalSpent string `json:"total_spent"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    VerifiedEmail bool `json:"verified_email"`
  
    Tags string `json:"tags"`
  
    LastOrderName string `json:"last_order_name"`
  
    DefaultAddress DefaultAddress `json:"default_address"`
  
    Addresses []DefaultAddress `json:"addresses"`
  

  
    api *API
  
}


func (api *API) Customers() ([]Customer, error) {
  res, status, err := api.request("/admin/customers.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Customer{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["customers"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Customer(id int64) (*Customer, error) {
  endpoint := fmt.Sprintf("/admin/customers/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Customer{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["customer"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewCustomer() *Customer {
  return &Customer{api: api}
}


func (obj *Customer) Save() (error) {
  endpoint := fmt.Sprintf("/admin/customers/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/customers.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Customer{}
  body["customer"] = obj

  buf := &bytes.Buffer{}
  err := json.NewEncoder(buf).Encode(body)

  if err != nil {
    return err
  }

  res, status, err := obj.api.request(endpoint, method, nil, buf)

  if err != nil {
    return err
  }

  if status != expectedStatus {
    r := errorResponse{}
    err = json.NewDecoder(res).Decode(&r)
    if err == nil {
      return fmt.Errorf("Status %d: %v", status, r.Errors)
    } else {
      return fmt.Errorf("Status %d, and error parsing body: %s", status, err)
    }
  }

  r := map[string]Customer{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["customer"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}









