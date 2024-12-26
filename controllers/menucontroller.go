package controllers

import (
	"log"
	"os"
	"spurt-cms/config"
	"spurt-cms/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubModule struct {
	Id         int
	ModuleName string
	IconPath   string
	Routes     []URL
	// FullAccessPermission bool
	Action bool
}

type URL struct {
	Id          int
	DisplayName string
	RouteName   string
}

type MenuMod struct {
	Id         int
	ParentId   int
	ModuleName string
	IconPath   string
	Routes     []URL // this is for single menu multiple permissions arr
	HrefRoute  URL
	Route      string      //this is for a tag href route
	SubModule  []SubModule // this is for submodules
	// FullAccessPermission bool
	EmptyCheck bool //this is flg for mainmenu hide if submodule is empty
}

type TablePageType struct {
	Id         int
	Title      string
	Slug       string
	CreatedBy  int
	CreatedOn  time.Time
	ModifiedBy int
	ModifiedOn time.Time
	DeletedBy  int
	DeletedOn  time.Time
	IsDeleted  int
}

type USR struct {
	TblModule         []MenuMod
	Name              string
	NameLength        int
	LimitedLengthName string
	RoleName          string
	ProfileImagePath  string
	Languages         []models.TblLanguage
	DefaultLanguage   models.TblLanguage
	CurrentLanguage   models.TblLanguage
	Personalize       models.TblUserPersonalize
	Footer            string
	NameString        string
	Expand            bool
	LogoPath          string
	ExpandLogoPath    string
	TblPageTypes      []TablePageType
}

type TblModulePermission struct {
	Id                   int
	RouteName            string
	DisplayName          string
	SlugName             string
	Description          string
	ModuleId             int
	CreatedBy            int
	CreatedOn            time.Time
	CreatedDate          string    `gorm:"-"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"DEFAULT:NULL"`
	ModuleName           string    `gorm:"<-:false"`
	FullAccessPermission int
	ParentId             int
	AssignPermission     int
	BreadcrumbName       string
	TblRolePermission    []TblRolePermission `gorm:"<-:false; foreignKey:PermissionId"`
	OrderIndex           int
}

type TblRolePermission struct {
	Id           int
	RoleId       int
	PermissionId int
	CreatedBy    int
	CreatedOn    time.Time
	CreatedDate  string `gorm:"<-:false"`
}

type TblModule struct {
	Id                   int
	ModuleName           string
	IsActive             int
	CreatedBy            int
	CreatedOn            time.Time
	CreatedDate          string `gorm:"<-:false"`
	DefaultModule        int
	ParentId             int
	IconPath             string
	OrderIndex           int
	TblModulePermission  []TblModulePermission `gorm:"<-:false; foreignKey:ModuleId"`
	Description          string
	DisplayName          string `gorm:"column:display_name;<-:false"`
	RouteName            string `gorm:"column:route_name;<-:false"`
	RouteParentId        int    `gorm:"column:parent_id;<-:false"`
	FullAccessPermission int
	GroupFlg             int
}

func GetLanguageList() []models.TblLanguage {

	var languagelist []models.TblLanguage

	err := models.FetchAllLanguage(&languagelist, TenantId)
	if err != nil {
		ErrorLog.Printf("Fectch all language error: %s", err)
		return []models.TblLanguage{}
	}

	for _, langlist := range languagelist {
		_, err := os.Stat(langlist.JsonPath)
		if err != nil {
			log.Println(err)
			return []models.TblLanguage{}
		}
	}

	return languagelist

}

func MainMenu() []TblModule {

	var parentModules []TblModule

	GetAllParentModule(&parentModules) //get parent default main module only

	return parentModules

}

func GetPageType() []TablePageType {

	var pagetypes []TablePageType

	GetDefaultPagetypes(&pagetypes) //get parent default main module only

	return pagetypes

}

func TabMenu(parentid int) []TblModule {

	var subModules []TblModule

	GetAllSubModule(&subModules, parentid)

	return subModules

}

func NewMenuController(c *gin.Context) USR {

	userid := c.GetInt("userid")

	var (
		permissionid     []int
		Final            []MenuMod
		personalize_data models.TblUserPersonalize
		firstn           string
		lastn            string
		expandflg        bool
		currentLanguage  models.TblLanguage
	)

	user, _, _ := NewTeamWP.GetUserById(userid, []int{})

	rolepermission, err := GetPermissionIDS(user.RoleId)

	pagetype := GetPageType()

	if err != nil {

		log.Println(err)

	}

	for _, val := range rolepermission {
		permissionid = append(permissionid, val.PermissionId)
	}

	LeftMain := MainMenu()

	for _, val := range LeftMain {

		var fin MenuMod
		fin.Id = val.Id
		fin.ModuleName = val.ModuleName
		fin.IconPath = val.IconPath

		TabManu := TabMenu(val.Id)

		var Suball []SubModule

		if user.RoleId != 1 && user.RoleId != 2 { //not admin users
			for _, tab := range TabManu {

				var sub SubModule
				sub.Id = tab.Id
				sub.ModuleName = tab.ModuleName
				sub.IconPath = tab.IconPath

				var (
					Url  []URL
					rout []TblModulePermission
					flg  bool
				)

				GetModulePermissions(&rout, tab.Id, permissionid)

				for _, url := range rout {
					if strings.Contains(url.RouteName, "/channel/entrylist/") {
						flg = true
						break
					}
				}

				for _, url := range rout {
					if url.ModuleId == tab.Id {
						if url.RouteName != "/channel/entrylist" {

							var urll URL
							if strings.Contains(url.RouteName, "/channel/entrylist/") {

								url.RouteName = "/channel/entrylist"
							}
							urll.Id = url.Id
							urll.DisplayName = url.DisplayName
							urll.RouteName = url.RouteName
							Url = append(Url, urll)

						} else if flg && url.RouteName == "/channel/entrylist" {

							var urll URL
							urll.Id = url.Id
							urll.DisplayName = url.DisplayName
							urll.RouteName = url.RouteName
							Url = append(Url, urll)

						}
					}
				}

				if flg {

					var urll URL
					urll.DisplayName = "Entries"
					urll.RouteName = "/channel/entrylist"
					Url = append(Url, urll)
				}

				sub.Routes = Url
				Suball = append(Suball, sub)
			}
		} else {

			for _, tab := range TabManu {

				if tab.ParentId == val.Id {

					var sub SubModule
					sub.Id = tab.Id
					sub.ModuleName = tab.ModuleName
					sub.IconPath = tab.IconPath

					var rout []TblModulePermission
					GetModulePermissions(&rout, tab.Id, []int{})

					var Url []URL
					for _, url := range rout {
						var urll URL
						urll.Id = url.Id
						urll.DisplayName = url.DisplayName
						urll.RouteName = url.RouteName
						Url = append(Url, urll)
					}

					sub.Routes = Url
					Suball = append(Suball, sub)

				}
			}
		}

		if len(Suball) > 0 && len(Suball[0].Routes) > 0 {

			fin.Route = Suball[0].Routes[0].RouteName

		}

		for _, val := range Suball {
			if len(val.Routes) != 0 {
				fin.Route = val.Routes[0].RouteName
				break
			}

		}

		fin.SubModule = Suball
		Final = append(Final, fin)

	}

	// for _, val := range Final {

	// 	fmt.Println(val)

	// 	fmt.Println()
	// }

	models.GetPersonalize(&personalize_data, userid, TenantId)

	var first = user.FirstName
	var last = user.LastName

	if first != "" {
		firstn = strings.ToUpper(first[:1])
	}
	if user.LastName != "" {
		lastn = strings.ToUpper(last[:1])
	}

	var nameLength = len(user.FirstName) + len(user.LastName)

	user.NameLength = nameLength

	if nameLength > 18 {
		if len(first) > 6 {
			user.LimitedLengthName = first[:5] + "..."

		} else {
			user.LimitedLengthName = first[:len(first)] + "..."
		}

	}

	var Name = firstn + lastn

	user.NameString = Name

	if user.ProfileImagePath != "" {
		if user.StorageType == "local" {
			user.ProfileImagePath = "/" + user.ProfileImagePath
		} else if user.StorageType == "aws" {
			user.ProfileImagePath = "/image-resize?name=" + user.ProfileImagePath
		}
	}

	if interfaceValue, exists := c.Get("currentLanguage"); exists {
		if structValue, ok := interfaceValue.(models.TblLanguage); ok {
			currentLanguage = structValue
		}
	}

	session1, _ := c.Cookie("mainmenu")

	if session1 != "" {
		expandflg = true
	} else {
		expandflg = false
	}

	tblGeneralSettings, _ := models.GetGeneralSettings(TenantId)

	if !strings.Contains(tblGeneralSettings.LogoPath, "/public/img") {

		if tblGeneralSettings.StorageType == "aws" {

			LogoPath = "/image-resize?name=" + tblGeneralSettings.LogoPath
			ExpantLogoPath = "/image-resize?name=" + tblGeneralSettings.ExpandLogoPath

		} else if tblGeneralSettings.StorageType == "local" {

			LogoPath = "Storage/" + tblGeneralSettings.LogoPath
			ExpantLogoPath = "Storage/" + tblGeneralSettings.ExpandLogoPath

		}

	} else {

		LogoPath = tblGeneralSettings.LogoPath

	}

	Mod := USR{
		TblModule:         Final,
		Name:              user.FirstName + " " + user.LastName,
		NameLength:        user.NameLength,
		LimitedLengthName: user.LimitedLengthName,
		RoleName:          user.RoleName,
		ProfileImagePath:  user.ProfileImagePath,
		Languages:         GetLanguageList(),
		// DefaultLanguage:  GetDefaultLanguageDetails(userid),
		Personalize:     personalize_data,
		CurrentLanguage: currentLanguage,
		Footer:          os.Getenv("COPYRIGHTS"),
		NameString:      user.NameString,
		Expand:          expandflg,
		LogoPath:        LogoPath,
		ExpandLogoPath:  ExpantLogoPath,
		TblPageTypes:    pagetype,
	}

	return Mod
}

func GetPermissionIDS(roleid int) ([]TblRolePermission, error) {

	var perm []TblRolePermission

	err := GetPermissionId(&perm, roleid)

	if err != nil {

		log.Println(err)

		return []TblRolePermission{}, err

	}

	return perm, nil
}

var db = config.SetupDB()

/**/
func GetAllParentModule(modules *[]TblModule) error {

	if err := db.Table("tbl_modules").Where("default_module = 0 and parent_id = 0").Order("tbl_modules.order_index asc").Find(&modules).
		Error; err != nil {

		return err
	}

	return nil
}

func GetAllSubModules(modules *[]TblModule, parentId []int) error {

	if err := db.Table("tbl_modules").Where("parent_id in (?)", parentId).Order("tbl_modules.order_index asc").Find(&modules).
		Error; err != nil {

		return err
	}

	return nil
}

func GetAllSubModule(modules *[]TblModule, moduleid int) error {

	if err := db.Table("tbl_modules").Where("parent_id = (?) and menu_type='tab'", moduleid).Order("tbl_modules.order_index asc").Find(&modules).
		Error; err != nil {

		return err
	}

	return nil
}

func GetModulePermissions(permission *[]TblModulePermission, modid int, ids []int) error {

	query := db.Table("tbl_module_permissions").Select("tbl_module_permissions.*,tbl_modules.module_name").Joins("inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id").Order("tbl_modules.order_index asc,tbl_module_permissions.order_index asc")

	if len(ids) > 0 {

		query = query.Where("tbl_module_permissions.id in (?)", ids)

	}

	if modid != 0 && modid > -1 {

		query = query.Where("module_id = (?)", modid)
	}

	query.Find(&permission)

	if err := query.Error; err != nil {

		return err
	}

	return nil
}

/*Get PermissionId By RoleId*/
func GetPermissionId(perm *[]TblRolePermission, roleid int) error {

	if err := db.Model(TblRolePermission{}).Where("role_id=?", roleid).Find(&perm).Error; err != nil {

		return err
	}

	return nil
}

/**/
func GetAllParentModules1(mod *[]TblModule) (err error) {

	if err := DB.Model(TblModule{}).Where("parent_id=0 and (tenant_id is NULL or tenant_id = ?)", TenantId).Order("order_index asc").Find(&mod).Error; err != nil {

		return err
	}

	return nil
}

/**/
func GetDefaultPagetypes(mod *[]TablePageType) (err error) {

	if err := DB.Table("tbl_page_types").Where("is_deleted=0").Find(&mod).Error; err != nil {

		return err
	}

	return nil
}

/**/
func GetAllSubModules1(mod *[]TblModule, ids []int) (err error) {

	if err := DB.Debug().Model(TblModule{}).Where("(tbl_modules.parent_id in (?) or id in(?)) and tbl_modules.assign_permission=0 and (tenant_id is NULL or tenant_id = ?)", ids, ids, TenantId).Order("order_index").Preload("TblModulePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("assign_permission =1 and tenant_id is NULL or tenant_id = ?", TenantId).Order("order_index asc")
	}).Find(&mod).Error; err != nil {

		return err
	}

	return nil
}

func Permissions() []TblModule {

	var (
		allmodule  []TblModule
		allmodules []TblModule
		parentid   []int //all parentid
		submod     []TblModule
	)

	GetAllParentModules1(&allmodule)
	for _, val := range allmodule {
		parentid = append(parentid, val.Id)
	}

	GetAllSubModules1(&submod, parentid)
	for _, val := range allmodule {

		if val.GroupFlg == 1 {

			var newmod TblModule
			newmod.Id = val.Id
			newmod.Description = val.Description
			newmod.CreatedBy = val.CreatedBy
			newmod.ModuleName = val.ModuleName
			newmod.IsActive = val.IsActive
			newmod.IconPath = val.IconPath
			newmod.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

			for _, sub := range submod {
				if sub.ParentId == val.Id {
					for _, getmod := range sub.TblModulePermission {
						if getmod.ModuleId == sub.Id {

							var modper TblModulePermission
							modper.Id = getmod.Id
							modper.Description = getmod.Description
							modper.DisplayName = getmod.DisplayName
							modper.ModuleName = getmod.ModuleName
							modper.RouteName = getmod.RouteName
							modper.CreatedBy = getmod.CreatedBy
							modper.Description = getmod.Description
							modper.TblRolePermission = getmod.TblRolePermission
							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")
							modper.FullAccessPermission = getmod.FullAccessPermission
							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}
				}

			}

			allmodules = append(allmodules, newmod)

		} else {

			for _, sub := range submod {

				if sub.ParentId == val.Id {

					var newmod TblModule
					newmod.Id = sub.Id
					newmod.Description = sub.Description
					newmod.CreatedBy = sub.CreatedBy
					newmod.ModuleName = sub.ModuleName
					newmod.IsActive = sub.IsActive
					newmod.IconPath = sub.IconPath
					newmod.CreatedDate = sub.CreatedOn.Format("02 Jan 2006 03:04 PM")

					for _, getmod := range sub.TblModulePermission {
						if getmod.ModuleId == sub.Id {

							var modper TblModulePermission
							modper.Id = getmod.Id
							modper.Description = sub.Description
							modper.DisplayName = getmod.DisplayName
							modper.ModuleName = getmod.ModuleName
							modper.RouteName = getmod.RouteName
							modper.CreatedBy = getmod.CreatedBy
							modper.Description = getmod.Description
							modper.TblRolePermission = getmod.TblRolePermission
							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")
							modper.FullAccessPermission = getmod.FullAccessPermission
							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}

					allmodules = append(allmodules, newmod)

				}

			}
		}

	}

	return allmodules
}
