package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type Shop struct {
  
    Address1 string `json:"address1"`
  
    City string `json:"city"`
  
    Country string `json:"country"`
  
    CountryCode string `json:"country_code"`
  
    CountryName string `json:"country_name"`
  
    CreatedAt time.Time `json:"created_at"`
  
    CustomerEmail string `json:"customer_email"`
  
    Currency string `json:"currency"`
  
    Domain string `json:"domain"`
  
    Email string `json:"email"`
  
    GoogleAppsDomain string `json:"google_apps_domain"`
  
    GoogleAppsLoginEnabled string `json:"google_apps_login_enabled"`
  
    Id int64 `json:"id"`
  
    Latitude time.Time `json:"latitude"`
  
    Longitude string `json:"longitude"`
  
    MoneyFormat string `json:"money_format"`
  
    MoneyWithCurrencyFormat string `json:"money_with_currency_format"`
  
    MyshopifyDomain string `json:"myshopify_domain"`
  
    Name string `json:"name"`
  
    PlanName string `json:"plan_name"`
  
    DisplayPlanName string `json:"display_plan_name"`
  
    PasswordEnabled string `json:"password_enabled"`
  
    Phone string `json:"phone"`
  
    Province string `json:"province"`
  
    ProvinceCode string `json:"province_code"`
  
    Public string `json:"public"`
  
    ShopOwner string `json:"shop_owner"`
  
    Source string `json:"source"`
  
    TaxShipping string `json:"tax_shipping"`
  
    TaxesIncluded string `json:"taxes_included"`
  
    CountyTaxes string `json:"county_taxes"`
  
    Timezone time.Time `json:"timezone"`
  
    Zip time.Time `json:"zip"`
  
    HasStorefront string `json:"has_storefront"`
  
  api *API
}


func (api *API) Shop(id int64) (*Shop, error) {
  endpoint := fmt.Sprintf("/admin/shops/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil, nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]Shop{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["shop"]

	if err != nil {
		return nil, err
  }

  result.api = api

  return &result, nil
}



