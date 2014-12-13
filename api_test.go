package shopify

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
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

func TestCreateWebhook(t *testing.T) {
	if !remoteEnabled {
		return
	}

	webhooks, err := api.Webhooks()

	if err != nil {
		fmt.Printf("Err fetching webhooks: %v", err)
	}

	for _, v := range webhooks {
		fmt.Printf("Existing webhook: %#v", v)
	}

	webhook := api.NewWebhook()

	webhook.Address = "https://aaa.ngrok.com/service/hook"
	webhook.Format = "json"
	webhook.Topic = "orders/delete"
	err = webhook.Save()

	if err != nil {
		t.Errorf("Error creating webhook: %v", err)
	}

	fmt.Printf("\n\nwebhooks are %#v\n\n", webhook)
}

func TestNewProduct(t *testing.T) {
	product := api.NewProduct()
	product.Title = "T-shirt"
	product.PublishedAt = time.Now()
	product.ProductType = "shirts"
	err := product.Save()
	if err != nil {
		t.Errorf("Error saving product: %s", err)
	}
	fmt.Printf("New product ID is: %d\n", product.Id)
}
