package shopify


import (
  
    "bytes"
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)


type Order struct {
  
    BuyerAcceptsMarketing bool `json:"buyer_accepts_marketing"`
  
    CancelReason string `json:"cancel_reason"`
  
    CancelledAt string `json:"cancelled_at"`
  
    CartToken string `json:"cart_token"`
  
    CheckoutToken string `json:"checkout_token"`
  
    ClosedAt string `json:"closed_at"`
  
    Confirmed bool `json:"confirmed"`
  
    CreatedAt time.Time `json:"created_at"`
  
    Currency string `json:"currency"`
  
    Email string `json:"email"`
  
    FinancialStatus string `json:"financial_status"`
  
    FulfillmentStatus string `json:"fulfillment_status"`
  
    Gateway string `json:"gateway"`
  
    Id int64 `json:"id"`
  
    LandingSite string `json:"landing_site"`
  
    LocationId string `json:"location_id"`
  
    Name string `json:"name"`
  
    Note string `json:"note"`
  
    Number int64 `json:"number"`
  
    ProcessedAt time.Time `json:"processed_at"`
  
    Reference string `json:"reference"`
  
    ReferringSite string `json:"referring_site"`
  
    SourceIdentifier string `json:"source_identifier"`
  
    SourceName string `json:"source_name"`
  
    SourceUrl string `json:"source_url"`
  
    SubtotalPrice time.Time `json:"subtotal_price"`
  
    TaxesIncluded bool `json:"taxes_included"`
  
    Test bool `json:"test"`
  
    Token string `json:"token"`
  
    TotalDiscounts string `json:"total_discounts"`
  
    TotalLineItemsPrice time.Time `json:"total_line_items_price"`
  
    TotalPrice time.Time `json:"total_price"`
  
    TotalPriceUsd string `json:"total_price_usd"`
  
    TotalTax string `json:"total_tax"`
  
    TotalWeight int64 `json:"total_weight"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    UserId string `json:"user_id"`
  
    BrowserIp string `json:"browser_ip"`
  
    LandingSiteRef string `json:"landing_site_ref"`
  
    OrderNumber int64 `json:"order_number"`
  
    DiscountCodes []interface{} `json:"discount_codes"`
  
    NoteAttributes []interface{} `json:"note_attributes"`
  
    ProcessingMethod string `json:"processing_method"`
  
    Source string `json:"source"`
  
    CheckoutId int64 `json:"checkout_id"`
  
    TaxLines []interface{} `json:"tax_lines"`
  
    Tags string `json:"tags"`
  
    LineItems []LineItem `json:"line_items"`
  
    ShippingLines []ShippingLine `json:"shipping_lines"`
  
    BillingAddress BillingAddress `json:"billing_address"`
  
    ShippingAddress BillingAddress `json:"shipping_address"`
  
    Fulfillments []interface{} `json:"fulfillments"`
  
    ClientDetails ClientDetail `json:"client_details"`
  
    Refunds []interface{} `json:"refunds"`
  
    Customer Customer `json:"customer"`
  

  
    api *API
  
}


func (api *API) Orders() ([]Order, error) {
  res, status, err := api.request("/admin/orders.json", "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]Order{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["orders"]

	if err != nil {
		return nil, err
  }

  for _, v := range result {
    v.api = api
  }

  return result, nil
}


func (api *API) Order(id int64) (*Order, error) {
  endpoint := fmt.Sprintf("/admin/orders/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Order{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["order"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}










func (api *API) NewOrder() *Order {
  return &Order{api: api}
}


func (obj *Order) Save() (error) {
  endpoint := fmt.Sprintf("/admin/orders/%d.json", obj.Id)
  method := "PUT"
  expectedStatus := 201

  if obj.Id == 0 {
    endpoint = fmt.Sprintf("/admin/orders.json")
    method = "POST"
    expectedStatus = 201
  }

  body := map[string]*Order{}
  body["order"] = obj

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

  r := map[string]Order{}
  err = json.NewDecoder(res).Decode(&r)

	if err != nil {
		return err
  }

  fmt.Printf("things are: %v\n\n", r)

  *obj = r["order"]

  fmt.Printf("things are: %v\n\n", res)

  return nil
}





