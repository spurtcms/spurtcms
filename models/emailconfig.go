package models

import "gorm.io/datatypes"

type TblEmailConfigurations struct {
	Id           int
	StmpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
}

type StmpConfig struct {
	Mail     string
	Password string
	Host     string
	Port     string
}

func GetMail(mail *TblEmailConfigurations) error {

	if err := db.Table("tbl_email_configurations").First(&mail).Error; err != nil {

		return err
	}

	return nil
}

func UpdateMail(mail *TblEmailConfigurations) (upadtemail *TblEmailConfigurations, err error) {

	if err := db.Debug().Table("tbl_email_configurations").Where("id =?", mail.Id).UpdateColumns(map[string]interface{}{"selected_type": mail.SelectedType, "stmp_config": mail.StmpConfig}).Error; err != nil {
		return &TblEmailConfigurations{}, err
	}
	return mail, nil
}
