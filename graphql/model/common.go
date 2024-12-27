package model

import (
	"spurt-cms/models"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	Model ModelConfig
)

type ModelConfig struct {
	DB *gorm.DB
}

type TblGraphqlSettings struct {
	Id          int `gorm:"primaryKey;auto_increment;"`
	TokenName   string
	Description string
	Duration    string
	CreatedBy   int
	CreatedOn   time.Time
	ModifiedBy  int
	ModifiedOn  time.Time
	DeletedBy   int
	DeletedOn   time.Time
	IsDeleted   int `gorm:"DEFAULT:0"`
	Token       string
	ExpiryTime  time.Time
	IsDefault   int `gorm:"Default:NULl"`
	TenantId    int
}

type TblMstrTenant struct {
	Id        int `gorm:"primaryKey;auto_increment"`
	TenantId  int
	IsDeleted int `gorm:"DEFAULT:0"`
	DeletedBy int `gorm:"DEFAULT:NULL"`
	DeletedOn time.Time
}

type TblStorageTypes struct {
	Id           int `gorm:"primaryKey;auto_increment"`
	Local        string
	Aws          datatypes.JSONMap
	Azure        datatypes.JSONMap
	Drive        datatypes.JSONMap
	SelectedType string 
	TenantId     int   
}

func init() {
	Model = ModelConfig{DB: models.DB}
}

func (model ModelConfig) GetApiSettings(apikey string, graphqlsettings *TblGraphqlSettings) error {

	if err := model.DB.Debug().Model(TblGraphqlSettings{}).Where("is_deleted = 0 and token = ?", apikey).First(&graphqlsettings).Error; err != nil {

		return err
	}

	return nil
}

func (model ModelConfig) GetTenantDetails(tenantId int, tenantDetails *TblMstrTenant) error {

	if err := model.DB.Debug().Model(TblMstrTenant{}).Where("is_deleted=0 and id=?", tenantId).First(&tenantDetails).Error; err != nil {

		return err
	}

	return nil
}

func GetStorageType(db *gorm.DB, storageType *TblStorageTypes) error {

	if err := db.Debug().Model(TblStorageTypes{}).First(&storageType).Error; err != nil {

		return err
	}

	return nil
}
