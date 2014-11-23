package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Checkout struct {
  
    AbandonedCheckoutUrl string `json:abandoned_checkout_url`
  
    BillingAddress string `json:billing_address`
  
    BuyerAcceptsMarketing string `json:buyer_accepts_marketing`
  
    CancelReason string `json:cancel_reason`
  
    CartToken string `json:cart_token`
  
    ClosedAt string `json:closed_at`
  
    CompletedAt string `json:completed_at`
  
    CreatedAt time.Time `json:created_at`
  
    Currency string `json:currency`
  
    Customer string `json:customer`
  
    DiscountCodes string `json:discount_codes`
  
    Email string `json:email`
  
    Gateway string `json:gateway`
  
    Id string `json:id`
  
    LandingSite string `json:landing_site`
  
    LineItems string `json:line_items`
  
    Note string `json:note`
  
    ReferringSite string `json:referring_site`
  
    ShippingAddress string `json:shipping_address`
  
    ShippingLines string `json:shipping_lines`
  
    SourceName string `json:source_name`
  
    SubtotalPrice float64 `json:subtotal_price`
  
    TaxLines float64 `json:tax_lines`
  
    TaxesIncluded string `json:taxes_included`
  
    Token string `json:token`
  
    TotalDiscounts string `json:total_discounts`
  
    TotalLineItemsPrice string `json:total_line_items_price`
  
    TotalPrice time.Time `json:total_price`
  
    TotalTax time.Time `json:total_tax`
  
    TotalWeight string `json:total_weight`
  
    UpdatedAt time.Time `json:updated_at`
  
}


// TODO implement Checkout.count

func (api *API) Checkout_index() (*[]Checkout, error) {
  res, status, err := api.request("/admin/checkouts.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Checkout{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}



