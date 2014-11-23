package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Webhook struct {
  
    Address string `json:address`
  
    CreatedAt time.Time `json:created_at`
  
    Fields string `json:fields`
  
    Format string `json:format`
  
    Id string `json:id`
  
    MetafieldNamespaces string `json:metafield_namespaces`
  
    Topic string `json:topic`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) Webhook_index() (*[]Webhook, error) {
  res, status, err := api.request("/admin/webhooks.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Webhook{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Webhook.count

// TODO implement Webhook.show

// TODO implement Webhook.create

// TODO implement Webhook.update

// TODO implement Webhook.destroy


