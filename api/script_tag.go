package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type ScriptTag struct {
  
    CreatedAt time.Time `json:created_at`
  
    Event string `json:event`
  
    Id string `json:id`
  
    Src string `json:src`
  
    UpdatedAt time.Time `json:updated_at`
  
}


func (api *API) ScriptTag_index() (*[]ScriptTag, error) {
  res, status, err := api.request("/admin/script_tags.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]ScriptTag{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement ScriptTag.count

// TODO implement ScriptTag.show

// TODO implement ScriptTag.create

// TODO implement ScriptTag.update

// TODO implement ScriptTag.destroy


