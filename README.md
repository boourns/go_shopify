go_shopify
==========

Idiomatic Shopify API client & app helper in Go

API Examples
========

__Initialize an api client and fetch products__

```go

import (
  "github.com/boourns/go_shopify"
)

func main() {
  api = shopify.API{
    URI: "https://shopname.myshopify.com/admin",
    Token: "(api client token)",
    Secret: "(api client secret for this shop)",
  }

  products := api.Products()
  // or
  product := api.Product(12345)
}
```

__ Create a new Product__
```go
product := api.NewProduct()
product.Title = "T-shirt"
product.PublishedAt = time.Now()
product.ProductType = "shirts"
err := product.Save()
if err != nil {
  fmt.Printf("Error saving product: %s", err)
}
fmt.Printf("New product ID is: %d\n", product.Id)  
```

Done
====
- App install flow (see example/main.go)
- Check signatures for admin and API proxy requests coming from Shopify
- store session for users
- store API keys for installed shops

TODO
====
- API client
