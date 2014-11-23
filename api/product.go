package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Product struct {
  
    BodyHtml string `json:body_html`
  
    CreatedAt time.Time `json:created_at`
  
    Handle string `json:handle`
  
    Id int64 `json:id`
  
    Images string `json:images`
  
    Options string `json:options`
  
    ProductType string `json:product_type`
  
    PublishedAt time.Time `json:published_at`
  
    PublishedScope string `json:published_scope`
  
    Tags string `json:tags`
  
    TemplateSuffix string `json:template_suffix`
  
    Title string `json:title`
  
    UpdatedAt time.Time `json:updated_at`
  
    Variants string `json:variants`
  
    Vendor string `json:vendor`
  
}


func (api *API) Product_index() (*[]Product, error) {
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

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Product.count

// TODO implement Product.show

// TODO implement Product.create

// TODO implement Product.update

// TODO implement Product.destroy


