package shopify_app

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
)

type ShopifyApp struct {
	APIKey      string
	APISecret   string
	RedirectURI string
}

func (s *ShopifyApp) AuthorizeURL(shop string, scopes string) url.URL {
	var u url.URL
	u.Scheme = "https"
	u.Host = shop
	u.Path = "/admin/oauth/authorize"
	q := u.Query()
	q.Set("client_id", s.APIKey)
	q.Set("scopes", scopes)
	q.Set("redirect_uri", s.RedirectURI)
	u.RawQuery = q.Encode()

	return u
}

func (s *ShopifyApp) CheckSignature(u *url.URL) bool {
	params := u.Query()
	signature := params["signature"]
	if signature == nil || len(signature) != 1 {
		return false
	}

	raw := md5.Sum([]byte(s.signatureString(u)))
	encrypted := hex.EncodeToString(raw[:])

	return 1 == subtle.ConstantTimeCompare([]byte(encrypted), []byte(signature[0]))
}

func (s *ShopifyApp) signatureString(u *url.URL) string {
	params := u.Query()

	keys := []string{}
	for k, _ := range params {
		if k != "signature" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	input := ""
	for _, k := range keys {
		input = fmt.Sprintf("%s%s=%s", input, k, params[k][0])
	}
	return input
}

func (s *ShopifyApp) AccessTokenFromCode(code string) string {
	return ""
}
