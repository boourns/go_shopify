package shopify

import (
	"fmt"
	"os"
	"testing"
)

var api = &API{}

func init() {
	api.URI = os.Getenv("SHOPIFY_API_HOST")
	api.Token = os.Getenv("SHOPIFY_API_TOKEN")
	api.Secret = os.Getenv("SHOPIFY_API_SECRET")

	if api.URI == "" || api.Token == "" || api.Secret == "" {
		panic("Set SHOPIFY_API_HOST=https://shop.myshopify.com SHOPIFY_API_TOKEN=<api client token> SHOPIFY_API_SECRET=<api permission secret>")
	}
}

func TestAPICanGet(t *testing.T) {
	data, response, err := api.request("/admin/shop.json", "GET", nil)
	if err != nil {
		t.Errorf("Error GET shop.json: %s", err)
	}
	if response != 200 {
		t.Errorf("Unexpected response: %d.  Data: %v", response, data)
	}
}

func TestCanGetIndex(t *testing.T) {
	products, err := api.Product_index()
	if err != nil {
		t.Errorf("Error GET Products: %s", err)
	}
	if len(*products) == 0 {
		t.Errorf("Received 0 products")
	}
	fmt.Printf("Received %d products\n", len(*products))
	fmt.Printf("First product: %v\n", products)
}
