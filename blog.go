package shopify


import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Blog struct {
  
    Commentable string `json:"commentable"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Feedburner string `json:"feedburner"`
  
    FeedburnerLocation string `json:"feedburner_location"`
  
    Handle string `json:"handle"`
  
    Id int64 `json:"id"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    Tags string `json:"tags"`
  

  
    api *API
  
}


func (api *API) Blogs() ([]Blog, error) {
  res, status, err := api.request("/admin/blogs.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Blog{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["blogs"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Blog(id int64) (*Blog, error) {
  endpoint := fmt.Sprintf("/admin/blogs/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Blog{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["blog"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewBlog() *Blog {
  return &Blog{api: api}
}


func (obj *Blog) Save() (error) {
  endpoint := fmt.Sprintf("/admin/blogs/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/blogs.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Blog{}
  body["blog"] = obj

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

  r := map[string]Blog{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["blog"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





