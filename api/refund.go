package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Refund struct {
  
    CreatedAt time.Time `json:created_at`
  
    Id int64 `json:id`
  
    Note string `json:note`
  
    RefundLineItems int64 `json:refund_line_items`
  
    Restock string `json:restock`
  
    Transactions string `json:transactions`
  
    UserId int64 `json:user_id`
  
}


// TODO implement Refund.show


