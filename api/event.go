package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Event struct {
  
    Arguments string `json:arguments`
  
    Body string `json:body`
  
    CreatedAt time.Time `json:created_at`
  
    Id string `json:id`
  
    Description string `json:description`
  
    Path string `json:path`
  
    Message string `json:message`
  
    SubjectId string `json:subject_id`
  
    SubjectType string `json:subject_type`
  
    Verb string `json:verb`
  
}


func (api *API) Event_index() (*[]Event, error) {
  res, status, err := api.request("/admin/events.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Event{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Event.show

// TODO implement Event.count


