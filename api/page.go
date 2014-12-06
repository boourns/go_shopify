package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Page struct {
  
    Author string `json:"author"`
  
    BodyHtml string `json:"body_html"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Handle string `json:"handle"`
  
    Id int64 `json:"id"`
  
    Metafield string `json:"metafield"`
  
    PublishedAt time.Time `json:"published_at"`
  
    ShopId int64 `json:"shop_id"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}


func (api *API) Pages() (*[]Page, error) {
  res, status, err := api.request("/admin/pages.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Page{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["page"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Page(id int64) (*Page, error) {
  endpoint := fmt.Sprintf("/admin/pages/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Page{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["page"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









