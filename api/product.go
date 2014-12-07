package shopify

import (
  
    "bytes"
  
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
  
    Variants []ProductVariant `json:"variants"`
  
    Vendor string `json:"vendor"`
  
  api *API
}


func (api *API) Products() (*[]Product, error) {
  res, status, err := api.request("/admin/products.json", "GET", nil, nil)

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

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) Product(id int64) (*Product, error) {
  endpoint := fmt.Sprintf("/admin/products/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

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

  result.api = api

  return &result, nil
}


func (api *API) NewProduct() *Product {
  return &Product{api: api}
}


func (obj *Product) Save() (error) {
  endpoint := fmt.Sprintf("/admin/products/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/products.json")
    method = "POST"
    expectedStatus = 201
  }

  buf := &bytes.Buffer{}
  err := json.NewEncoder(buf).Encode(obj)

  if err != nil {
    return err
  }

  res, status, err := obj.api.request(endpoint, method, nil, buf)

  if err != nil {
    return err
  }

  if status != expectedStatus {
    return fmt.Errorf("Status returned: %d", status)
  }

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





