package shopify

import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type ProductImage struct {
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Position string `json:"position"`
  
    ProductId string `json:"product_id"`
  
    VariantIds []string `json:"variant_ids"`
  
    Src string `json:"src"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
  api *API
}


func (api *API) ProductImages() (*[]ProductImage, error) {
  res, status, err := api.request("/admin/product_images.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]ProductImage{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["product_image"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) ProductImage(id int64) (*ProductImage, error) {
  endpoint := fmt.Sprintf("/admin/product_images/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]ProductImage{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["product_image"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewProductImage() *ProductImage {
  return &ProductImage{api: api}
}


func (obj *ProductImage) Save() (error) {
  endpoint := fmt.Sprintf("/admin/product_images/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/product_images.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*ProductImage{}
  body["product_image"] = obj

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

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





