package models

import (
	"time"
)

type TblMstrBlocks struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	SlugName         string    `gorm:"type:character varying"`
	ChannelSlugname  string    `gorm:"type:character varying"`
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
	ChannelID        string    `gorm:"type:character varying"`
	ChannelNames     []string  `gorm:"-"`
}

type TblMstrForms struct {
	Id                   int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormTitle            string    `gorm:"type:character varying"`
	FormSlug             string    `gorm:"type:character varying"`
	FormData             string    `gorm:"type:character varying"`
	Status               int       `gorm:"type:integer"`
	IsActive             int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn            time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"type:integer;DEFAULT:0"`
	TenantId             int       `gorm:"type:integer"`
	Uuid                 string    `gorm:"type:character varying"`
	Username             string    `gorm:"<-:false"`
	ProfileImagePath     string    `gorm:"<-:false"`
	NameString           string    `gorm:"<-:false"`
	FirstName            string    `gorm:"<-:false"`
	LastName             string    `gorm:"<-:false"`
	DateString           string    `gorm:"-"`
	CreatedDate          string    `gorm:"-:migration;<-:false"`
	ModifiedDate         string    `gorm:"-:migration;<-:false"`
	FormImagePath        string    `gorm:"type:character varying"`
	FormDescription      string    `gorm:"type:character varying"`
	ChannelId            string    `gorm:"type:character varying"`
	ChannelName          string    `gorm:"type:character varying"`
	Channelnamearr       []string  `gorm:"-"`
	FormPreviewImagepath string    `gorm:"type:character varying"`
	FormPreviewImagename string    `gorm:"type:character varying"`
}
type TblMstrchannel struct {
	Id                 int       `gorm:"column:id"`
	ChannelName        string    `gorm:"column:channel_name"`
	ChannelDescription string    `gorm:"column:channel_description"`
	SlugName           string    `gorm:"column:slug_name"`
	FieldGroupId       int       `gorm:"column:field_group_id"`
	IsActive           int       `gorm:"column:is_active"`
	IsDeleted          int       `gorm:"column:is_deleted"`
	CreatedOn          time.Time `gorm:"column:created_on"`
	CreatedBy          int       `gorm:"column:created_by"`
	ModifiedOn         time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"column:modified_by;DEFAULT:NULL"`
	DateString         string    `gorm:"-"`
	ProfileImagePath   string    `gorm:"<-:false"`
	AuthorDetails      TblUser   `gorm:"foreignKey:Id;references:CreatedBy"`
	ChannelType        string    `gorm:"column:channel_type"`
	TenantId           int       `gorm:"column:tenant_id"`
	Username           string    `gorm:"<-:false"`
	FirstName          string    `gorm:"<-:false"`
	LastName           string    `gorm:"<-:false"`
	NameString         string    `gorm:"<-:false"`
	ImagePath          string    `gorm:"column:image_path"`
}

type TblMstrIntegration struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	ClientId     string    `gorm:"type:character varying"`
	ClientSecret string    `gorm:"type:character varying"`
	CoverImage   string    `gorm:"type:character varying"`
	GatewayName  string    `gorm:"type:character varying"`
	GatewayDesc  string    `gorm:"type:character varying"`
	IsActive     int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	TenantId     string    `gorm:"type:character varying"`
}
