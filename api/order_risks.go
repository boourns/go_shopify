package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type OrderRisks struct {
  
    CauseCancel string `json:cause_cancel`
  
    Display string `json:display`
  
    Id int64 `json:id`
  
    OrderId int64 `json:order_id`
  
    Message string `json:message`
  
    Recommendation string `json:recommendation`
  
    Score string `json:score`
  
    Source string `json:source`
  
}


// TODO implement OrderRisks.create

func (api *API) OrderRisks_index() (*[]OrderRisks, error) {
  res, status, err := api.request("/admin/order_risks.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]OrderRisks{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement OrderRisks.show

// TODO implement OrderRisks.update

// TODO implement OrderRisks.destroy


