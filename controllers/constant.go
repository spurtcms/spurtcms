package controllers

import (
	"log"
	"os"
	"regexp"
	"spurt-cms/config"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Limit  int = 10
	MLimit int = 6
)

var (
	TZONE, _           = time.LoadLocation(os.Getenv("TIME_ZONE"))
	CurrentLanugageId  int
	Datelayout         = "02 Jan 2006 03:04 PM"
	LogoPath           string
	ExpantLogoPath     string
	PermissionModRoute = []string{}
	DB                 = config.SetupDB()
	ErrorLog           *log.Logger
	WarnLog            *log.Logger
	Lastlogin          = make(map[int]string)
	Lang_codearr       = []string{"en", "es", "fr", "ger"}
	Token              string
	maxSize            = 8 * 1024 * 1024 // 8MB
	SecretKey          = os.Getenv("JWT_SECRET")
	OwndeskUrl         = os.Getenv("OWNDESK_URL")
	CompanyProfilePath = os.Getenv("COMPANY_PROFILE_PATH")
)

var TenantId int

func Tenant_Id(c *gin.Context) {
	TenantId = c.GetInt("tenant_id")
}

func RemoveSpecialCharacter(str string) string {

	stri := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")

	return stri
}
