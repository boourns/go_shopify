package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Article struct {
  
    Author string `json:"author"`
  
    BlogId int64 `json:"blog_id"`
  
    BodyHtml string `json:"body_html"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Metafield string `json:"metafield"`
  
    Published string `json:"published"`
  
    PublishedAt time.Time `json:"published_at"`
  
    SummaryHtml string `json:"summary_html"`
  
    Tags string `json:"tags"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    UserId int64 `json:"user_id"`
  
}


func (api *API) Articles() (*[]Article, error) {
  res, status, err := api.request("/admin/articles.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Article{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["article"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Article(id int64) (*Article, error) {
  endpoint := fmt.Sprintf("/admin/articles/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Article{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["article"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}













