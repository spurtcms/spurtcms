package models

import (
	"time"
)

type TblBlock struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	BlockDescription string    `gorm:"type:text"`
	BlockContent     string    `gorm:"type:text"`
	BlockCss         string    `gorm:"type:text"`
	CoverImage       string    `gorm:"type:character varying"`
	IconImage        string    `gorm:"type:character varying"`
	TenantId         int       `gorm:"type:integer"`
	Prime            int       `gorm:"type:integer"`
	IsActive         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedBy       int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	ProfileImagePath string    `gorm:"<-:false"`
	FirstName        string    `gorm:"<-:false"`
	LastName         string    `gorm:"<-:false"`
	Username         string    `gorm:"<-:false"`
	TagValueArr      []string  `gorm:"-"`
	TagValue         string    `gorm:"<-:false;"`
	NameString       string    `gorm:"-"`
	Actions          string    `gorm:"<-:false"`
	CreatedDate      string    `gorm:"-:migration;<-:false"`
	ModifiedDate     string    `gorm:"-:migration;<-:false"`
}
type TblTemplates struct {
	Id               int
	TemplateName     string
	TemplateSlug     string
	TemplateModuleId int
	ImageName        string
	ImagePath        string
	DeployLink       string
	GithubLink       string
	CreatedOn        time.Time
	CreatedBy        int
	IsActive         int       `gorm:"DEFAULT:1"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	TenantId         int
	ChannelId        int
}

type TblTemplateModules struct {
	Id                 int
	TemplateModuleName string
	TemplateModuleSlug string
	Description        string
	IsActive           int
	CreatedOn          time.Time
	CreatedBy          int
	TenantId           int
	Templates          []TblTemplates `gorm:"foreignKey:TemplateModuleId"`
	TemplateCount      int            `gorm:"-"`
	ChannelId          int
}

type TblChannel struct {
	Id                 int
	ChannelName        string
	ChannelDescription string
	SlugName           string
	FieldGroupId       int
	IsActive           int
	IsDeleted          int
	CreatedBy          int
	ModifiedOn         time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	DateString         string
	EntriesCount       int
	ProfileImagePath   string `gorm:"-;<-:false"`
	ChannelType        string
	TenantId           int
	TemplatesCount     int `gorm:"-"`
}

func GetTemplateModuleList(keyword string, tenantId int, channelId string) (tempModuleList []TblTemplateModules, count int64, err error) {

	query := DB.Debug().Preload("Templates", "is_active = 1 and is_deleted = 0 and lower(trim(template_name)) like lower(trim(?)) ", "%"+keyword+"%")

	if channelId != "" {
		query = query.Where("is_active = 1 and tenant_id = ? and channel_id = ?", tenantId, channelId)

	} else {
		query = query.Where("is_active = 1 and tenant_id = ?", tenantId)
	}

	listQuery := query.Find(&tempModuleList)
	if listQuery.Error != nil {

		return []TblTemplateModules{}, -1, listQuery.Error
	}

	countQuery := query.Count(&count)
	if countQuery.Error != nil {

		return []TblTemplateModules{}, -1, countQuery.Error
	}

	return tempModuleList, count, nil
}

type TemplatesCount struct {
	ChannelId     int
	TemplateCount int
}

func GetChannelBasedTemplateCount(channelId string, tenantId int) (templatesCount TemplatesCount, err error) {

	var count TemplatesCount

	query := DB.Debug().Table("tbl_templates").Select("count(id) as template_count, channel_id").Where("is_deleted = 0 AND is_active = 1 AND tenant_id = ? and channel_id = ?", tenantId, channelId)

	err = query.Group("channel_id").Find(&count).Error
	if err != nil {
		return TemplatesCount{}, err
	}

	return count, nil
}

func GetChannelDetailsWithTemplateCount(channelId string, tenantId int) (channelDetails TblChannel, err error) {

	query := DB.Debug().Where("is_deleted = 0 and is_active = 1 and tenant_id = ? and id = ? ", tenantId, channelId).First(&channelDetails)
	if query.Error != nil {
		return TblChannel{}, query.Error
	}

	templateCount, err := GetChannelBasedTemplateCount(channelId, tenantId)
	if query.Error != nil {
		return TblChannel{}, query.Error
	}

	channelDetails.EntriesCount = templateCount.TemplateCount

	return channelDetails, err

}
func BlockLists(limit, offset int, filter Filter, tenantid int) (block []TblBlock, Totalblock int64, err error) {

	query := DB.Select("tbl_blocks.*,max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username, STRING_AGG(tbl_block_tags.tag_name, ', ') as tag_value ,(case when (select id from tbl_block_collections where tbl_block_collections.block_id = tbl_blocks.id and is_deleted = 0 limit 1) is not null then 'true' else 'false' end ) as actions ").Table("tbl_blocks").Joins("inner join tbl_block_tags on tbl_block_tags.block_id = tbl_blocks.id").Joins("left join tbl_block_collections on tbl_block_collections.block_id = tbl_blocks.id").Joins("inner join tbl_users on case when tbl_block_collections.id is not null then tbl_users.id = tbl_block_collections.user_id else tbl_users.id = tbl_blocks.created_by end").Where("tbl_blocks.is_deleted = ?  and tbl_blocks.tenant_id =?", 0, tenantid).Group("tbl_blocks.id").Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_blocks.title)) LIKE LOWER(TRIM(?))  ", "%"+filter.Keyword+"%")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&block)

		return block, 0, err

	}

	query.Find(&block).Count(&Totalblock)

	return block, Totalblock, err

}
