package models

import "time"

type TblGeneralSetting struct {
	Id             int
	CompanyName    string
	LogoPath       string
	ExpandLogoPath string
	DateFormat     string
	TimeFormat     string
	TimeZone       string
	LanguageId     int
	ModifiedBy     int
	ModifiedOn     time.Time
	StorageType    string
}

type TblTimeZone struct {
	Id       int
	Timezone string
	TenantId int
}

func GetGeneralSettings(tenantid int) (gensetting TblGeneralSetting, err error) {

	if gerr := DB.Debug().Table("tbl_general_settings").Where("tenant_id = ?", tenantid).First(&gensetting).Error; gerr != nil {

		return gensetting, gerr
	}

	return gensetting, nil
}

func UpdateGeneralSettings(gensetting TblGeneralSetting, tenantid int) error {

	if err := DB.Debug().Table("tbl_general_settings").Where("tenant_id = ?", tenantid).UpdateColumns(map[string]interface{}{"company_name": gensetting.CompanyName, "logo_path": gensetting.LogoPath, "expand_logo_path": gensetting.ExpandLogoPath, "date_format": gensetting.DateFormat, "time_format": gensetting.TimeFormat, "time_zone": gensetting.TimeZone, "language_id": gensetting.LanguageId, "modified_by": gensetting.ModifiedBy, "modified_on": gensetting.ModifiedOn, "storage_type": gensetting.StorageType}).Error; err != nil {

		return err

	}

	return nil

}

func UpdateLanuguageInGeneral(languageid int, tenantid int) error {

	if err := DB.Table("tbl_general_settings").Where(" tenant_id = ?", tenantid).UpdateColumns(map[string]interface{}{"language_id": languageid}).Error; err != nil {

		return err

	}

	return nil

}

func GetTimeZones(tenantid int) (timzones []TblTimeZone, err error) {

	result := DB.Table("tbl_timezones").Where("tenant_id is NULL or tenant_id = ?", tenantid).Find(&timzones)

	if result.Error != nil {
		return nil, result.Error
	}

	return timzones, nil
}

func LastLoginActivity(userid, tenantid int) error {

	log_time := time.Now().UTC()

	if err := DB.Table("tbl_users").Where("id=? and tenant_id=?", userid, tenantid).Update("last_login", log_time).Error; err != nil {

		return err
	}
	return nil
}
