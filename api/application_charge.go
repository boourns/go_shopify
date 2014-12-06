package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type ApplicationCharge struct {
  
    ConfirmationUrl string `json:"confirmation_url"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    Price string `json:"price"`
  
    ReturnUrl string `json:"return_url"`
  
    Status string `json:"status"`
  
    Test string `json:"test"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}




func (api *API) ApplicationCharge(id int64) (*ApplicationCharge, error) {
  endpoint := fmt.Sprintf("/admin/application_charges/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]ApplicationCharge{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["application_charge"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


func (api *API) ApplicationCharges() (*[]ApplicationCharge, error) {
  res, status, err := api.request("/admin/application_charges.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]ApplicationCharge{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["application_charge"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}





