package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type RecurringApplicationCharge struct {
  
    ActivatedOn string `json:activated_on`
  
    BillingOn string `json:billing_on`
  
    CancelledOn string `json:cancelled_on`
  
    ConfirmationUrl string `json:confirmation_url`
  
    CreatedAt time.Time `json:created_at`
  
    Id int64 `json:id`
  
    Name string `json:name`
  
    Price string `json:price`
  
    ReturnUrl string `json:return_url`
  
    Test string `json:test`
  
    TrialDays string `json:trial_days`
  
    TrialEndsOn string `json:trial_ends_on`
  
    UpdatedAt time.Time `json:updated_at`
  
}


// TODO implement RecurringApplicationCharge.create

// TODO implement RecurringApplicationCharge.show

func (api *API) RecurringApplicationCharge_index() (*[]RecurringApplicationCharge, error) {
  res, status, err := api.request("/admin/recurring_application_charges.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]RecurringApplicationCharge{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement RecurringApplicationCharge.activate

// TODO implement RecurringApplicationCharge.destroy


