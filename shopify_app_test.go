package shopify_app

import (
	"net/url"
	"testing"
)

var app ShopifyApp

func init() {
	app = ShopifyApp{APIKey: "asdf", APISecret: "1234", RedirectURI: "http://localhost:4000"}
}

func TestAuthorizeURL(t *testing.T) {
	redir := app.AuthorizeURL("burnsmod.myshopify.com", "read_orders")

	expected := "https://burnsmod.myshopify.com/admin/oauth/authorize?client_id=asdf&redirect_uri=http%3A%2F%2Flocalhost%3A4000&scopes=read_orders"

	if redir.String() != expected {
		t.Errorf("Expected %s, got %s", expected, redir.String())
	}
}

func TestSignatureString(t *testing.T) {
	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=31b9fcfbd98a3650b8523bcc92f8c5d2")
	expected := "code=asdfshop=burnsmod.myshopify.comtimestamp=1337178173"

	if output := app.signatureString(u); output != expected {
		t.Errorf("expected %s output %s", expected, output)
	}
}

func TestSignature(t *testing.T) {
	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=fa105044ed00f097f098ce87a62cdc67")

	if app.CheckSignature(u) != true {
		t.Errorf("signature checking failed")
	}
}
