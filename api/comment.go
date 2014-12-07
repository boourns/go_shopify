package shopify

import (
  
    "bytes"
  
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
  
  api *API
}


func (api *API) Comments() (*[]Comment, error) {
  res, status, err := api.request("/admin/comments.json", "GET", nil, nil)

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

  for _, v := range result {
    v.api = api
  }

  return &result, nil
}




func (api *API) Comment(id int64) (*Comment, error) {
  endpoint := fmt.Sprintf("/admin/comments/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

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

  result.api = api

  return &result, nil
}


func (api *API) NewComment() *Comment {
  return &Comment{api: api}
}


func (obj *Comment) Save() (error) {
  endpoint := fmt.Sprintf("/admin/comments/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/comments.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Comment{}
  body["comment"] = obj

  buf := &bytes.Buffer{}
  err := json.NewEncoder(buf).Encode(body)

  if err != nil {
    return err
  }

  res, status, err := obj.api.request(endpoint, method, nil, buf)

  if err != nil {
    return err
  }

  if status != expectedStatus {
    r := errorResponse{}
    err = json.NewDecoder(res).Decode(&r)
    if err == nil {
      return fmt.Errorf("Status %d: %v", status, r.Errors)
    } else {
      return fmt.Errorf("Status %d, and error parsing body: %s", status, err)
    }
  }

  fmt.Printf("things are: %v\n\n", res)

  return nil
}













