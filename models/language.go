package models

import (
	"fmt"
	"time"
)

type TblLanguage struct {
	Id                int
	LanguageName      string
	LanguageCode      string
	CreatedOn         time.Time
	CreatedBy         int
	ModifiedOn        time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL"`
	DeletedOn         time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted         int       `gorm:"DEFAULT:0"`
	ImagePath         string
	IsStatus          int
	IsDefault         int
	JsonPath          string
	DefaultLanguageId int    `gorm:"<-:false"`
	DateString        string `gorm:"-"`
	TenantId          int
	StorageType       string
}

type TblUser struct {
	Id                   int       `gorm:"column:id"`
	Uuid                 string    `gorm:"column:uuid"`
	FirstName            string    `gorm:"column:first_name"`
	LastName             string    `gorm:"column:last_name"`
	RoleId               int       `gorm:"column:role_id"`
	Email                string    `gorm:"column:email"`
	Username             string    `gorm:"column:username"`
	Password             string    `gorm:"column:password"`
	MobileNo             string    `gorm:"column:mobile_no"`
	IsActive             int       `gorm:"column:is_active"`
	ProfileImage         string    `gorm:"column:profile_image"`
	ProfileImagePath     string    `gorm:"column:profile_image_path"`
	StorageType          string    `gorm:"column:storage_type"`
	DataAccess           int       `gorm:"column:data_access"`
	CreatedOn            time.Time `gorm:"column:created_on"`
	CreatedBy            int       `gorm:"column:created_by"`
	ModifiedOn           time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"column:modified_by;DEFAULT:NULL"`
	LastLogin            time.Time `gorm:"column:last_login;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"column:is_deleted"`
	DeletedOn            time.Time `gorm:"column:deleted_on;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"column:deleted_by;DEFAULT:NULL"`
	ModuleName           string    `gorm:"-"`
	RouteName            string    `gorm:"-:migration;<-:false"`
	DisplayName          string    `gorm:"-:migration;<-:false"`
	Description          string    `gorm:"-"`
	ModuleId             int       `gorm:"-:migration;<-:false"`
	PermissionId         int       `gorm:"-"`
	FullAccessPermission int       `gorm:"-:migration;<-:false"`
	RoleName             string    `gorm:"-:migration;<-:false"`
	DefaultLanguageId    int       `gorm:"column:default_language_id"`
	NameString           string    `gorm:"-"`
	TenantId             string
	Otp                  int       `gorm:"column:otp"`
	OtpExpiry            time.Time `gorm:"column:otp_expiry;DEFAULT:NULL"`
	NameLength           int       `gorm:"-:migration;<-:false"`
	LimitedLengthName    string    `gorm:"-:migration;<-:false"`
	S3FolderName         string    `gorm:"column:s3_folder_name"`
	Subdomain            string
	GoTemplateDefault    int `gorm:"column:go_template_default"`
	ArticleCount         int `gorm:"type:integer"`
}

type QUERY struct {
	UserId     int
	DataAccess int
}

func UpdateSubDomain(id int, subdomain string) error {

	result := DB.Table("tbl_users").Where("id = ?", id).Update("subdomain", subdomain)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetuserBySubdomain(subdomain string) (userdetils TblUser, err error) {

	if err := DB.Table("tbl_users").Where("is_deleted = 0 and subdomain = ? ", subdomain).First(&userdetils).Error; err != nil {

		return TblUser{}, err
	}

	return userdetils, nil
}
func (q QUERY) GetLanguageList(language *[]TblLanguage, limit, offset int, filter Filter, flag bool, tenantid string) ([]TblLanguage, int64) {

	var Total_languages int64

	var emptyslice []TblLanguage

	query := DB.Table("tbl_languages").Where("tbl_languages.is_deleted=? and tbl_languages.is_status=? ", 0, 1).Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_languages.language_name)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_languages.language_code)) LIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
		// .Or("LOWER(TRIM(tbl_languages.language_code)) ILIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%")

	}

	if flag {

		query.Find(&language)

		return *language, 0

	}

	if q.DataAccess == 1 {

		query = query.Where("tbl_languages.created_by = ?", q.UserId)
	}

	if limit != 0 && !flag {

		query.Offset(offset).Limit(limit).Find(&language)

		return *language, 0

	} else {

		query.Find(&language).Count(&Total_languages)

		return emptyslice, Total_languages
	}

}

func GetLanguageDetails(langData *TblLanguage, langid, userid int, tenantid string) error {

	var userDetails TblUser

	if err := DB.Table("tbl_users").Where("id= ? and (tenant_id is NULL or tenant_id = ?)", userid, tenantid).First(&userDetails).Error; err != nil {

		userDetails = TblUser{}
	}

	if userDetails.DefaultLanguageId == langid {

		if err := DB.Debug().Table("tbl_languages").Select("tbl_languages.*,tbl_users.default_language_id").Joins("left join tbl_users on  tbl_users.default_language_id = tbl_languages.Id").Where("tbl_languages.is_deleted=0 and tbl_users.is_deleted = 0 and tbl_languages.id = ? and tbl_users.id = ? and (tbl_languages.tenant_id is NULL or tbl_languages.tenant_id = ?)", langid, userid, tenantid).First(&langData).Error; err != nil {

			return err
		}

	} else {

		if err := DB.Debug().Table("tbl_languages").Where("is_deleted = 0 and id=?", langid).Where("(tenant_id is NULL or tenant_id =?)", tenantid).First(&langData).Error; err != nil {

			return err
		}

	}

	return nil
}

func AddNewLanguage(langData *TblLanguage, userid int) error {

	if langData.IsDefault == 1 {

		if err := DB.Debug().Table("tbl_languages").Create(&langData).Error; err != nil {

			return err
		}

		if err := DB.Debug().Table("tbl_users").Where("is_deleted=0 and id=?", userid).Update("default_language_id", langData.Id).Error; err != nil {

			return err
		}

	} else {

		if err := DB.Debug().Table("tbl_languages").Create(&langData).Error; err != nil {

			return err
		}

	}

	return nil
}

func UpdateLanguage(langData *TblLanguage, userid int, tenantid string) error {

	var default_lang TblLanguage

	if langData.IsDefault == 1 {

		if err := DB.Table("tbl_users").Where("is_deleted=0 and id=? and (tenant_id is NULL or tenant_id = ?)", userid, tenantid).Update("default_language_id", langData.Id).Error; err != nil {

			return err
		}

		if err := DB.Table("tbl_languages").Where("id=? and (tenant_id is NULL or tenant_id = ?)", langData.Id, tenantid).UpdateColumns(map[string]interface{}{"language_name": langData.LanguageName, "language_code": langData.LanguageCode, "is_status": langData.IsStatus, "image_path": langData.ImagePath, "modified_on": langData.ModifiedOn, "modified_by": langData.ModifiedBy, "storage_type": langData.StorageType}).Error; err != nil {

			return err
		}

	} else {

		err := GetDefaultLanguage(&default_lang, userid, tenantid)

		if err != nil {

			return err
		}

		if default_lang.Id != langData.Id {

			if err := DB.Table("tbl_languages").Where("id=? and (tenant_id is NULL or tenant_id = ?)", langData.Id, tenantid).UpdateColumns(map[string]interface{}{"language_name": langData.LanguageName, "language_code": langData.LanguageCode, "is_status": langData.IsStatus, "image_path": langData.ImagePath, "modified_on": langData.ModifiedOn, "modified_by": langData.ModifiedBy, "storage_type": langData.StorageType}).Error; err != nil {

				return err
			}

		} else {

			if err := DB.Table("tbl_languages").Where("id=? and (tenant_id is NULL or tenant_id = ?)", langData.Id, tenantid).UpdateColumns(map[string]interface{}{"language_name": langData.LanguageName, "language_code": langData.LanguageCode, "is_status": langData.IsStatus, "image_path": langData.ImagePath, "modified_on": langData.ModifiedOn, "modified_by": langData.ModifiedBy, "storage_type": langData.StorageType}).Error; err != nil {

				return err
			}

			MakeDefaultLanguage(userid, tenantid)

		}

	}

	return nil
}

func DeleteLanguage(langData *TblLanguage, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("id=? and (tenant_id is NULL or tenant_id = ?)", langData.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": langData.IsDeleted, "deleted_by": langData.DeletedBy, "deleted_on": langData.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

func FetchAllLanguage(languages *[]TblLanguage, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("is_deleted=0 and (tenant_id is NULL or tenant_id = ?)", tenantid).Order("id desc").Find(&languages).Error; err != nil {

		return err
	}

	return nil
}

func GetDefaultLanguage(default_lang *TblLanguage, userid int, tenantid string) error {

	if err := DB.Table("tbl_languages").Select("tbl_languages.*").Joins("inner join tbl_users on tbl_users.default_language_id = tbl_languages.id").Where("tbl_languages.is_deleted=0  and tbl_users.id=? and (tbl_languages.tenant_id is NULL or tbl_languages.tenant_id = ?)", userid, tenantid).First(&default_lang).Error; err != nil {

		return err
	}

	return nil
}

func GetLanguagebyLanguagecode(langData *TblLanguage, langCode string, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("is_deleted = 0 and language_code=? and (tenant_id is NULL or tenant_id = ?)", langCode, tenantid).First(&langData).Error; err != nil {

		return err
	}

	return nil
}

func MakeDefaultLanguage(userid int, tenantid string) error {

	var AutoDefaultLanguage TblLanguage

	if err := DB.Table("tbl_languages").Where("is_deleted = 0 and language_code='en' and (tenant_id is NULL or tenant_id = ?)", tenantid).First(&AutoDefaultLanguage).Error; err != nil {

		return err
	}

	// if err := db.Table("tbl_languages").Where("is_deleted = 0 and language_code='en'").Update("is_default", 1).Error; err != nil {

	// 	return err
	// }

	if err := DB.Table("tbl_users").Where("is_deleted=0 and id=? and (tenant_id is NULL or tenant_id = ?)", userid, tenantid).Update("default_language_id", AutoDefaultLanguage.Id).Error; err != nil {

		return err
	}

	return nil
}

func EliminateDefaultMany(tenantid string) error {

	if err := DB.Table("tbl_languages").Where("is_deleted=0 and is_default=1 and language_code!=? and (tenant_id is NULL or tenant_id = ?)", "en", tenantid).Update("is_default", 0).Error; err != nil {

		return err
	}

	return nil
}

func Checklang(language *TblLanguage, language_name string, langid int, tenantid string) error {

	if langid == 0 {
		if err := DB.Table("tbl_languages").Where("LOWER(TRIM(language_name))=LOWER(TRIM(?)) and is_deleted = 0 and (tenant_id is NULL or tenant_id = ?)", language_name, tenantid).First(&language).Error; err != nil {

			return err
		}
	} else {

		if err := DB.Table("tbl_languages").Where("LOWER(TRIM(language_name))=LOWER(TRIM(?)) and id not in(?) and is_deleted= 0 and (tenant_id is NULL or tenant_id = ?)", language_name, langid, tenantid).First(&language).Error; err != nil {

			return err
		}
	}
	return nil
}

func GetIdByLanguageName(langData *TblLanguage, LangName string, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("is_deleted = 0 and language_name=? and (tenant_id is NULL or tenant_id = ?)", LangName, tenantid).First(&langData).Error; err != nil {

		return err
	}

	return nil
}

func GetLanguageById(langData *TblLanguage, id int) error {

	if err := DB.Table("tbl_languages").Where("is_deleted =0 and id =?", id).First(&langData).Error; err != nil {

		return err
	}

	return nil
}

func Languageisactive(langData *TblLanguage, id, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("id=? and (tenant_id is NULL or tenant_id = ?)", id, tenantid).UpdateColumns(map[string]interface{}{"is_status": langData.IsStatus, "modified_by": langData.ModifiedBy, "modified_on": langData.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func RemoveLanguageImagePath(ImagePath string, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("image_path=? and (tenant_id is NULL or tenant_id = ?)", ImagePath, tenantid).UpdateColumns(map[string]interface{}{"image_path": ""}).Error; err != nil {

		return err
	}

	return nil
}
func MultiSelectDeleteLanguage(langData *TblLanguage, languageid []int, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("id in (?) and (tenant_id is NULL or tenant_id = ?)", languageid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": langData.IsDeleted, "deleted_by": langData.DeletedBy, "deleted_on": langData.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

func MultiSelectLanguageisactive(langData *TblLanguage, id []int, tenantid string) error {

	if err := DB.Table("tbl_languages").Where("id in (?) and (tenant_id is NULL or tenant_id = ?)", id, tenantid).UpdateColumns(map[string]interface{}{"is_status": langData.IsStatus, "modified_by": langData.ModifiedBy, "modified_on": langData.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func SetDefaultLang(langId, defaultVal int, tenantId string) (language TblLanguage, err error) {
	fmt.Println("Seumm")

	if defaultVal != 1 {

		if err := DB.Table("tbl_users").Where("is_deleted=0 and is_active = 1 and (tenant_id is NULL or tenant_id = ?)", tenantId).Update("default_language_id", langId).Error; err != nil {

			return TblLanguage{}, err
		}

		if err := DB.Table("tbl_languages").Where("is_deleted = 0 and id= ? ", langId).Update("is_default", 1).Error; err != nil {

			return TblLanguage{}, err
		}

		if err := DB.Table("tbl_languages").Where("is_deleted = 0 and id != ? ", langId).Update("is_default", 0).Error; err != nil {

			return TblLanguage{}, err
		}

		if err := DB.Table("tbl_languages").Where("is_deleted = 0 and id = ? ", langId).First(&language).Error; err != nil {

			return TblLanguage{}, err
		}

	}

	return language, nil
}
