package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"spurt-cms/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v82"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/datatypes"
)

func MasterIntegrationList(c *gin.Context) {
	var (
		limitVal int
		offset   int
	)

	// Pagination
	limit := c.Query("limit")
	pageNo, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if limit == "" {
		limitVal = Limit
	} else {
		limitVal, _ = strconv.Atoi(limit)
	}
	if pageNo > 0 {
		offset = (pageNo - 1) * limitVal
	}

	endURL := os.Getenv("MASTER_INTEGRATION_ENDPOINTURL")
	if endURL == "" {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "MASTER_INTEGRATION_ENDPOINTURL not configured",
		})
		return
	}

	// Make GET request
	req, err := http.NewRequest("GET", endURL, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": err.Error()})
		return
	}

	query := req.URL.Query()
	query.Add("limit", strconv.Itoa(limitVal))
	query.Add("offset", strconv.Itoa(offset))
	query.Add("page", strconv.Itoa(pageNo))
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	_, tabName, _ := ModuleRouteName(c)
	masterConnect := true

	var response MasterIntegrationResponse

	if err != nil {
		fmt.Println("❌ Connection error:", err)
		masterConnect = false
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("❌ Invalid response:", resp.StatusCode)
			masterConnect = false
		} else {
			bodyBytes, _ := io.ReadAll(resp.Body)
			if err := json.Unmarshal(bodyBytes, &response); err != nil {
				fmt.Println("❌ JSON decode error:", err)
				fmt.Println("Response:", string(bodyBytes))
				masterConnect = false
			} else {
				fmt.Printf("✅ Parsed %d integrations\n", len(response.IntegrationList))
			}
		}
	}

	if !masterConnect {
		response = MasterIntegrationResponse{
			BlockCount:      0,
			IntegrationList: []models.IntegrationView{},
		}
	}

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	// Render HTML
	c.HTML(http.StatusOK, "masterintegration.html", gin.H{
		"csrf":        csrf.GetToken(c),
		"Menu":        menu,
		"Count":       response.BlockCount,
		"title":       "Integration",
		"linktitle":   "Integration",
		"translate":   translate,
		"Tabmenu":     tabName,
		"Integration": response.IntegrationList,
	})
}

// func MasterIntegrationList(c *gin.Context) {
// 	var (
// 		limt   int
// 		offset int
// 	)
// 	limit := c.Query("limit")
// 	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	if limit == "" {
// 		limt = Limit
// 	} else {
// 		limt, _ = strconv.Atoi(limit)
// 	}
// 	if pageno != 0 {
// 		offset = (pageno - 1) * limt
// 	}

// 	endurl := os.Getenv("MASTER_INTEGRATION_ENDPOINTURL")

// 	req, err := http.NewRequest("GET", endurl, nil)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
// 		return
// 	}
// 	query := req.URL.Query()

// 	query.Add("limit", strconv.Itoa(limt))
// 	query.Add("offset", strconv.Itoa(offset))
// 	query.Add("page", strconv.Itoa(pageno))
// 	req.URL.RawQuery = query.Encode()
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	_, TabName, _ := ModuleRouteName(c)

// 	masterconnect := true

// 	if err != nil || resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error connecting to master server:", err)
// 		masterconnect = false
// 	} else {
// 		defer resp.Body.Close()
// 	}

// 	var responseData ResponseData

// 	fmt.Println("masterconnect:",masterconnect)
// 	if masterconnect {
// 		bodyBytes, err := io.ReadAll(resp.Body)
// 		if err == nil {
// 			fmt.Println("Error response:", err)
// 			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
// 			err = json.NewDecoder(resp.Body).Decode(&responseData)
// 			if err != nil {
// 				masterconnect = false
// 			}
// 		} else {
// 			masterconnect = false
// 		}
// 	}

// 	fmt.Println("responseData::",responseData.IntegrationList)

// 	if !masterconnect {
// 		responseData = ResponseData{
// 			IntegrationList: []models.TblMstrIntegration{},
// 			BlockCount:      0,
// 		}
// 	}

// 	translate, _ := TranslateHandler(c)

// 	menu := NewMenuController(c)

// 	c.HTML(200, "masterintegration.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "Count": responseData.BlockCount, "title": "Integration", "linktitle": "Integration", "translate": translate, "Tabmenu": TabName, "Integration": responseData.IntegrationList})

// }

func PaymentIntgrationList(c *gin.Context) {
	var viewData []models.IntegrationView

	Integration, Total_Integration, _ := models.ListIntegration(TenantId, "")
	for _, item := range Integration {
		var wrapper struct {
			Credentials models.Credentials `json:"credentials"`
		}

		if err := json.Unmarshal(item.Credentials, &wrapper); err != nil {
			fmt.Println("❌ Unmarshal error:", err)
			continue
		}

		stoarge := wrapper.Credentials

		viewData = append(viewData, models.IntegrationView{
			Id:              item.Id,
			Credentials:     stoarge,
			IsActive:        item.IsActive,
			IntegrationType: item.IntegrationType,
		})
	}

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	c.HTML(200, "integration.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "title": "Integration", "linktitle": "Integration", "translate": translate, "Tabmenu": TabName, "IntegrationList": viewData, "Count": Total_Integration})

}

func AddIntegrationToCollection(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("Id"))
	TenantId := c.GetString("tenant_id") // Make sure you’ve set this somewhere (middleware or session)
	userId := c.GetInt("userid")

	endURL := os.Getenv("MASTER_INTEGRATION_ENDPOINTURL")
	queryURL := fmt.Sprintf("%s?id=%d", endURL, Id)

	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("❌ Connection error:", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to master API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("❌ Invalid response:", resp.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Master API returned an invalid response"})
		return
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	// ✅ Expecting JSON like: { "integrationdetail": { ... } }
	type MasterIntegrationDetailResponse struct {
		IntegrationDetail models.IntegrationView `json:"integrationdetail"`
	}

	var response MasterIntegrationDetailResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		fmt.Println("❌ JSON decode error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse master integration JSON"})
		return
	}

	masterData := response.IntegrationDetail

	// ---- CHECK IF ALREADY EXISTS ----
	exists, _ := models.CheckInterationReapt(masterData.Credentials.GatewayName, TenantId)
	fmt.Println("exists", exists)
	if exists {
		c.SetCookie("Alert-msg", "Integrationreapt", 3600, "", "", false, false)
		c.Redirect(301, "/admin/integration/master")
		return
	}

	// ---- COMBINE PAYMENT & S3 JSON ----
	combined := map[string]interface{}{
		"credentials": masterData.Credentials,
	}

	combinedJSON, _ := json.Marshal(combined)

	// ---- CREATE LOCAL INTEGRATION ----
	now := time.Now().UTC()
	newIntegration := models.TblIntegration{
		Credentials:     datatypes.JSON(combinedJSON),
		IntegrationType: masterData.IntegrationType,
		IsActive:        0,
		CreatedBy:       userId,
		CreatedOn:       now,
		TenantId:        TenantId,
	}

	if err := models.AddtoCollection(newIntegration); err != nil {
		log.Println("❌ Error adding to collection:", err)
		c.AbortWithError(500, err)
		return
	}

	// ---- SUCCESS RESPONSE ----
	c.SetCookie("get-toast", "Collection Added Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/integration")
}

func IntegrationEdit(c *gin.Context) {

	IntegrationId, _ := strconv.Atoi(c.DefaultQuery("id", ""))

	EditIntegration, _ := models.EditIntegration(IntegrationId, TenantId)

	c.JSON(200, gin.H{"Integration": EditIntegration})

}

func GetStripeValue(c *gin.Context) {

	Getstripe, _ := models.GetStripeIntegration( TenantId)

	c.JSON(200, gin.H{"Integration": Getstripe})

}

// func IntegrationEdit(c *gin.Context) {

// 	IntegrationId, _ := strconv.Atoi(c.DefaultQuery("id", ""))

// 	endurl := os.Getenv("MASTER_INTEGRATION_ENDPOINTURL")

// 	req, err := http.NewRequest("GET", endurl, nil)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
// 		return
// 	}
// 	query := req.URL.Query()

// 	query.Add("id", strconv.Itoa(IntegrationId))
// 	req.URL.RawQuery = query.Encode()
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	masterconnect := true

// 	if err != nil || resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error connecting to master server:", err)
// 		masterconnect = false
// 	} else {
// 		defer resp.Body.Close()
// 	}

// 	type integrationres struct {
// 		Integraiondetail models.TblMstrIntegration `json:"integrationdetail"`
// 	}
// 	var responseData integrationres
// 	var bodyBytes []byte
// 	if masterconnect {
// 		bodyBytes, err = io.ReadAll(resp.Body)
// 		if err != nil {
// 			masterconnect = false
// 			fmt.Println("Error reading response body:", err)
// 		} else {

// 			err = json.Unmarshal(bodyBytes, &responseData)
// 			if err != nil {
// 				masterconnect = false
// 				fmt.Println("Error decoding JSON:", err)
// 			}
// 		}
// 	}

// 	if !masterconnect {
// 		responseData = integrationres{
// 			Integraiondetail: models.TblMstrIntegration{},
// 		}
// 	}

// 	// EditIntegration, _ := models.EditIntegration(IntegrationId, TenantId)

// 	c.JSON(200, gin.H{"Integration": responseData.Integraiondetail})

// }

func IntegrationManage(c *gin.Context) {
	IntegrationType := c.PostForm("IntegrationType")
	fmt.Println("IntegrationType", IntegrationType)
	// Prepare credentials map
	var creds map[string]interface{}

	if IntegrationType == "Payment" {
		creds = map[string]interface{}{
			"access_key":       c.PostForm("ClientId"),     // optional: read from form
			"secret_key":       c.PostForm("ClientSecret"), // optional: read from form
			"gateway_name":     c.PostForm("GatewayName"),
			"gateway_desc":     c.PostForm("GatewayDesc"),
			"dollar_to_rupees": c.PostForm("DollartoRupees"),
		}
	} else { // Storage
		creds = map[string]interface{}{
			"region":       c.PostForm("Region"),
			"access_key":   c.PostForm("AccessKeyId"),
			"secret_key":   c.PostForm("SecretKey"),
			"bucket_name":  c.PostForm("BucketName"),
			"gateway_name": c.PostForm("GatewayName"),
			"gateway_desc": c.PostForm("GatewayDesc"),
		}
	}

	// Marshal credentials map to JSON
	CredentialsSON, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("❌ Error marshalling credentials:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process credentials"})
		return
	}

	IntegrationId, _ := strconv.Atoi(c.PostForm("IntegrationId"))
	userid := c.GetInt("userid")
	isactive, _ := strconv.Atoi(c.PostForm("integration_activestatus"))

	updatetime := time.Now().UTC()

	// Build integration struct
	UpdateIntegration := models.TblIntegration{
		Id:              IntegrationId,
		Credentials:     datatypes.JSON(CredentialsSON),
		IntegrationType: IntegrationType,
		IsActive:        isactive,
		ModifiedOn:      updatetime,
		ModifiedBy:      userid,
	}
	// Call model function to save/update
	if err := models.ManageIntegration(UpdateIntegration, TenantId); err != nil {
		fmt.Println("❌ Error updating integration:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update integration"})
		fmt.Println("err", err)
		return
	}

	// Set success cookie and redirect
	c.SetCookie("get-toast", "Integration Updated Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/integration")
}

func HandleStripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusServiceUnavailable, "Error reading body")
		return
	}

	var event stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err == nil {
			fmt.Println("Payment Success:", session.ID)
			if err != nil {
				log.Fatal("Create Subscription Error :", err)
				c.AbortWithError(500, err)
			}
			// storeToDB(session.ID, "success", session.CustomerEmail)
		}

	case "payment_intent.payment_failed":
		var pi stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &pi); err == nil {
			fmt.Println("Payment Failed:", pi.ID)
			// storeToDB(pi.ID, "failed", pi.ReceiptEmail)
		}

	case "payment_intent.processing":
		var pi stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &pi); err == nil {
			fmt.Println("Payment Pending:", pi.ID)
			// storeToDB(pi.ID, "pending", pi.ReceiptEmail)
		}

	default:
		fmt.Println("Unhandled event type:", event.Type)
	}

	c.Status(http.StatusOK)
}
