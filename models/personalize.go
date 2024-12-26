package models

import "time"

type TblUserPersonalize struct {
	Id                  int
	UserId              int
	MenuBackgroundColor string
	FontColor           string
	IconColor           string
	LogoPath            string
	ExpandLogoPath      string
	CreatedOn           time.Time
	ModifiedOn          time.Time `gorm:"default:null"`
	TenantId            int
}

func GetPersonalize(personalize *TblUserPersonalize, id, tenantid int) error {

	if err := DB.Table("tbl_user_personalizes").Select("tbl_user_personalizes.*").Where("user_id=? and tenant_id = ?", id, tenantid).First(&personalize).Error; err != nil {

		return err
	}

	return nil
}

func CreatePersonalize(personalize *TblUserPersonalize) error {

	if err := DB.Table("tbl_user_personalizes").Create(&personalize).Error; err != nil {

		return err
	}

	return nil

}

func UpdatePersonalize(personalize *TblUserPersonalize, id int, tenantid int) error {

	if err := DB.Table("tbl_user_personalizes").Where("user_id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"menu_background_color": personalize.MenuBackgroundColor, "logo_path": personalize.LogoPath, "expand_logo_path": personalize.ExpandLogoPath, "modified_on": personalize.ModifiedOn}).Error; err != nil {

		return err
	}
	return nil

}
