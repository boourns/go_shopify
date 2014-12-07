package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Event struct {
  
    Arguments []string `json:"arguments"`
  
    Body string `json:"body"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Id int64 `json:"id"`
  
    Description string `json:"description"`
  
    Path string `json:"path"`
  
    Message string `json:"message"`
  
    SubjectId int64 `json:"subject_id"`
  
    SubjectType string `json:"subject_type"`
  
    Verb string `json:"verb"`
  
  api *API
}


func (api *API) Events() (*[]Event, error) {
  res, status, err := api.request("/admin/events.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Event{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["event"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}


func (api *API) Event(id int64) (*Event, error) {
  endpoint := fmt.Sprintf("/admin/events/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Event{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["event"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}





