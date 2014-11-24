package main

import (
	"fmt"
	"github.com/boourns/go_shopify"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
)

var app *shopify.App

// Change these to actual secret keys before use
var store = sessions.NewCookieStore([]byte("this-is-a-dummy-authentication-key32"), []byte("this-is-a-dummy-encryption-key32"))

var tokens map[string]string

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
	app = &shopify.App{
		RedirectURI: redirect,
		APIKey:      key,
		APISecret:   secret,
	}

	tokens = map[string]string{}
}

func getSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "shopify_app")
	if err != nil {
		panic(err)
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	return session
}

func serveInstall(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if len(params["error"]) == 1 {
		log.Printf("Install error: %s", params["error"])
	} else if len(params["code"]) == 1 {

		// auth callback from shopify
		if app.AdminSignatureOk(r.URL) != true {
			http.Error(w, "Invalid signature", 401)
			log.Printf("Invalid signature from Shopify")
			return
		}

		if len(params["shop"]) != 1 {
			http.Error(w, "Expected 'shop' param", 400)
			log.Printf("Invalid signature from Shopify")
			return
		}

		shop := params["shop"][0]
		token, _ := app.AccessToken(shop, params["code"][0])

		// persist this token
		tokens[shop] = token

		// log in user
		session := getSession(r)
		session.Values["current_shop"] = shop
		err := session.Save(r, w)
		if err != nil {
			panic(err)
		}

		log.Printf("logged in as %s, redirecting to admin", shop)

		http.Redirect(w, r, "/admin", 302)

	} else if len(params["install_shop"]) == 1 {
		// install request, redirect to Shopify
		shop := params["install_shop"][0]
		log.Printf("starting oauth flow")

		http.Redirect(w, r, app.AuthorizeURL(shop, "read_themes,write_themes"), 302)
	}
}

func serveAppProxy(w http.ResponseWriter, r *http.Request) {
	if app.AppProxySignatureOk(r.URL) {
		http.ServeFile(w, r, "static/app_proxy.html")
	} else {
		http.Error(w, "Unauthorized", 401)
	}
}

// initial page served when visited as embedded app inside Shopify admin
func serveAdmin(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	params := r.URL.Query()

	// signed request from Shopify?
	if app.AdminSignatureOk(r.URL) {
		log.Printf("signed request!")
		session.Values["current_shop"] = params["shop"][0]
		session.Save(r, w)
	} else if _, ok := session.Values["current_shop"]; !ok {
		log.Printf("no current_shop")
		// not logged in and not signed request
		http.Error(w, "Unauthorized", 401)
		return
	}

	shop, _ := session.Values["current_shop"].(string)

	// if we don't have an access token for the shop, obtain one now.
	if _, ok := tokens[shop]; !ok {
		http.Redirect(w, r, app.AuthorizeURL(shop, "read_themes,write_themes"), 302)
		return
	}

	// they're logged in
	http.ServeFile(w, r, "static/admin.html")
}

func main() {
	http.HandleFunc("/install", serveInstall)
	http.HandleFunc("/admin", serveAdmin)
	http.HandleFunc("/app_proxy/", serveAppProxy)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/home.html")
	})

	http.ListenAndServe("0.0.0.0:4000", context.ClearHandler(http.DefaultServeMux))
}
