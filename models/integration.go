package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// type TblIntegration struct {
// 	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
// 	ClientId     string    `gorm:"type:character varying"`
// 	ClientSecret string    `gorm:"type:character varying"`
// 	CoverImage   string    `gorm:"type:character varying"`
// 	GatewayName  string    `gorm:"type:character varying"`
// 	GatewayDesc  string    `gorm:"type:character varying"`
// 	IsActive     int       `gorm:"type:integer"`
// 	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
// 	CreatedBy    int       `gorm:"type:integer"`
// 	IsDeleted    int       `gorm:"type:integer"`
// 	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	DeletedBy    int       `gorm:"DEFAULT:NULL"`
// 	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
// 	TenantId     string    `gorm:"type:character varying"`
// }

// type TblIntegration struct {
// 	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
// 	ClientId     string    `gorm:"type:character varying"`
// 	ClientSecret string    `gorm:"type:character varying"`
// 	CoverImage   string    `gorm:"type:character varying"`
// 	GatewayName  string    `gorm:"type:character varying"`
// 	GatewayDesc  string    `gorm:"type:character varying"`
// 	IsActive     int       `gorm:"type:integer"`
// 	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
// 	CreatedBy    int       `gorm:"type:integer"`
// 	IsDeleted    int       `gorm:"type:integer"`
// 	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	DeletedBy    int       `gorm:"DEFAULT:NULL"`
// 	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
// 	TenantId     string    `gorm:"type:character varying"`
// }

type TblIntegration struct {
	Id              int            `gorm:"primaryKey;auto_increment;type:serial" json:"id"`
	Credentials     datatypes.JSON `gorm:"type:jsonb" json:"credentials"`
	IntegrationType string         `gorm:"type:character varying" json:"integration_type"`
	IsActive        int            `gorm:"type:integer" json:"is_active"`
	CreatedOn       time.Time      `gorm:"type:timestamp without time zone" json:"created_on"`
	CreatedBy       int            `gorm:"type:integer" json:"created_by"`
	IsDeleted       int            `gorm:"type:integer" json:"is_deleted"`
	DeletedOn       time.Time      `gorm:"type:timestamp without time zone;DEFAULT:NULL" json:"deleted_on"`
	DeletedBy       int            `gorm:"DEFAULT:NULL" json:"deleted_by"`
	ModifiedOn      time.Time      `gorm:"type:timestamp without time zone;DEFAULT:NULL" json:"modified_on"`
	ModifiedBy      int            `gorm:"DEFAULT:NULL" json:"modified_by"`
	TenantId        string         `gorm:"type:character varying" json:"tenant_id"`
}

func ListIntegration(tenantid string, Keyword string) ([]TblIntegration, int64, error) {
	var Total_Integration int64

	var Integration []TblIntegration
	query := DB.Table("tbl_integrations").Where("tenant_id=?", tenantid).Order("id asc")

	// if Keyword != "" {
	// 	query = query.Debug().
	// 		Where("LOWER(TRIM(tbl_integrations.gateway_name)) LIKE LOWER(TRIM(?))", "%"+Keyword+"%")

	// }

	query.Count(&Total_Integration).Find(&Integration)

	return Integration, Total_Integration, nil
}

func EditIntegration(Id int, tenantid string) (TblIntegration, error) {
	var Integration TblIntegration
	if err := DB.Table("tbl_integrations").Where("id=? and tenant_id=?", Id, tenantid).First(&Integration).Error; err != nil {

		return TblIntegration{}, err
	}

	return Integration, nil

}

func GetStripeIntegration(tenantid string) (TblIntegration, error) {
	var Integration TblIntegration
	if err := DB.Table("tbl_integrations").Where("tenant_id=? and is_active='1'", tenantid).Where("credentials->'credentials'->>'gateway_name' = ?", "Stripe").First(&Integration).Error; err != nil {

		return TblIntegration{}, err
	}

	return Integration, nil

}

func ManageIntegration(integration TblIntegration, tenantId string) error {
	// Prepare the individual JSON fields to update
	// Example: update gateway_name and gateway_desc inside credentials JSON
	credMap := map[string]interface{}{}
	if err := json.Unmarshal(integration.Credentials, &credMap); err != nil {
		return fmt.Errorf("failed to unmarshal credentials JSON: %w", err)
	}

	for key, value := range credMap {
		// Convert value to JSON string
		valueBytes, _ := json.Marshal(value)
		if err := DB.Exec(`
			UPDATE tbl_integrations
			SET credentials = jsonb_set(credentials, '{credentials,"`+key+`"}', ?::jsonb, true),
				modified_on = ?,
				modified_by = ?
			WHERE id = ? AND tenant_id = ?`,
			string(valueBytes), integration.ModifiedOn, integration.ModifiedBy, integration.Id, tenantId).Error; err != nil {
			return err
		}
	}

	// Update non-JSON fields if needed
	if err := DB.Table("tbl_integrations").
		Where("id = ? AND tenant_id = ?", integration.Id, tenantId).
		Updates(map[string]interface{}{
			"integration_type": integration.IntegrationType,
			"is_active":        integration.IsActive,
		}).Error; err != nil {
		return err
	}

	return nil
}

func AddtoCollection(Integration TblIntegration) error {
	if err := DB.Table("tbl_integrations").Create(&Integration).Error; err != nil {
		return err
	}
	return nil
}

func CheckInterationReapt(name string, tenantid string) (bool, error) {
	var integration TblIntegration
	err := DB.Table("tbl_integrations").
		Where("credentials->'credentials'->>'gateway_name' = ? AND tenant_id = ?", name, tenantid).Debug().
		First(&integration).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func StripePayment(SubscriptionName string, SubscriptionAmount int64, SuccessURL string, CancelURL string, PaymentType string, MemberId string, MemberEmail string) (string, error) {
	stripe.Key = os.Getenv("Stripe_Secret_Key")

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(SubscriptionName),
					},
					UnitAmount: stripe.Int64(SubscriptionAmount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(SuccessURL),
		CancelURL:  stripe.String(CancelURL),
		Metadata: map[string]string{
			"user_id":      MemberId,
			"subscription": SubscriptionName,
			"email":        MemberEmail,
		},
	}

	s, err := session.New(params)
	if err != nil {
		return "", err
	}

	return s.URL, nil
}

func CreatePayPalOrderAndGetApprovalLink(amount float64) (string, error) {
	accessToken, _ := getAccessToken()
	if accessToken == "" {
		return "", fmt.Errorf("failed to get PayPal access token")
	}

	body := map[string]interface{}{
		"intent": "CAPTURE",
		"purchase_units": []map[string]interface{}{
			{
				"amount": map[string]string{
					"currency_code": "USD",
					"value":         fmt.Sprintf("%.2f", amount),
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api-m.sandbox.paypal.com/v2/checkout/orders", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("create request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("paypal request error: %w", err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode error: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("paypal create order failed: %v", result)
	}

	// Extract approval link
	links := result["links"].([]interface{})
	for _, l := range links {
		link := l.(map[string]interface{})
		if link["rel"] == "approve" {
			return link["href"].(string), nil
		}
	}

	return "", fmt.Errorf("approval link not found")
}

func getAccessToken() (string, error) {

	clientID := os.Getenv("Access_Client")
	secret := os.Getenv("Access_Secret")

	if clientID == "" || secret == "" {
		return "", fmt.Errorf("missing PayPal credentials in environment variables")
	}

	data := "grant_type=client_credentials"
	req, err := http.NewRequest("POST", "https://api-m.sandbox.paypal.com/v1/oauth2/token", strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.SetBasicAuth(clientID, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("PayPal token error [%d]: %s", res.StatusCode, string(bodyBytes))
	}

	var tokenData struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(bodyBytes, &tokenData); err != nil {
		return "", fmt.Errorf("failed to parse token JSON: %w", err)
	}

	if tokenData.AccessToken == "" {
		return "", fmt.Errorf("access token is empty")
	}

	return tokenData.AccessToken, nil
}

func GetS3Credential(tenantid string, gatewayName string) (TblIntegration, error) {
	var integrations TblIntegration

	// JSONB query
	if err := DB.Table("tbl_integrations").
		Where("tenant_id = ? AND credentials->'credentials'->>'gateway_name' = ?", tenantid, gatewayName).
		Order("id ASC").
		First(&integrations).Error; err != nil {
		return TblIntegration{}, err
	}

	return integrations, nil
}
