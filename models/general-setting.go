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
}

func GetGeneralSettings() (gensetting TblGeneralSetting, err error) {

	if gerr := db.Table("tbl_general_settings").Where("id=1").First(&gensetting).Error; gerr != nil {

		return gensetting, gerr
	}

	return gensetting, nil
}

func UpdateGeneralSettings(gensetting TblGeneralSetting) error {

	if err := db.Table("tbl_general_settings").Where("id=1").UpdateColumns(map[string]interface{}{"company_name": gensetting.CompanyName, "logo_path": gensetting.LogoPath, "expand_logo_path": gensetting.ExpandLogoPath, "date_format": gensetting.DateFormat, "time_format": gensetting.TimeFormat, "time_zone": gensetting.TimeZone, "language_id": gensetting.LanguageId, "modified_by": gensetting.ModifiedBy, "modified_on": gensetting.ModifiedOn}).Error; err != nil {

		return err

	}

	return nil

}

func UpdateLanuguageInGeneral(languageid int) error {

	if err := db.Table("tbl_general_settings").Where("id=1").UpdateColumns(map[string]interface{}{"language_id": languageid}).Error; err != nil {

		return err

	}

	return nil

}

func GetTimeZones() (timzones []string, err error) {

	result := db.Table("tbl_timezones").Select("timezone").Find(&timzones)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return timzones, nil
}
