package main

import (
	"fmt"
	"github.com/boourns/go_shopify"
	"log"
	"net/http"
	"os"
)

var appl *shopify.App

// set Callback URL to http://localhost:4000/installed

// this endpoint will be used to handle the oauth callback from Shopify
const defaultRedirect = "http://localhost:4000/install"

func init() {
	var key, secret, redirect string

	if key = os.Getenv("SHOPIFY_API_KEY"); key == "" {
		panic("Set SHOPIFY_API_KEY")
	}

	if secret = os.Getenv("SHOPIFY_API_SECRET"); secret == "" {
		panic("Set SHOPIFY_API_SECRET")
	}

	if redirect = os.Getenv("SHOPIFY_API_REDIRECT"); redirect == "" {
		fmt.Printf("SHOPIFY_API_REDIRECT not set, defaulting to %s", defaultRedirect)
		redirect = defaultRedirect
	}

	// use ngrok to test an embedded app with HTTPS
	appl = &shopify.App{
		RedirectURI: redirect,
		APIKey:      key,
		APISecret:   secret,
	}
}

func serveInstall(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// TODO: session for logged in user

	if len(params["error"]) == 1 {
		log.Printf("Install error: %s", params["error"])
	} else if len(params["code"]) == 1 {

		// auth callback from shopify
		if appl.CheckSignature(r.URL) != true {
			http.Error(w, "Invalid signature", 401)
			log.Printf("Invalid signature from Shopify")
			return
		}

		if len(params["shop"]) != 1 {
			http.Error(w, "Expected 'shop' param", 400)
			log.Printf("Invalid signature from Shopify")
			return
		}

		token, _ := appl.AccessToken(params["shop"][0], params["code"][0])

		fmt.Printf("token is %v", token)

		http.Redirect(w, r, "/settings", 302)

	} else if len(params["install_shop"]) == 1 {
		// install request, redirect to Shopify
		shop := params["install_shop"][0]

		http.Redirect(w, r, appl.AuthorizeURL(shop, "read_themes,write_themes"), 302)
	}
}

// initial page served when visited as embedded app inside Shopify
func serveSettings(w http.ResponseWriter, r *http.Request) {
	// TOOD: save session, render this page for each shop
	http.ServeFile(w, r, "static/settings.html")
}

func main() {
	http.HandleFunc("/install", serveInstall)
	http.HandleFunc("/settings", serveSettings)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/home.html")
	})

	http.ListenAndServe("0.0.0.0:4000", nil)
}
