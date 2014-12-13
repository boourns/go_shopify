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
  
    ProductType string `json:"product_type"`
  
    PublishedAt time.Time `json:"published_at"`
  
    PublishedScope string `json:"published_scope"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    Vendor string `json:"vendor"`
  
    Tags string `json:"tags"`
  
    Variants []Variant `json:"variants"`
  
    Options []Option `json:"options"`
  
    Images []interface{} `json:"images"`
  

  
    api *API
  
}


func (api *API) Products() ([]Product, error) {
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

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
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

  body := map[string]*Product{}
  body["product"] = obj

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

  r := map[string]Product{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["product"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





