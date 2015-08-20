package shopify



type Option struct {
  
    Id int64 `json:"id"`
  
    Name string `json:"name"`
  
    Position int64 `json:"position"`
  
    ProductId int64 `json:"product_id"`
  
    Values []string `json:"values"` 
    
  
}



