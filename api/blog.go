package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Blog struct {
  
    Commentable string `json:"commentable"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Feedburner string `json:"feedburner"`
  
    FeedburnerLocation string `json:"feedburner_location"`
  
    Handle string `json:"handle"`
  
    Id int64 `json:"id"`
  
    Metafield string `json:"metafield"`
  
    Tags string `json:"tags"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
}


func (api *API) Blogs() (*[]Blog, error) {
  res, status, err := api.request("/admin/blogs.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Blog{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Blog(id int64) (*Blog, error) {
  endpoint := fmt.Sprintf("/admin/blogs/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Blog{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["blog"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









