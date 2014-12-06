package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type Redirect struct {
  
    Id int64 `json:"id"`
  
    Path string `json:"path"`
  
    Target string `json:"target"`
  
}


func (api *API) Redirects() (*[]Redirect, error) {
  res, status, err := api.request("/admin/redirects.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Redirect{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["redirect"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) Redirect(id int64) (*Redirect, error) {
  endpoint := fmt.Sprintf("/admin/redirects/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Redirect{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["redirect"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









