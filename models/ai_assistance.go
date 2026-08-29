package models

import (
	"database/sql"
	"time"
)

type TblApps struct {
	Id            int                    `gorm:"primaryKey;auto_increment;type:serial"`
	Title         string                 `gorm:"type:character varying"`
	Description   string                 `gorm:"type:character varying"`
	FieldJsonPath string                 `gorm:"type:character varying"`
	IconName      string                 `gorm:"type:character varying"`
	IconPath      string                 `gorm:"type:character varying"`
	CreatedBy     int                    `gorm:"type:integer"`
	CreatedOn     time.Time              `gorm:"type:timestamp without time zone"`
	ModifiedBy    int                    `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn    time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive      int                    `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted     int                    `gorm:"type:integer;DEFAULT:0"`
	DeletedBy     int                    `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn     time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId      string                    `gorm:"type:character varying;DEFAULT:NULL"`
	CustomFields  map[string]interface{} `gorm:"-"`
	ChannelSlug   string                 `gorm:"type:character varying"`
}
type TblAiPrompt struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	AppId        int       `gorm:"type:integer;DEFAULT:NULL"`
	MasterId     int       `gorm:"type:integer;DEFAULT:NULL"`
	ChildId      int       `gorm:"type:integer;DEFAULT:NULL"`
	PromptLevel  int       `gorm:"type:integer;DEFAULT:NULL"`
	LanguageId   int       `gorm:"type:integer;DEFAULT:NULL"`
	SystemPrompt string    `gorm:"type:character varying"`
	UserPrompt   string    `gorm:"type:character varying"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive     int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     string    `gorm:"type:character varying;DEFAULT:NULL"`
}

func GetAppList(keyword string, tenant string) (tblapp []TblApps, count int64, err error) {

	query := DB.Debug().Model(TblApps{}).Where("tbl_apps.is_deleted = 0 and tbl_apps.is_active = 1")

	if keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_apps.title)) ILIKE LOWER(TRIM(?)))", "%"+keyword+"%")

	}

	listQuery := query.Find(&tblapp)

	if listQuery.Error != nil {

		return tblapp, -1, listQuery.Error
	}

	countQuery := query.Count(&count)
	if countQuery.Error != nil {

		return tblapp, -1, countQuery.Error
	}

	return tblapp, count, nil
}
func GetAiPrompts(aiprompt *TblAiPrompt, id int) error {

	if err := DB.Table("tbl_ai_prompts").Select("tbl_ai_prompts.*").Where("is_deleted=0 and id = ?", id).First(&aiprompt).Error; err != nil {

		return err
	}

	return nil
}

func ArticleCountUpdate(userid int, tenantid string) error {
	var currentCount sql.NullInt64

	if err := DB.Debug().Table("tbl_users").Select("article_count").Where("id = ? AND tenant_id =?", userid, tenantid).Scan(&currentCount).Error; err != nil {
		return err
	}

	var newCount int64
	if !currentCount.Valid {
		newCount = 1
	} else {
		newCount = currentCount.Int64 + 1
	}

	if err := DB.Debug().Table("tbl_users").Where("id = ? AND tenant_id=?", userid, tenantid).UpdateColumns(map[string]interface{}{"article_count": newCount}).Error; err != nil {
		return err
	}

	return nil
}

//Get ArticleCount//

func GetArticleCount(userid int, tenantid string) (TblUser, error) {

	var userDetails TblUser

	if err := DB.Table("tbl_users").Where("id=? and tenant_id=?", userid, tenantid).First(&userDetails).Error; err != nil {

		return TblUser{}, err
	}

	return userDetails, nil
}
