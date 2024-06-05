package postgres

import (
	"spurt-cms/controllers"
	"time"

	"gorm.io/datatypes"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:integer"`
}

type TblUser struct {
	Id                int       `gorm:"primaryKey;type:serial"`
	Uuid              string    `gorm:"type:character varying"`
	FirstName         string    `gorm:"type:character varying"`
	LastName          string    `gorm:"type:character varying"`
	RoleId            int       `gorm:"type:integer"`
	Email             string    `gorm:"type:character varying"`
	Username          string    `gorm:"type:character varying"`
	Password          string    `gorm:"type:character varying"`
	MobileNo          string    `gorm:"type:character varying"`
	IsActive          int       `gorm:"type:integer"`
	ProfileImage      string    `gorm:"type:character varying"`
	ProfileImagePath  string    `gorm:"type:character varying"`
	DataAccess        int       `gorm:"type:integer"`
	DefaultLanguageId int       `gorm:"type:integer"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:integer"`
	LastLogin         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:integer"`
	DeletedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL"`
}

type TblCategories struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	CategoryName string    `gorm:"type:character varying"`
	CategorySlug string    `gorm:"type:character varying"`
	Description  string    `gorm:"type:character varying"`
	ImagePath    string    `gorm:"type:character varying"`
	ParentId     int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL;type:integer"`
}

type TblChannel struct {
	Id                 int    `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelName        string `gorm:"type:character varying"`
	ChannelDescription string
	SlugName           string    `gorm:"type:character varying"`
	FieldGroupId       int       `gorm:"type:integer"`
	IsActive           int       `gorm:"type:integer"`
	IsDeleted          int       `gorm:"type:integer"`
	CreatedOn          time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy          int       `gorm:"type:integer"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
}

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid             string    `gorm:"type:character varying"`
	FirstName        string    `gorm:"type:character varying"`
	LastName         string    `gorm:"type:character varying"`
	Email            string    `gorm:"type:character varying"`
	MobileNo         string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	ProfileImage     string    `gorm:"type:character varying"`
	ProfileImagePath string    `gorm:"type:character varying"`
	LastLogin        int       `gorm:"type:integer"`
	MemberGroupId    int       `gorm:"type:integer"`
	Password         string    `gorm:"type:character varying"`
	Username         string    `gorm:"DEFAULT:NULL"`
	Otp              int       `gorm:"DEFAULT:NULL"`
	OtpExpiry        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
}

type TblAccessControl struct {
	Id                int       `gorm:"primaryKey;auto_increment;type:serial"`
	AccessControlName string    `gorm:"type:character varying"`
	AccessControlSlug string    `gorm:"type:character varying"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:integer"`
	DeletedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL"`
}

type TblAccessControlPages struct {
	Id                       int       `gorm:"primaryKey;auto_increment;type:serial"`
	AccessControlUserGroupId int       `gorm:"type:integer"`
	SpacesId                 int       `gorm:"type:integer"`
	PageGroupId              int       `gorm:"type:integer"`
	PageId                   int       `gorm:"type:integer"`
	CreatedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy                int       `gorm:"type:integer"`
	ModifiedOn               time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy               int       `gorm:"DEFAULT:NULL"`
	IsDeleted                int       `gorm:"type:integer"`
	DeletedOn                time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy                int       `gorm:"DEFAULT:NULL"`
	ChannelId                int       `gorm:"type:integer"`
	EntryId                  int       `gorm:"type:integer"`
}

type TblAccessControlUserGroup struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	AccessControlId int       `gorm:"type:integer"`
	MemberGroupId   int       `gorm:"type:integer"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy       int       `gorm:"type:integer"`
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted       int       `gorm:"type:integer"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL"`
}

type TblEmailTemplate struct {
	Id                  int       `gorm:"primaryKey;auto_increment;type:serial"`
	TemplateName        string    `gorm:"type:character varying"`
	TemplateSubject     string    `gorm:"type:character varying"`
	TemplateMessage     string    `gorm:"type:text"`
	TemplateDescription string    `gorm:"type:text"`
	ModuleId            int       `gorm:"type:integer"`
	CreatedOn           time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy           int       `gorm:"type:integer"`
	ModifiedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"DEFAULT:NULL"`
	DeletedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy           int       `gorm:"DEFAULT:NULL"`
	IsDeleted           int       `gorm:"DEFAULT:0"`
	IsActive            int       `gorm:"type:integer"`
	IsDefault           int       `gorm:"type:integer"`
}

type TblField struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName        string    `gorm:"type:character varying"`
	FieldDesc        string    `gorm:"type:character varying"`
	FieldTypeId      int       `gorm:"type:integer"`
	MandatoryField   int       `gorm:"type:integer"`
	OptionExist      int       `gorm:"type:integer"`
	InitialValue     string    `gorm:"type:character varying"`
	Placeholder      string    `gorm:"type:character varying"`
	OrderIndex       int       `gorm:"type:integer"`
	ImagePath        string    `gorm:"type:character varying"`
	DatetimeFormat   string    `gorm:"type:character varying"`
	TimeFormat       string    `gorm:"type:character varying"`
	Url              string    `gorm:"type:character varying"`
	SectionParentId  int       `gorm:"type:integer"`
	CharacterAllowed int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
}

type TblFieldGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	GroupName  string    `gorm:"type:character varying"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
}

type TblFieldOption struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	OptionName  string    `gorm:"type:character varying"`
	OptionValue string    `gorm:"type:character varying"`
	FieldId     int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
}

type TblFieldType struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	TypeName   string    `gorm:"type:character varying"`
	TypeSlug   string    `gorm:"type:character varying"`
	IsActive   int       `gorm:"type:integer"`
	IsDeleted  int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblLanguage struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	LanguageName string    `gorm:"type:character varying"`
	LanguageCode string    `gorm:"type:character varying"`
	ImagePath    string    `gorm:"type:character varying"`
	IsStatus     int       `gorm:"type:integer"`
	IsDefault    int       `gorm:"type:integer"`
	JsonPath     string    `gorm:"type:character varying"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"DEFAULT:0"`
}

type TblModule struct {
	Id                   int       `gorm:"primaryKey;type:serial"`
	ModuleName           string    `gorm:"type:character varying"`
	IsActive             int       `gorm:"type:integer"`
	DefaultModule        int       `gorm:"type:integer"`
	ParentId             int       `gorm:"type:integer"`
	IconPath             string    `gorm:"type:character varying"`
	AssignPermission     int       `gorm:"type:integer"`
	Description          string    `gorm:"type:character varying"`
	OrderIndex           int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone"`
	MenuType             string    `gorm:"type:character varying"`
	FullAccessPermission int       `gorm:"type:integer"`
	GroupFlg             int       `gorm:"type:integer"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey;type:serial"`
	RouteName            string    `gorm:"type:character varying;unique"`
	DisplayName          string    `gorm:"type:character varying"`
	SlugName             string    `gorm:"type:character varying"`
	Description          string    `gorm:"type:character varying"`
	ModuleId             int       `gorm:"type:integer"`
	FullAccessPermission int       `gorm:"type:integer"`
	ParentId             int       `gorm:"type:integer"`
	AssignPermission     int       `gorm:"type:integer"`
	BreadcrumbName       string    `gorm:"type:character varying"`
	OrderIndex           int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
}

type TblChannelCategorie struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId  int       `gorm:"type:integer"`
	CategoryId string    `gorm:"type:character varying"`
	CreatedAt  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
}

type TblGroupField struct {
	Id        int `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId int `gorm:"type:integer"`
	FieldId   int `gorm:"type:integer"`
}

type TblUserPersonalize struct {
	Id                  int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId              int       `gorm:"type:integer"`
	MenuBackgroundColor string    `gorm:"type:character varying"`
	FontColor           string    `gorm:"type:character varying"`
	IconColor           string    `gorm:"type:character varying"`
	LogoPath            string    `gorm:"type:character varying"`
	ExpandLogoPath      string    `gorm:"type:character varying"`
	CreatedOn           time.Time `gorm:"type:timestamp without time zone"`
	ModifiedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblRoleUser struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId     int       `gorm:"type:integer"`
	UserId     int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblChannelEntrie struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title           string    `gorm:"type:character varying"`
	Slug            string    `gorm:"type:character varying"`
	Description     string    `gorm:"type:character varying"`
	UserId          int       `gorm:"type:integer"`
	ChannelId       int       `gorm:"type:integer"`
	Status          int       `gorm:"type:integer"` //0-draft 1-publish 2-unpublish
	CoverImage      string    `gorm:"type:character varying"`
	ThumbnailImage  string    `gorm:"type:character varying"`
	MetaTitle       string    `gorm:"type:character varying"`
	MetaDescription string    `gorm:"type:character varying"`
	Keyword         string    `gorm:"type:character varying"`
	CategoriesId    string    `gorm:"type:character varying"`
	RelatedArticles string    `gorm:"type:character varying"`
	Feature         int       `gorm:"DEFAULT:0"`
	ViewCount       int       `gorm:"DEFAULT:0"`
	CreateTime      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	PublishedTime   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ImageAltTag     string    `gorm:"type:character varying"`
	Author          string    `gorm:"type:character varying"`
	SortOrder       int       `gorm:"type:integer"`
	Excerpt         string    `gorm:"type:character varying"`
	ReadingTime     int       `gorm:"type:integer"`
	Tags            string    `gorm:"type:character varying"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy       int       `gorm:"type:integer"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	ModifiedOn      time.Time `gorm:"DEFAULT:NULL"`
	IsActive        int       `gorm:"type:integer"`
	IsDeleted       int       `gorm:"DEFAULT:0"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL"`
}

type TblChannelEntryField struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName      string    `gorm:"type:character varying"`
	FieldValue     string    `gorm:"type:character varying"`
	ChannelEntryId int       `gorm:"type:integer"`
	FieldId        int       `gorm:"type:integer"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy      int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	DeletedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId        int               `gorm:"type:integer"`
	ProfilePage     string            `gorm:"type:character varying"`
	ProfileName     string            `gorm:"type:character varying"`
	ProfileSlug     string            `gorm:"type:character varying"`
	CompanyLogo     string            `gorm:"type:character varying"`
	CompanyName     string            `gorm:"type:character varying"`
	CompanyLocation string            `gorm:"type:character varying"`
	About           string            `gorm:"type:character varying"`
	Linkedin        string            `gorm:"type:character varying"`
	Website         string            `gorm:"type:character varying"`
	Twitter         string            `gorm:"type:character varying"`
	SeoTitle        string            `gorm:"type:character varying"`
	SeoDescription  string            `gorm:"type:character varying"`
	SeoKeyword      string            `gorm:"type:character varying"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:integer"`
	CreatedOn       time.Time         `gorm:"type:timestamp without time zone"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL"`
	ModifiedOn      time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL"`
	DeletedOn       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblMemberNotesHighlights struct {
	Id                      int               `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId                int               `gorm:"type:integer"`
	PageId                  int               `gorm:"type:integer"`
	NotesHighlightsContent  string            `gorm:"type:character varying"`
	NotesHighlightsType     string            `gorm:"type:character varying"`
	HighlightsConfiguration datatypes.JSONMap `gorm:"type:jsonb"`
	CreatedBy               int               `gorm:"type:integer"`
	CreatedOn               time.Time         `gorm:"type:timestamp without time zone"`
	ModifiedBy              int               `gorm:"type:integer"`
	ModifiedOn              time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy               int               `gorm:"type:integer"`
	DeletedOn               time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted               int               `gorm:"type:integer"`
}

type TblStorageType struct {
	Id           int               `gorm:"primaryKey;auto_increment;type:serial"`
	Local        string            `gorm:"type:character varying"`
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:character varying"`
}

type TblEmailConfiguration struct {
	Id           int               `gorm:"primaryKey;auto_increment;type:serial"`
	StmpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:character varying"`
}

type TblGeneralSetting struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	CompanyName    string    `gorm:"type:character varying"`
	LogoPath       string    `gorm:"type:character varying"`
	ExpandLogoPath string    `gorm:"type:character varying"`
	DateFormat     string    `gorm:"type:character varying"`
	TimeFormat     string    `gorm:"type:character varying"`
	TimeZone       string    `gorm:"type:character varying"`
	LanguageId     int       `gorm:"type:integer"`
	ModifiedBy     int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
}
type TblMemberSetting struct {
	Id                int       `gorm:"primaryKey;auto_increment;type:serial"`
	AllowRegistration int       `gorm:"type:integer"`
	MemberLogin       string    `gorm:"type:character varying"`
	ModifiedBy        int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	NotificationUsers string    `gorm:"type:character varying"`
}

type TblGraphqlSettings struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Token       string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Duration    string    `gorm:"type:character varying"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
}

type TblTimezone struct {
	Id       int    `gorm:"primaryKey;auto_increment;type:serial"`
	Timezone string `gorm:"type:character varying"`
}

func MigrationTables() {

	err := controllers.DB.AutoMigrate(

		TblAccessControl{},
		TblAccessControlPages{},
		TblAccessControlUserGroup{},
		TblCategories{},
		TblChannelCategorie{},
		TblChannelEntrie{},
		TblChannelEntryField{},
		TblChannel{},
		TblEmailTemplate{},
		TblFieldGroup{},
		TblFieldOption{},
		TblFieldType{},
		TblField{},
		TblLanguage{},
		TblMemberGroup{},
		TblMemberNotesHighlights{},
		TblMemberProfile{},
		TblMember{},
		TblModulePermission{},
		TblModule{},
		TblRole{},
		TblRolePermission{},
		TblRoleUser{},
		TblUser{},
		TblGroupField{},
		TblUserPersonalize{},
		TblStorageType{},
		TblEmailConfiguration{},
		TblGeneralSetting{},
		TblMemberSetting{},
		TblGraphqlSettings{},
		TblTimezone{},
	)

	if err != nil {

		panic(err)

	}

}
