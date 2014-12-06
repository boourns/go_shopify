package shopify

import (
  
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
  
}


func (api *API) ProductImages() (*[]ProductImage, error) {
  res, status, err := api.request("/admin/product_images.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]ProductImage{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) ProductImage(id int64) (*ProductImage, error) {
  endpoint := fmt.Sprintf("/admin/product_images/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

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

  return &result, nil
}









