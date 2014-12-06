package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Comment struct {
  
    ArticleId int64 `json:"article_id"`
  
    Author string `json:"author"`
  
    BlogId int64 `json:"blog_id"`
  
    Body string `json:"body"`
  
    BodyHtml string `json:"body_html"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Email string `json:"email"`
  
    Id int64 `json:"id"`
  
    Ip string `json:"ip"`
  
    PublishedAt time.Time `json:"published_at"`
  
    Status string `json:"status"`
  
    UpdatedAt string `json:"updated_at"`
  
    UserAgent string `json:"user_agent"`
  
}


func (api *API) Comments() (*[]Comment, error) {
  res, status, err := api.request("/admin/comments.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Comment{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["comment"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Comment(id int64) (*Comment, error) {
  endpoint := fmt.Sprintf("/admin/comments/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Comment{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["comment"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}

















