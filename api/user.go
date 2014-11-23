package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
)

type User struct {
  
    AccountOwner string `json:account_owner`
  
    Bio string `json:bio`
  
    Email string `json:email`
  
    FirstName string `json:first_name`
  
    Id int64 `json:id`
  
    Im string `json:im`
  
    LastName string `json:last_name`
  
    Permissions string `json:permissions`
  
    Phone string `json:phone`
  
    Pin string `json:pin`
  
    ReceiveAnnouncements int64 `json:receive_announcements`
  
    ScreenName string `json:screen_name`
  
    Url string `json:url`
  
    UserType string `json:user_type`
  
}


func (api *API) User_index() (*[]User, error) {
  res, status, err := api.request("/admin/users.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]User{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement User.show


