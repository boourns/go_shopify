package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type SmartCollection struct {
  
    BodyHtml string `json:body_html`
  
    Handle string `json:handle`
  
    Id string `json:id`
  
    Image string `json:image`
  
    PublishedAt time.Time `json:published_at`
  
    PublishedScope string `json:published_scope`
  
    Rules string `json:rules`
  
    Disjunctive string `json:disjunctive`
  
    SortOrder string `json:sort_order`
  
    TemplateSuffix string `json:template_suffix`
  
    Title string `json:title`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) SmartCollection_index() (*[]SmartCollection, error) {
  res, status, err := api.request("/admin/smart_collections.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]SmartCollection{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement SmartCollection.count

// TODO implement SmartCollection.show

// TODO implement SmartCollection.create

// TODO implement SmartCollection.update

// TODO implement SmartCollection.order

// TODO implement SmartCollection.destroy


