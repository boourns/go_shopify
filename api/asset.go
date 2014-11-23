package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Asset struct {
  
    Attachment string `json:attachment`
  
    ContentType string `json:content_type`
  
    CreatedAt time.Time `json:created_at`
  
    Key string `json:key`
  
    PublicUrl string `json:public_url`
  
    Size int64 `json:size`
  
    SourceKey string `json:source_key`
  
    Src string `json:src`
  
    ThemeId int64 `json:theme_id`
  
    UpdatedAt time.Time `json:updated_at`
  
    Value string `json:value`
  
}


func (api *API) Asset_index() (*[]Asset, error) {
  res, status, err := api.request("/admin/assets.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Asset{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Asset.show

// TODO implement Asset.update

// TODO implement Asset.destroy


