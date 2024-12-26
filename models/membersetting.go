package models

import (
	"time"
)

type TblMemberSetting struct {
	Id                int
	AllowRegistration int
	MemberLogin       string
	ModifiedOn        time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL"`
	NotificationUsers string
}

func GetMemberSettings(tenantid int) (membersetting *[]TblMemberSetting, error bool) {

	if err := DB.Where("tenant_id = ?", tenantid).Table("tbl_member_settings").Find(&membersetting).Error; err != nil {

		return &[]TblMemberSetting{}, false
	}

	return membersetting, true
}

func UpdateMemberSetting(membersetting *TblMemberSetting, tenantid int) error {

	if err := DB.Model(TblMemberSetting{}).Where("id=? and tenant_id = ?", membersetting.Id, tenantid).UpdateColumns(map[string]interface{}{"allow_registration": membersetting.AllowRegistration, "member_login": membersetting.MemberLogin, "notification_users": membersetting.NotificationUsers, "modified_on": membersetting.ModifiedOn, "modified_by": membersetting.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}
