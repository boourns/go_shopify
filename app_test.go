package shopify

import (
	"net/url"
	"testing"
)

var app App

func init() {
	app = App{APIKey: "asdf", APISecret: "1234", RedirectURI: "http://localhost:4000"}
}

func TestAuthorizeURL(t *testing.T) {
	redir := app.AuthorizeURL("burnsmod.myshopify.com", "read_orders")

	expected := "https://burnsmod.myshopify.com/admin/oauth/authorize?client_id=asdf&redirect_uri=http%3A%2F%2Flocalhost%3A4000&scope=read_orders"

	if redir != expected {
		t.Errorf("Expected %s, got %s", expected, redir)
	}
}

func TestSignatureString(t *testing.T) {
	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=31b9fcfbd98a3650b8523bcc92f8c5d2")
	expected := "1234code=asdfshop=burnsmod.myshopify.comtimestamp=1337178173"

	if output := app.signatureString(u, true); output != expected {
		t.Errorf("expected %s output %s", expected, output)
	}

	if output := app.signatureString(u, false); output != "code=asdfshop=burnsmod.myshopify.comtimestamp=1337178173" {
		t.Errorf("expected %s output %s", expected, output)
	}
}

func TestAdminSignatureOk(t *testing.T) {
	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=bd28a1a098688d8937e991aef3bc80ab")

	if app.AdminSignatureOk(u) != true {
		t.Errorf("signature checking failed")
	}
}

func TestAppProxySignatureOk(t *testing.T) {
	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=6be68200c0175f0f28c9246c9c6fa1cc152e119be597d5df1ea950106c25bf7f")

	if app.AppProxySignatureOk(u) != true {
		t.Errorf("signature checking failed")
	}
}

func TestIgnoreSignature(t *testing.T) {

	a := App{APIKey: "asdf", APISecret: "1234", RedirectURI: "http://localhost:4000", IgnoreSignature: true}

	u, _ := url.Parse("https://app.com/?shop=burnsmod.myshopify.com&code=asdf&timestamp=1337178173&signature=ffff")

	if a.AdminSignatureOk(u) != true {
		t.Errorf("IgnoreSignature didn't work for Admin")
	}

	if a.AppProxySignatureOk(u) != true {
		t.Errorf("IgnoreSignature didn't work for AppProxy")
	}
}
