package shopify

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var api API
var remoteEnabled = false

func init() {
	if os.Getenv("SHOPIFY_API_TOKEN") != "" && os.Getenv("SHOPIFY_API_SECRET") != "" && os.Getenv("SHOPIFY_API_HOST") != "" {
		remoteEnabled = true
		api = API{
			URI:    os.Getenv("SHOPIFY_API_HOST"),
			Token:  os.Getenv("SHOPIFY_API_TOKEN"),
			Secret: os.Getenv("SHOPIFY_API_SECRET"),
		}
	} else {
		log.Printf("Remote tests disabled, set SHOPIFY_API_KEY, SHOPIFY_API_SECRET, SHOPIFY_API_HOST")
	}
}

func TestReadProducts(t *testing.T) {
	if !remoteEnabled {
		return
	}

	products, err := api.Product(389374712)

	if err != nil {
		t.Errorf("Error fetching products: %v", err)
	}

	fmt.Printf("\n\nproducts are %#v\n\n", products)
}
