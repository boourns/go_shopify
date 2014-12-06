package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Product struct {
  
    BodyHtml string `json:"body_html"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Handle string `json:"handle"`
  
    Id int64 `json:"id"`
  
    Images []interface{} `json:"images"`
  
    Options []interface{} `json:"options"`
  
    ProductType string `json:"product_type"`
  
    PublishedAt time.Time `json:"published_at"`
  
    PublishedScope string `json:"published_scope"`
  
    Tags string `json:"tags"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    Variants []interface{} `json:"variants"`
  
    Vendor string `json:"vendor"`
  
}


func (api *API) Products() (*[]Product, error) {
  res, status, err := api.request("/admin/products.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Product{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["product"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Product(id int64) (*Product, error) {
  endpoint := fmt.Sprintf("/admin/products/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Product{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["product"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









