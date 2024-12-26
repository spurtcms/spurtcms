package models

import "gorm.io/datatypes"

type TblStorageType struct {
	Id           int
	Local        string
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string
}

type AWS struct {
	AccessId   string `json:"accessid"`
	AccessKey  string `json:"accesskey"`
	Region     string `json:"region"`
	BucketName string `json:"bucketname"`
}

type Azure struct {
	StorageAccount string `json:"storageaccount"`
	AccountKey     string `json:"accountkey"`
	ContainerName  string `json:"containername"`
}

func UpdateStorageType(stype TblStorageType, _ int) (TblStorageType, error) {

	query := DB.Debug().Model(TblStorageType{}).Where("id=1")

	if stype.SelectedType == "local" {

		query = query.Omit("aws,azure,drive")

	} else if stype.SelectedType == "aws" {

		query = query.Omit("local,azure,drive")

	} else if stype.SelectedType == "azure" {

		query = query.Omit("aws,local,drive")

	} else if stype.SelectedType == "drive" {

		query = query.Omit("aws,azure,local")

	}

	query = query.UpdateColumns(map[string]interface{}{"local": stype.Local, "aws": stype.Aws, "azure": stype.Azure, "drive": stype.Drive, "selected_type": stype.SelectedType})

	if err := query.Error; err != nil {

		return TblStorageType{}, err
	}

	return stype, nil
}

func GetStorageValue(_ int) (tblstorgetype TblStorageType, err error) {

	if err := DB.Model(TblStorageType{}).Where("id=1").First(&tblstorgetype).Error; err != nil {

		return TblStorageType{}, err
	}

	return tblstorgetype, nil
}
