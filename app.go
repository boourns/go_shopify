package shopify

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
)

type App struct {
	APIKey          string
	APISecret       string
	RedirectURI     string
	IgnoreSignature bool
}

func (s *App) AuthorizeURL(shop string, scopes string) string {
	var u url.URL
	u.Scheme = "https"
	u.Host = shop
	u.Path = "/admin/oauth/authorize"
	q := u.Query()
	q.Set("client_id", s.APIKey)
	q.Set("scope", scopes)
	q.Set("redirect_uri", s.RedirectURI)
	u.RawQuery = q.Encode()

	return u.String()
}

func (s *App) AdminSignatureOk(u *url.URL) bool {
	if s.IgnoreSignature {
		return true
	}

	params := u.Query()
	signature := params["signature"]
	if signature == nil || len(signature) != 1 {
		return false
	}

	raw := md5.Sum([]byte(s.signatureString(u, true)))
	encrypted := hex.EncodeToString(raw[:])

	return 1 == subtle.ConstantTimeCompare([]byte(encrypted), []byte(signature[0]))
}

func (s *App) AppProxySignatureOk(u *url.URL) bool {
	if s.IgnoreSignature {
		return true
	}

	params := u.Query()
	signature := params["signature"]
	if signature == nil || len(signature) != 1 {
		return false
	}

	mac := hmac.New(sha256.New, []byte(s.APISecret))
	mac.Write([]byte(s.signatureString(u, false)))
	calculated := hex.EncodeToString(mac.Sum(nil))

	return 1 == subtle.ConstantTimeCompare([]byte(signature[0]), []byte(calculated))
}

func (s *App) signatureString(u *url.URL, prependSig bool) string {
	params := u.Query()

	keys := []string{}
	for k, _ := range params {
		if k != "signature" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	input := ""
	if prependSig {
		input = s.APISecret
	}
	for _, k := range keys {
		input = fmt.Sprintf("%s%s=%s", input, k, params[k][0])
	}
	return input
}

type tokenResponse struct {
	AccessToken string `json:access_token`
}

func (s *App) AccessToken(shop string, code string) (string, error) {
	url := fmt.Sprintf("https://%s/admin/oauth/access_token.json", shop)

	data := map[string]string{
		"client_id":     s.APIKey,
		"client_secret": s.APISecret,
		"code":          code,
	}

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	transport := &http.Transport{}
	response, err := transport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	token := map[string]string{}
	err = json.NewDecoder(response.Body).Decode(&token)

	if err != nil {
		return "", err
	}

	if _, ok := token["error"]; ok {
		return "", fmt.Errorf("%s", token["error"])
	}

	if _, ok := token["access_token"]; !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return token["access_token"], nil
}
