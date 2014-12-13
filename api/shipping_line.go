package shopify


import (
  
    "time"
  
)


type ShippingLine struct {
  
    Code string `json:"code"`
  
    Price time.Time `json:"price"`
  
    Source string `json:"source"`
  
    Title string `json:"title"`
  
    TaxLines []interface{} `json:"tax_lines"`
  

  
}



