package mysql

import (
	"spurt-cms/controllers"
	"time"

	"gorm.io/datatypes"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblUser struct {
	Id                int       `gorm:"primaryKey"`
	Uuid              string    `gorm:"type:varchar(255)"`
	FirstName         string    `gorm:"type:varchar(255)"`
	LastName          string    `gorm:"type:varchar(255)"`
	RoleId            TblRole   `gorm:"type:int;foreignkey:Id"`
	Email             string    `gorm:"type:varchar(255)"`
	Username          string    `gorm:"type:varchar(255)"`
	Password          string    `gorm:"type:varchar(255)"`
	MobileNo          string    `gorm:"type:varchar(255)"`
	IsActive          int       `gorm:"type:int"`
	ProfileImage      string    `gorm:"type:varchar(255)"`
	ProfileImagePath  string    `gorm:"type:varchar(255)"`
	DataAccess        int       `gorm:"type:int"`
	DefaultLanguageId int       `gorm:"type:int"`
	CreatedOn         time.Time `gorm:"type:datetime"`
	CreatedBy         int       `gorm:"type:int"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	LastLogin         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:int"`
	DeletedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblCategories struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	CategoryName string    `gorm:"type:varchar(255)"`
	CategorySlug string    `gorm:"type:varchar(255)"`
	Description  string    `gorm:"type:varchar(255)"`
	ImagePath    string    `gorm:"type:varchar(255)"`
	ParentId     int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	CreatedBy    int       `gorm:"type:int"`
	ModifiedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted    int       `gorm:"type:int"`
	DeletedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:int;DEFAULT:NULL;type:int"`
}

type TblChannel struct {
	Id                 int       `gorm:"primaryKey;auto_increment"`
	ChannelName        string    `gorm:"type:varchar(255)"`
	ChannelDescription string    `gorm:"type:varchar(255)"`
	SlugName           string    `gorm:"type:varchar(255)"`
	FieldGroupId       int       `gorm:"type:int"`
	IsActive           int       `gorm:"type:int"`
	IsDeleted          int       `gorm:"type:int"`
	CreatedOn          time.Time `gorm:"type:datetime"`
	CreatedBy          int       `gorm:"type:int"`
	ModifiedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	Uuid             string    `gorm:"type:varchar(255)"`
	FirstName        string    `gorm:"type:varchar(255)"`
	LastName         string    `gorm:"type:varchar(255)"`
	Email            string    `gorm:"type:varchar(255)"`
	MobileNo         string    `gorm:"type:varchar(255)"`
	IsActive         int       `gorm:"type:int"`
	ProfileImage     string    `gorm:"type:varchar(255)"`
	ProfileImagePath string    `gorm:"type:varchar(255)"`
	LastLogin        int       `gorm:"type:int"`
	MemberGroupId    int       `gorm:"type:int"`
	Password         string    `gorm:"type:varchar(255)"`
	Username         string    `gorm:"type:varchar(255)"`
	Otp              int       `gorm:"DEFAULT:NULL;type:int"`
	OtpExpiry        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblAccessControl struct {
	Id                int       `gorm:"primaryKey;auto_increment"`
	AccessControlName string    `gorm:"type:varchar(255)"`
	AccessControlSlug string    `gorm:"type:varchar(255)"`
	CreatedOn         time.Time `gorm:"type:datetime"`
	CreatedBy         int       `gorm:"type:int"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted         int       `gorm:"type:int"`
	DeletedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblAccessControlPages struct {
	Id                       int       `gorm:"primaryKey;auto_increment"`
	AccessControlUserGroupId int       `gorm:"type:int"`
	SpacesId                 int       `gorm:"type:int"`
	PageGroupId              int       `gorm:"type:int"`
	PageId                   int       `gorm:"type:int"`
	CreatedOn                time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy                int       `gorm:"type:int"`
	ModifiedOn               time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy               int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted                int       `gorm:"type:int"`
	DeletedOn                time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy                int       `gorm:"DEFAULT:NULL;type:int"`
	ChannelId                int       `gorm:"type:int"`
	EntryId                  int       `gorm:"type:int"`
}

type TblAccessControlUserGroup struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	AccessControlId int       `gorm:"type:int"`
	MemberGroupId   int       `gorm:"type:int"`
	CreatedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy       int       `gorm:"type:int"`
	ModifiedOn      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted       int       `gorm:"type:int"`
	DeletedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblEmailTemplate struct {
	Id                  int       `gorm:"primaryKey;auto_increment"`
	TemplateName        string    `gorm:"type:varchar(255)"`
	TemplateSubject     string    `gorm:"type:varchar(255)"`
	TemplateMessage     string    `gorm:"type:text"`
	TemplateDescription string    `gorm:"type:text"`
	ModuleId            int       `gorm:"type:int"`
	CreatedOn           time.Time `gorm:"type:datetime"`
	CreatedBy           int       `gorm:"type:int"`
	ModifiedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn           time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy           int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted           int       `gorm:"DEFAULT:0"`
	IsActive            int       `gorm:"type:int"`
	IsDefault           int       `gorm:"type:int"`
}

type TblField struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	FieldName        string    `gorm:"type:varchar(255)"`
	FieldDesc        string    `gorm:"type:varchar(255)"`
	FieldTypeId      int       `gorm:"type:int"`
	MandatoryField   int       `gorm:"type:int"`
	OptionExist      int       `gorm:"type:int"`
	InitialValue     string    `gorm:"type:varchar(255)"`
	Placeholder      string    `gorm:"type:varchar(255)"`
	OrderIndex       int       `gorm:"type:int"`
	ImagePath        string    `gorm:"type:varchar(255)"`
	DatetimeFormat   string    `gorm:"type:varchar(255)"`
	TimeFormat       string    `gorm:"type:varchar(255)"`
	Url              string    `gorm:"type:varchar(255)"`
	SectionParentId  int       `gorm:"type:int"`
	CharacterAllowed int       `gorm:"type:int"`
	CreatedOn        time.Time `gorm:"type:datetime"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblFieldGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	GroupName  string    `gorm:"type:varchar(255)"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	CreatedBy  int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblFieldOption struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	OptionName  string    `gorm:"type:varchar(255)"`
	OptionValue string    `gorm:"type:varchar(255)"`
	FieldId     int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblFieldType struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	TypeName   string    `gorm:"type:varchar(255)"`
	TypeSlug   string    `gorm:"type:varchar(255)"`
	IsActive   int       `gorm:"type:int"`
	IsDeleted  int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	ModifiedBy int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblLanguage struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	LanguageName string    `gorm:"type:varchar(255)"`
	LanguageCode string    `gorm:"type:varchar(255)"`
	ImagePath    string    `gorm:"type:varchar(255)"`
	IsStatus     int       `gorm:"type:int"`
	IsDefault    int       `gorm:"type:int"`
	JsonPath     string    `gorm:"type:varchar(255)"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	CreatedBy    int       `gorm:"type:int"`
	ModifiedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted    int       `gorm:"DEFAULT:0"`
}

type TblModule struct {
	Id                   int       `gorm:"primaryKey"`
	ModuleName           string    `gorm:"type:varchar(255)"`
	IsActive             int       `gorm:"type:int"`
	DefaultModule        int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	IconPath             string    `gorm:"type:varchar(255)"`
	AssignPermission     int       `gorm:"type:int"`
	Description          string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	MenuType             string    `gorm:"type:varchar(255)"`
	FullAccessPermission int       `gorm:"type:int"`
	GroupFlg             int       `gorm:"type:int"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey"`
	RouteName            string    `gorm:"type:varchar(255)"`
	DisplayName          string    `gorm:"type:varchar(255)"`
	SlugName             string    `gorm:"type:varchar(255)"`
	Description          string    `gorm:"type:varchar(255)"`
	ModuleId             int       `gorm:"type:int"`
	FullAccessPermission int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	AssignPermission     int       `gorm:"type:int"`
	BreadcrumbName       string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn           time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	RoleId       int       `gorm:"type:int"`
	PermissionId int       `gorm:"type:int"`
	CreatedBy    int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime"`
}

type TblChannelCategorie struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	ChannelId  int       `gorm:"type:int"`
	CategoryId string    `gorm:"type:varchar(255)"`
	CreatedAt  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
}

type TblGroupField struct {
	Id        int `gorm:"primaryKey;auto_increment"`
	ChannelId int `gorm:"type:int"`
	FieldId   int `gorm:"type:int"`
}

type TblUserPersonalize struct {
	Id                  int       `gorm:"primaryKey;auto_increment"`
	UserId              int       `gorm:"type:int"`
	MenuBackgroundColor string    `gorm:"type:varchar(255)"`
	FontColor           string    `gorm:"type:varchar(255)"`
	IconColor           string    `gorm:"type:varchar(255)"`
	LogoPath            string    `gorm:"type:varchar(255)"`
	ExpandLogoPath      string    `gorm:"type:varchar(255)"`
	CreatedOn           time.Time `gorm:"type:datetime"`
	ModifiedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblRoleUser struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	RoleId     int       `gorm:"type:int"`
	UserId     int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblChannelEntrie struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	Title           string    `gorm:"type:varchar(255)"`
	Slug            string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:varchar(255)"`
	UserId          int       `gorm:"type:int"`
	ChannelId       int       `gorm:"type:int"`
	Status          int       `gorm:"type:int"` //0-draft 1-publish 2-unpublish
	CoverImage      string    `gorm:"type:varchar(255)"`
	ThumbnailImage  string    `gorm:"type:varchar(255)"`
	MetaTitle       string    `gorm:"type:varchar(255)"`
	MetaDescription string    `gorm:"type:varchar(255)"`
	Keyword         string    `gorm:"type:varchar(255)"`
	CategoriesId    string    `gorm:"type:varchar(255)"`
	RelatedArticles string    `gorm:"type:varchar(255)"`
	Feature         int       `gorm:"DEFAULT:0"`
	ViewCount       int       `gorm:"DEFAULT:0"`
	CreateTime      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	PublishedTime   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ImageAltTag     string    `gorm:"type:varchar(255)"`
	Author          string    `gorm:"type:varchar(255)"`
	SortOrder       int       `gorm:"type:int"`
	Excerpt         string    `gorm:"type:varchar(255)"`
	ReadingTime     int       `gorm:"type:int"`
	Tags            string    `gorm:"type:varchar(255)"`
	CreatedOn       time.Time `gorm:"type:datetime"`
	CreatedBy       int       `gorm:"type:int"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn      time.Time `gorm:"DEFAULT:NULL;type:datetime"`
	IsActive        int       `gorm:"type:int"`
	IsDeleted       int       `gorm:"DEFAULT:0"`
	DeletedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblChannelEntryField struct {
	Id             int       `gorm:"primaryKey;auto_increment"`
	FieldName      string    `gorm:"type:varchar(255)"`
	FieldValue     string    `gorm:"type:varchar(255)"`
	ChannelEntryId int       `gorm:"type:int"`
	FieldId        int       `gorm:"type:int"`
	CreatedOn      time.Time `gorm:"type:datetime"`
	CreatedBy      int       `gorm:"type:int"`
	ModifiedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy      int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment"`
	MemberId        int               `gorm:"type:int"`
	ProfilePage     string            `gorm:"type:varchar(255)"`
	ProfileName     string            `gorm:"type:varchar(255)"`
	ProfileSlug     string            `gorm:"type:varchar(255)"`
	CompanyLogo     string            `gorm:"type:varchar(255)"`
	CompanyName     string            `gorm:"type:varchar(255)"`
	CompanyLocation string            `gorm:"type:varchar(255)"`
	About           string            `gorm:"type:varchar(255)"`
	Linkedin        string            `gorm:"type:varchar(255)"`
	Website         string            `gorm:"type:varchar(255)"`
	Twitter         string            `gorm:"type:varchar(255)"`
	SeoTitle        string            `gorm:"type:varchar(255)"`
	SeoDescription  string            `gorm:"type:varchar(255)"`
	SeoKeyword      string            `gorm:"type:varchar(255)"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:int"`
	CreatedOn       time.Time         `gorm:"type:datetime"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn      time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn       time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblMemberNotesHighlights struct {
	Id                      int               `gorm:"primaryKey;auto_increment"`
	MemberId                int               `gorm:"type:int"`
	PageId                  int               `gorm:"type:int"`
	NotesHighlightsContent  string            `gorm:"type:varchar(255)"`
	NotesHighlightsType     string            `gorm:"type:varchar(255)"`
	HighlightsConfiguration datatypes.JSONMap `gorm:"type:jsonb"`
	CreatedBy               int               `gorm:"type:int"`
	CreatedOn               time.Time         `gorm:"type:datetime"`
	ModifiedBy              int               `gorm:"type:int"`
	ModifiedOn              time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy               int               `gorm:"type:int"`
	DeletedOn               time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted               int               `gorm:"type:int"`
}

type TblStorageType struct {
	Id           int               `gorm:"primaryKey;auto_increment"`
	Local        string            `gorm:"type:varchar(255)"`
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
}

type TblEmailConfiguration struct {
	Id           int               `gorm:"primaryKey;auto_increment;"`
	StmpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
}

type TblGeneralSetting struct {
	Id             int       `gorm:"primaryKey;auto_increment"`
	CompanyName    string    `gorm:"type:varchar(255)"`
	LogoPath       string    `gorm:"type:varchar(255)"`
	ExpandLogoPath string    `gorm:"type:varchar(255)"`
	DateFormat     string    `gorm:"type:varchar(255)"`
	TimeFormat     string    `gorm:"type:varchar(255)"`
	TimeZone       string    `gorm:"type:varchar(255)"`
	LanguageId     int       `gorm:"type:int"`
	ModifiedBy     int       `gorm:"type:int"`
	ModifiedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblMemberSetting struct {
	Id                int       `gorm:"primaryKey;auto_increment;"`
	AllowRegistration int       `gorm:"type:int"`
	MemberLogin       string    `gorm:"type:varchar(255)"`
	ModifiedBy        int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	NotificationUsers string    `gorm:"type:varchar(255)"`
}

type TblGraphqlSettings struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Token       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Duration    string    `gorm:"type:varchar(255)"`
	CreatedBy   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int;DEFAULT:0"`
}

type TblTimezone struct {
	Id       int    `gorm:"primaryKey;auto_increment"`
	Timezone string `gorm:"type:varchar(255)"`
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
