package shopify


import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Page struct {
  
    Author string `json:"author"`
  
    BodyHtml string `json:"body_html"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Handle string `json:"handle"`
  
    Id int64 `json:"id"`
  
    PublishedAt time.Time `json:"published_at"`
  
    ShopId int64 `json:"shop_id"`
  
    TemplateSuffix string `json:"template_suffix"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  

  
    api *API
  
}


func (api *API) Pages() ([]Page, error) {
  res, status, err := api.request("/admin/pages.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Page{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["pages"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}




func (api *API) Page(id int64) (*Page, error) {
  endpoint := fmt.Sprintf("/admin/pages/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Page{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["page"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}


func (api *API) NewPage() *Page {
  return &Page{api: api}
}


func (obj *Page) Save() (error) {
  endpoint := fmt.Sprintf("/admin/pages/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/pages.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Page{}
  body["page"] = obj

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

  r := map[string]Page{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["page"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





