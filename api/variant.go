package shopify


import (
  
    "time"
  
)


type Variant struct {
  
    Barcode string `json:"barcode"`
  
    CompareAtPrice string `json:"compare_at_price"`
  
    CreatedAt time.Time `json:"created_at"`
  
    FulfillmentService string `json:"fulfillment_service"`
  
    Grams int64 `json:"grams"`
  
    Id int64 `json:"id"`
  
    InventoryManagement string `json:"inventory_management"`
  
    InventoryPolicy string `json:"inventory_policy"`
  
    Option1 string `json:"option1"`
  
    Option2 string `json:"option2"`
  
    Option3 string `json:"option3"`
  
    Position int64 `json:"position"`
  
    Price string `json:"price"`
  
    ProductId int64 `json:"product_id"`
  
    RequiresShipping bool `json:"requires_shipping"`
  
    Sku string `json:"sku"`
  
    Taxable bool `json:"taxable"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    InventoryQuantity int64 `json:"inventory_quantity"`
  
    OldInventoryQuantity int64 `json:"old_inventory_quantity"`
  
    ImageId string `json:"image_id"`
  

  
}



