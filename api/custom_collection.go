package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type CustomCollection struct {
  
    BodyHtml string `json:body_html`
  
    Handle string `json:handle`
  
    Image string `json:image`
  
    Id string `json:id`
  
    Metafield string `json:metafield`
  
    Published string `json:published`
  
    PublishedAt time.Time `json:published_at`
  
    PublishedScope string `json:published_scope`
  
    SortOrder string `json:sort_order`
  
    TemplateSuffix string `json:template_suffix`
  
    Title string `json:title`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) CustomCollection_index() (*[]CustomCollection, error) {
  res, status, err := api.request("/admin/custom_collections.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]CustomCollection{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement CustomCollection.count

// TODO implement CustomCollection.show

// TODO implement CustomCollection.create

// TODO implement CustomCollection.update

// TODO implement CustomCollection.destroy


