package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Order struct {
  
    BillingAddress string `json:billing_address`
  
    BrowserIp string `json:browser_ip`
  
    BuyerAcceptsMarketing string `json:buyer_accepts_marketing`
  
    CancelReason string `json:cancel_reason`
  
    CancelledAt string `json:cancelled_at`
  
    CartToken string `json:cart_token`
  
    ClientDetails string `json:client_details`
  
    ClosedAt string `json:closed_at`
  
    CreatedAt time.Time `json:created_at`
  
    Currency string `json:currency`
  
    Customer string `json:customer`
  
    DiscountCodes string `json:discount_codes`
  
    Email string `json:email`
  
    FinancialStatus string `json:financial_status`
  
    Fulfillments time.Time `json:fulfillments`
  
    FulfillmentStatus string `json:fulfillment_status`
  
    Tags string `json:tags`
  
    Gateway string `json:gateway`
  
    Id string `json:id`
  
    LandingSite string `json:landing_site`
  
    LineItems int64 `json:line_items`
  
    Name time.Time `json:name`
  
    Note string `json:note`
  
    NoteAttributes string `json:note_attributes`
  
    Number string `json:number`
  
    OrderNumber int64 `json:order_number`
  
    PaymentDetails string `json:payment_details`
  
    ProcessedAt time.Time `json:processed_at`
  
    ProcessingMethod string `json:processing_method`
  
    ReferringSite string `json:referring_site`
  
    Refund string `json:refund`
  
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


func (api *API) Order_index() (*[]Order, error) {
  res, status, err := api.request("/admin/orders.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Order{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}


// TODO implement Order.show

// TODO implement Order.count

// TODO implement Order.close

// TODO implement Order.open

// TODO implement Order.cancel

// TODO implement Order.create

// TODO implement Order.update

// TODO implement Order.destroy


