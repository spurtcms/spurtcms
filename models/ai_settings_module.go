package models

import (
	"time"
)

type TblAiSettingsModule struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	AIModule    string    `gorm:"type:character varying"`
	ApiKey      string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	AiModel     string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	TenantId    string       `gorm:"type:integer"`
	DateString  string    `gorm:"-"`
}

func CreateAiModule(module TblAiSettingsModule) (bool, error) {

	if err := DB.Create(&module).Error; err != nil {
		return false, err
	}

	return true, nil

}

func ListAiModule(limit, offset int, filter Filter, tenantid string) (list []TblAiSettingsModule, count int64, err error) {

	var aimodule []TblAiSettingsModule

	query := DB.Table("tbl_ai_settings_modules").Where("is_deleted=0 and tenant_id=?", tenantid).Order("created_on desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_ai_settings_modules.ai_module)) like LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&aimodule)

		return aimodule, count, nil

	}

	query.Find(&aimodule).Count(&count)

	return aimodule, count, nil

}

func StatusChange(id, isactive, modifiedby int, tenantid string) (bool, error) {

	result := DB.Table("tbl_ai_settings_modules").
		Where("id = ? AND tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": isactive})

	if result.Error != nil {
		return false, result.Error
	}

	result = DB.Table("tbl_ai_settings_modules").
		Where("id != ? AND tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": 0})

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func UpdateAiModule(module TblAiSettingsModule) (bool, error) {

	result := DB.Table("tbl_ai_settings_modules").
		Where("id = ? AND tenant_id = ?", module.Id, module.TenantId).UpdateColumns(map[string]interface{}{"ai_module": module.AIModule, "api_key": module.ApiKey, "description": module.Description, "ai_model": module.AiModel, "modified_by": module.ModifiedBy, "modified_on": module.ModifiedOn})

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func DeleteModule(id, userid int, tenantid string) (bool, error) {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	result := DB.Table("tbl_ai_settings_modules").
		Where("id = ? AND tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": "1", "deleted_by": userid, "modified_on": deletedon})

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func MultiSelectDelete(moduleids []int, userid int, tenantid string) (bool, error) {

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err := DB.Table("tbl_ai_settings_modules").Where("id in (?) and tenant_id=?", moduleids, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": "1", "deleted_on": deletedon, "deleted_by": userid}).Error; err != nil {

		return false, err
	}

	return true, nil

}
