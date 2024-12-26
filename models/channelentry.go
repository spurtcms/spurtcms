package models

import (
	"time"

	chn "github.com/spurtcms/channels"
	"gorm.io/gorm"
)

var IST, _ = time.LoadLocation("Asia/Kolkata")

type TblCategory struct {
	Id                 int
	CategoryName       string
	CategorySlug       string
	Description        string
	ImagePath          string
	CreatedOn          time.Time
	CreatedBy          int
	ModifiedOn         time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted          int
	DeletedOn          time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy          int       `gorm:"DEFAULT:NULL"`
	ParentId           int
	CreatedDate        string   `gorm:"-"`
	ModifiedDate       string   `gorm:"-"`
	DateString         string   `gorm:"-"`
	ParentCategoryName string   `gorm:"-"`
	Parent             []string `gorm:"-"`
	ParentWithChild    []Result `gorm:"-"`
}

type Result struct {
	CategoryName string
}

type TblChannelEntries struct {
	Id                   int
	Title                string `form:"title" binding:"required"`
	Slug                 string `form:"slug" binding:"required"`
	Description          string
	UserId               int
	ChannelId            int
	Status               int //0-draft 1-publish 2-unpublish
	IsActive             int
	IsDeleted            int       `gorm:"DEFAULT:0"`
	DeletedBy            int       `gorm:"DEFAULT:NULL"`
	DeletedOn            time.Time `gorm:"DEFAULT:NULL"`
	CreatedOn            time.Time
	CreatedBy            int
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"DEFAULT:NULL"`
	CoverImage           string
	ThumbnailImage       string
	MetaTitle            string `form:"metatitle" binding:"required"`
	MetaDescription      string `form:"metadesc" binding:"required"`
	Keyword              string `form:"keywords" binding:"required"`
	CategoriesId         string
	RelatedArticles      string
	CreatedDate          string                 `gorm:"-"`
	ModifiedDate         string                 `gorm:"-"`
	Username             string                 `gorm:"<-:false"`
	TblChannelEntryField []TblChannelEntryField `gorm:"<-:false; foreignKey:ChannelEntryId"`
	Category             []TblCategory          `gorm:"<-:false; foreignKey:Id"`
	CategoryGroup        string                 `gorm:"<-:false"`
	ChannelName          string                 `gorm:"<-:false"`
	Cno                  int                    `gorm:"<-:false"`
}

type TblChannelEntryField struct {
	Id             int
	FieldName      string
	FieldValue     string
	ChannelEntryId int
	FieldId        int
	CreatedOn      time.Time
	CreatedBy      int
	ModifiedOn     time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	FieldTypeId    int       `gorm:"<-:false"`
}

type ExportData struct {
	Title       string
	CreatedBy   string
	CreatedDate string
	Status      string
}

type TblModule struct {
	Id                   int
	ModuleName           string
	IsActive             int
	CreatedBy            int
	CreatedOn            time.Time
	CreatedDate          string `gorm:"<-:false"`
	DefaultModule        int
	ParentId             int
	IconPath             string
	OrderIndex           int
	Description          string
	DisplayName          string `gorm:"column:display_name;<-:false"`
	RouteName            string `gorm:"column:route_name;<-:false"`
	RouteParentId        int    `gorm:"column:parent_id;<-:false"`
	FullAccessPermission int
	GroupFlg             int
}

func Exportentriesdata(exportdata *[]chn.Tblchannelentries, id []int, tenantid int) error {

	if err := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("inner join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted=0 and tbl_channel_entries.id IN ? and (tbl_channel_entries.tenant_id is NULL or tbl_channel_entries.tenant_id = ?)", id, tenantid).Order("tbl_channel_entries.id desc").Preload("TblChannelEntryField", func(db *gorm.DB) *gorm.DB {
		return db.Select("tbl_channel_entry_fields.*,tbl_fields.field_type_id").Joins("inner join tbl_fields on tbl_fields.Id = tbl_channel_entry_fields.field_id")

	}).Find(&exportdata).Error; err != nil {
		return err
	}
	return nil

}

func Entryid(entry string, tenantid int) (Id int, err error) {

	var modules TblModule

	if err := DB.Table("tbl_modules").Where("module_name=? and parent_id!=?", entry, 0).First(&modules).Error; err != nil {

		return 0, err
	}

	id := modules.Id

	return id, nil
}
