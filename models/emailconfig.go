package models

import "gorm.io/datatypes"

type TblEmailConfigurations struct {
	Id           int
	SmtpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
}

type SmtpConfig struct {
	Mail     string
	Password string
	Host     string
	Port     string
}

func GetMail(mail *TblEmailConfigurations, tenantid string) error {

	if err := DB.Table("tbl_email_configurations").Where("tenant_id = ?", tenantid).First(&mail).Error; err != nil {

		return err
	}

	return nil
}

func UpdateMail(mail *TblEmailConfigurations, tenantid string) (upadtemail *TblEmailConfigurations, err error) {

	if err := DB.Debug().Table("tbl_email_configurations").Where("tenant_id = ?", tenantid).UpdateColumns(map[string]interface{}{"selected_type": mail.SelectedType, "smtp_config": mail.SmtpConfig}).Error; err != nil {
		return &TblEmailConfigurations{}, err
	}
	return mail, nil
}
