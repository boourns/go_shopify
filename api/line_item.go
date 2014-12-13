package shopify


import (
  
    "time"
  
)


type LineItem struct {
  
    AppliedDiscounts []interface{} `json:"applied_discounts"`
  
    CompareAtPrice string `json:"compare_at_price"`
  
    FulfillmentService string `json:"fulfillment_service"`
  
    GiftCard bool `json:"gift_card"`
  
    Grams int64 `json:"grams"`
  
    LinePrice time.Time `json:"line_price"`
  
    Price string `json:"price"`
  
    ProductId int64 `json:"product_id"`
  
    Properties string `json:"properties"`
  
    Quantity int64 `json:"quantity"`
  
    RequiresShipping bool `json:"requires_shipping"`
  
    Sku string `json:"sku"`
  
    TaxLines []interface{} `json:"tax_lines"`
  
    Taxable bool `json:"taxable"`
  
    Title string `json:"title"`
  
    VariantId int64 `json:"variant_id"`
  
    VariantTitle string `json:"variant_title"`
  
    Vendor string `json:"vendor"`
  

  
}



