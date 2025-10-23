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
	TenantId    string
}

type TblMstrTenant struct {
	Id        int `gorm:"primaryKey;auto_increment"`
	TenantId  string
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
	TenantId     string
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

func (model ModelConfig) GetTenantDetails(tenantId string, tenantDetails *TblMstrTenant) error {

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

func (model ModelConfig) GettenantByapikey(apikey string) (tenant TblGraphqlSettings, err error) {

	if err := model.DB.Debug().Model(TblGraphqlSettings{}).Where("is_deleted = 0 and token = ?", apikey).First(&tenant).Error; err != nil {

		return TblGraphqlSettings{}, err
	}

	return tenant, nil
}

func (model ModelConfig) GetUserbyTenantId(tenantid string) (user models.TblUser, err error) {

	if err := model.DB.Debug().Table("tbl_users").Where("is_deleted = 0 and tenant_id = ?", tenantid).First(&user).Error; err != nil {

		return models.TblUser{}, err
	}

	return user, nil
}
