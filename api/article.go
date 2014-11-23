package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Article struct {
  
    Author string `json:author`
  
    BlogId int64 `json:blog_id`
  
    BodyHtml string `json:body_html`
  
    CreatedAt time.Time `json:created_at`
  
    Id int64 `json:id`
  
    Metafield string `json:metafield`
  
    Published string `json:published`
  
    PublishedAt time.Time `json:published_at`
  
    SummaryHtml string `json:summary_html`
  
    Tags string `json:tags`
  
    TemplateSuffix string `json:template_suffix`
  
    Title string `json:title`
  
    UpdatedAt time.Time `json:updated_at`
  
    UserId string `json:user_id`
  
}


func (api *API) Article_index() (*[]Article, error) {
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

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Article.count

// TODO implement Article.show

// TODO implement Article.create

// TODO implement Article.update

// TODO implement Article.authors

// TODO implement Article.tags

// TODO implement Article.destroy


