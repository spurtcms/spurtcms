package controllers

import (
	"fmt"
	"spurt-cms/models"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/datatypes"
)

func MediaSettings(c *gin.Context) {

	menu := NewMenuController(c)

	ModuleName, TabName, _ := ModuleRouteName(c)

	Storagetype, err := models.GetStorageValue(TenantId)

	if err != nil {
		fmt.Println(err)
	}

	var aws models.AWS

	if Storagetype.Aws != nil {

		aws.AccessId = Storagetype.Aws["AccessId"].(string)
		aws.AccessKey = Storagetype.Aws["AccessKey"].(string)
		aws.Region = Storagetype.Aws["Region"].(string)
		aws.BucketName = Storagetype.Aws["BucketName"].(string)

	}

	var azure models.Azure

	if Storagetype.Azure != nil {
		azure.StorageAccount = Storagetype.Azure["AzureAccount"].(string)
		azure.AccountKey = Storagetype.Azure["AzureKey"].(string)
		azure.ContainerName = Storagetype.Azure["AzureContainer"].(string)
	}

	translate, _ := TranslateHandler(c)

	c.HTML(200, "media_settings.html", gin.H{"Menu": menu, "Mediamenu": true, "title": ModuleName, "Tabmenu": TabName, "HeadTitle": "Media", "csrf": csrf.GetToken(c), "aws": aws, "azure": azure, "local": Storagetype.Local, "Storage": Storagetype, "translate": translate})
}

func MediaStorageUpdate(c *gin.Context) {

	selectedtype := c.Request.PostFormValue("type")

	local := c.PostForm("local")
	awsid := c.PostForm("awsid")
	awskey := c.PostForm("awskey")
	awsregion := c.PostForm("awsregion")
	awsbucketname := c.PostForm("awsbucketname")
	azureacc := c.PostForm("azureacc")
	azurekey := c.PostForm("azurekey")
	azurecontainer := c.PostForm("azurecontainer")

	var Stype models.TblStorageType

	Stype.Local = local
	Stype.Aws = datatypes.JSONMap{"AccessId": awsid, "AccessKey": awskey, "Region": awsregion, "BucketName": awsbucketname}
	Stype.Azure = datatypes.JSONMap{"AzureAccount": azureacc, "AzureKey": azurekey, "AzureContainer": azurecontainer}
	Stype.SelectedType = selectedtype

	_, err := models.UpdateStorageType(Stype, TenantId)

	if err != nil {
		fmt.Println(err)
		c.SetCookie("get-toast", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/media/settings/")
		return
	}

	c.SetCookie("get-toast", "Media Settings Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/media/settings/")

}
