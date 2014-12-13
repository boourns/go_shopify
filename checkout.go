package shopify


import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Checkout struct {
  
    BuyerAcceptsMarketing bool `json:"buyer_accepts_marketing"`
  
    CartToken string `json:"cart_token"`
  
    ClosedAt string `json:"closed_at"`
  
    CompletedAt string `json:"completed_at"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Currency string `json:"currency"`
  
    Email string `json:"email"`
  
    Gateway string `json:"gateway"`
  
    Id int64 `json:"id"`
  
    LandingSite string `json:"landing_site"`
  
    Note string `json:"note"`
  
    ReferringSite string `json:"referring_site"`
  
    ShippingLines []interface{} `json:"shipping_lines"`
  
    SourceIdentifier string `json:"source_identifier"`
  
    SourceName string `json:"source_name"`
  
    SourceUrl string `json:"source_url"`
  
    SubtotalPrice time.Time `json:"subtotal_price"`
  
    TaxesIncluded bool `json:"taxes_included"`
  
    Token time.Time `json:"token"`
  
    TotalDiscounts string `json:"total_discounts"`
  
    TotalLineItemsPrice time.Time `json:"total_line_items_price"`
  
    TotalPrice time.Time `json:"total_price"`
  
    TotalTax string `json:"total_tax"`
  
    TotalWeight int64 `json:"total_weight"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    LineItems []LineItem `json:"line_items"`
  
    Name string `json:"name"`
  
    NoteAttributes []NoteAttribute `json:"note_attributes"`
  
    Source string `json:"source"`
  
    DiscountCodes []interface{} `json:"discount_codes"`
  
    AbandonedCheckoutUrl string `json:"abandoned_checkout_url"`
  
    TaxLines []interface{} `json:"tax_lines"`
  
    BillingAddress BillingAddress `json:"billing_address"`
  
    ShippingAddress BillingAddress `json:"shipping_address"`
  
    Customer Customer `json:"customer"`
  

  
    api *API
  
}




func (api *API) Checkouts() ([]Checkout, error) {
  res, status, err := api.request("/admin/checkouts.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Checkout{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["checkouts"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}



