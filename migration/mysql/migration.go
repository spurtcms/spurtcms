package mysql

import (
	"spurt-cms/controllers"
	"time"

	"gorm.io/datatypes"
)

type TblRoles struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:LONGTEXT"`
	Slug        string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId    int       `gorm:"type:int;"`
}

type TblUsers struct {
    Id                int       `gorm:"primaryKey;autoIncrement;type:int"`
    Uuid              string    `gorm:"type:varchar(255)"`
    FirstName         string    `gorm:"type:varchar(255)"`
    LastName          string    `gorm:"type:varchar(255)"`
    RoleId            int       `gorm:"type:int"`
    Email             string    `gorm:"type:varchar(255)"`
    Username          string    `gorm:"type:varchar(255)"`
    Password          string    `gorm:"type:varchar(255)"`
    MobileNo          string    `gorm:"type:varchar(255)"`
    IsActive          int       `gorm:"type:int"`
    ProfileImage      string    `gorm:"type:varchar(255)"`
    ProfileImagePath  string    `gorm:"type:varchar(255)"`
    StorageType       string    `gorm:"type:varchar(255)"`
    DataAccess        int       `gorm:"type:int"`
    DefaultLanguageId int       `gorm:"type:int"`
    CreatedOn         time.Time `gorm:"type:datetime"`
    CreatedBy         int       `gorm:"type:int"`
    ModifiedOn        time.Time `gorm:"type:datetime;default:NULL"`
    ModifiedBy        int       `gorm:"type:int;default:NULL"`
    LastLogin         time.Time `gorm:"type:datetime;default:NULL"`
    IsDeleted         int       `gorm:"type:int"`
    DeletedOn         time.Time `gorm:"type:datetime;default:NULL"`
    DeletedBy         int       `gorm:"type:int;default:NULL"`
    TenantId          string    `gorm:"type:varchar(255)"`
    Otp               int       `gorm:"type:int;default:NULL"`
    OtpExpiry         time.Time `gorm:"type:datetime;default:NULL"`
    S3FolderName      string    `gorm:"type:varchar(255)"`
    ArticleCount      int       `gorm:"type:int"`
    TotalCount        int       `gorm:"type:int"`
}

type TblCategories struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	CategoryName string    `gorm:"type:varchar(255)"`
	CategorySlug string    `gorm:"type:varchar(255)"`
	Description  string    `gorm:"type:LONGTEXT"`
	ImagePath    string    `gorm:"type:varchar(255)"`
	ParentId     int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	CreatedBy    int       `gorm:"type:int"`
	ModifiedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted    int       `gorm:"type:int"`
	DeletedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:int;DEFAULT:NULL;type:int"`
	TenantId     int       `gorm:"type:int;"`
}

type TblChannels struct {
    Id                 int       `gorm:"primaryKey;autoIncrement;type:int"`
    ChannelName        string    `gorm:"type:varchar(255)"`
    ChannelUniqueId    string    `gorm:"type:varchar(255)"`
    ChannelDescription string    `gorm:"type:varchar(255)"`
    SlugName           string    `gorm:"type:varchar(255)"`
    FieldGroupId       int       `gorm:"type:int"`
    IsActive           int       `gorm:"type:int"`
    IsDeleted          int       `gorm:"type:int"`
    CreatedOn          time.Time `gorm:"type:datetime"`
    CreatedBy          int       `gorm:"type:int"`
    ModifiedOn         time.Time `gorm:"type:datetime;default:NULL"`
    ModifiedBy         int       `gorm:"type:int;default:NULL"`
    ChannelType        string    `gorm:"type:varchar(255)"`
    CollectionCount    int       `gorm:"type:int"`
    CloneCount         int       `gorm:"type:int"`
    TenantId           string    `gorm:"type:varchar(255)"`
    ImagePath          string    `gorm:"type:varchar(255)"`
}


type TblMemberGroups struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:LONGTEXT"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId    int       `gorm:"type:int;"`
}

type TblMembers struct {
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
	StorageType      string    `gorm:"type:varchar(255)"`
	TenantId         int       `gorm:"type:int;"`
}

type TblAccessControls struct {
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
	TenantId          int       `gorm:"type:int;"`
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
	TenantId                 int       `gorm:"type:int;"`
}

type TblAccessControlUserGroups struct {
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
	TenantId        int       `gorm:"type:int;"`
}

type TblEmailTemplates struct {
	Id                  int       `gorm:"primaryKey;auto_increment"`
	TemplateName        string    `gorm:"type:varchar(255)"`
	TemplateSlug        string    `gorm:"type:varchar(255)"`
	TemplateSubject     string    `gorm:"type:varchar(255)"`
	TemplateMessage     string    `gorm:"type:LONGTEXT"`
	TemplateDescription string    `gorm:"type:LONGTEXT"`
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
	TenantId            int       `gorm:"type:int;"`
}

type TblFields struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	FieldName        string    `gorm:"type:varchar(255)"`
	FieldDesc        string    `gorm:"type:LONGTEXT"`
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
	TenantId         int       `gorm:"type:int;"`
}

type TblFieldGroups struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	GroupName  string    `gorm:"type:varchar(255)"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	CreatedBy  int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId   string    `gorm:"type:int;"`
}

type TblFieldOptions struct {
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
	TenantId    int       `gorm:"type:int;"`
	OrderIndex  int       `gorm:"DEFAULT:NULL;type:int"`
}

type TblFieldTypes struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	TypeName   string    `gorm:"type:varchar(255)"`
	TypeSlug   string    `gorm:"type:varchar(255)"`
	IsActive   int       `gorm:"type:int"`
	IsDeleted  int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	ModifiedBy int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:int;"`
}

type TblLanguages struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	LanguageName string    `gorm:"type:varchar(255)"`
	LanguageCode string    `gorm:"type:varchar(255)"`
	ImagePath    string    `gorm:"type:varchar(255)"`
	StorageType  string    `gorm:"type:varchar(255)"`
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
	TenantId     int       `gorm:"type:int;"`
}

type TblModules struct {
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
	TenantId             int       `gorm:"type:int;"`
}

type TblModulePermissions struct {
	Id                   int       `gorm:"primaryKey"`
	RouteName            string    `gorm:"type:varchar(255)"`
	DisplayName          string    `gorm:"type:varchar(255)"`
	SlugName             string    `gorm:"type:varchar(255)"`
	Description          string    `gorm:"type:LONGTEXT"`
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
	TenantId             int       `gorm:"type:int;"`
}

type TblRolePermissions struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	RoleId       int       `gorm:"type:int"`
	PermissionId int       `gorm:"type:int"`
	CreatedBy    int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	TenantId     int       `gorm:"type:int;"`
}

type TblChannelCategories struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	ChannelId  int       `gorm:"type:int"`
	CategoryId string    `gorm:"type:varchar(255)"`
	CreatedAt  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	TenantId   string    `gorm:"type:int;"`
}

type TblGroupFields struct {
	Id        int    `gorm:"primaryKey;auto_increment"`
	ChannelId int    `gorm:"type:int"`
	FieldId   int    `gorm:"type:int"`
	TenantId  string `gorm:"type:int;"`
}

type TblUserPersonalizes struct {
	Id                  int       `gorm:"primaryKey;auto_increment"`
	UserId              int       `gorm:"type:int"`
	MenuBackgroundColor string    `gorm:"type:varchar(255)"`
	FontColor           string    `gorm:"type:varchar(255)"`
	IconColor           string    `gorm:"type:varchar(255)"`
	LogoPath            string    `gorm:"type:varchar(255)"`
	ExpandLogoPath      string    `gorm:"type:varchar(255)"`
	CreatedOn           time.Time `gorm:"type:datetime"`
	ModifiedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId            int       `gorm:"type:int;"`
}

type TblRoleUsers struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	RoleId     int       `gorm:"type:int"`
	UserId     int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:int;"`
}

type TblChannelEntries struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	Title           string    `gorm:"type:varchar(255)"`
	Slug            string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"LONGTEXT"`
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
	TenantId        int       `gorm:"type:int;"`
	Uuid            string    `gorm:"type:varchar(255)"`
	ParentId        int       `gorm:"type:int;"`
	OrderIndex      int       `gorm:"type:int;"`
	MembergroupId   string    `gorm:"type:varchar(255)"`
	CtaId           int       `gorm:"type:int"`
}

type TblChannelEntryFields struct {
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
	TenantId       int       `gorm:"type:int;"`
}

type TblMemberProfiles struct {
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
	StorageType     string            `gorm:"type:varchar(255)"`
	ClaimDate       time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId        int               `gorm:"type:int;"`
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
	TenantId                int               `gorm:"type:int;"`
}

type TblStorageTypes struct {
	Id           int               `gorm:"primaryKey;auto_increment"`
	Local        string            `gorm:"type:varchar(255)"`
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
	TenantId     int               `gorm:"type:int;"`
}

type TblEmailConfigurations struct {
	Id           int               `gorm:"primaryKey;auto_increment;"`
	StmpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:varchar(255)"`
	TenantId     int               `gorm:"type:int;"`
}

type TblGeneralSettings struct {
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
	TenantId       int       `gorm:"type:int;"`
	StorageType    string    `gorm:"type:varchar(255)"`
}

type TblMemberSettings struct {
	Id                int       `gorm:"primaryKey;auto_increment;"`
	AllowRegistration int       `gorm:"type:int"`
	MemberLogin       string    `gorm:"type:varchar(255)"`
	ModifiedBy        int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	NotificationUsers string    `gorm:"type:varchar(255)"`
	TenantId          int       `gorm:"type:int;"`
}

type TblGraphqlSettings struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	TokenName   string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:LONGTEXT"`
	Duration    string    `gorm:"type:varchar(255)"`
	CreatedBy   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int;DEFAULT:0"`
	Token       string    `gorm:"type:varchar(255)"`
	IsDefault   int       `gorm:"type:int;DEFAULT:0"`
	ExpiryTime  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId    int       `gorm:"type:int;"`
}

type TblTimezones struct {
	Id       int    `gorm:"primaryKey;auto_increment"`
	Timezone string `gorm:"type:varchar(255)"`
	tenantid string `gorm:"type:int;"`
}

type TblForms struct {
	Id                   int       `gorm:"primaryKey;autoIncrement"`
	FormTitle            string    `gorm:"type:varchar(255)"`
	FormSlug             string    `gorm:"type:varchar(255)"`
	FormData             string    `gorm:"type:LONGTEXT"`
	Status               int       `gorm:"type:int"`
	IsActive             int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:timestamp;default:NULL"`
	ModifiedBy           int       `gorm:"type:int;default:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp;default:NULL"`
	DeletedBy            int       `gorm:"type:int;default:NULL"`
	DeletedOn            time.Time `gorm:"type:timestamp;default:NULL"`
	IsDeleted            int       `gorm:"type:int;default:0"`
	TenantId             int       `gorm:"type:int"`
	Uuid                 string    `gorm:"type:varchar(255)"`
	FormImagePath        string    `gorm:"type:varchar(255)"`
	FormDescription      string    `gorm:"type:varchar(255)"`
	ChannelId            int       `gorm:"type:int"`
	ChannelName          string    `gorm:"type:varchar(255)"`
	FormPreviewImagepath string    `gorm:"type:varchar(255)"`
	FormPreviewImagename string    `gorm:"type:varchar(255)"`
}

type TblFormResponse struct {
	Id           int       `gorm:"primaryKey;autoIncrement"`
	FormId       int       `gorm:"type:int"`
	FormResponse string    `gorm:"type:LONGTEXT"`
	UserId       int       `gorm:"type:int"`
	IsActive     int       `gorm:"type:int"`
	IsDeleted    int       `gorm:"type:int;default:0"`
	CreatedBy    int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:timestamp;default:NULL"`
	TenantId     int       `gorm:"type:int"`
	EntryId      int       `gorm:"type:int"`
}

type TblFormRegistrations struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId     int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int;DEFAULT:0"`
	TenantId   string    `gorm:"type:int;"`
}

type TblFormFields struct {
	Id            int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId        int       `gorm:"type:int"`
	MasterFieldId int       `gorm:"type:int"`
	FieldName     string    `gorm:"type:varchar(255)"`
	Mandatory     int       `gorm:"type:int"`
	OptionExists  int       `gorm:"type:int"`
	CreatedBy     int       `gorm:"type:int"`
	CreatedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted     int       `gorm:"type:int;DEFAULT:0"`
	TenantId      int       `gorm:"type:int;"`
}

type TblFormValues struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldId    int       `gorm:"type:int"`
	FieldValue string    `gorm:"type:varchar(255)"`
	RegisterId int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int;DEFAULT:0"`
	TenantId   string    `gorm:"type:int;"`
}

type TblFormMasterFields struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName        string    `gorm:"type:varchar(255)"`
	FieldDescription string    `gorm:"type:LONGTEXT"`
	IconPath         string    `gorm:"type:varchar(255)"`
	IsActive         int       `gorm:"type:int"`
	CreatedBy        int       `gorm:"type:int"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int;DEFAULT:0"`
	TenantId         int       `gorm:"type:int;"`
}

type TblFormFieldOptions struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldId    int       `gorm:"type:int"`
	Options    string    `gorm:"type:varchar(255)"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int;DEFAULT:0"`
	TenantId   string    `gorm:"type:int;"`
}

type TblPageTypes struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	Title      string    `gorm:"type:varchar(255)"`
	Slug       string    `gorm:"type:varchar(255)"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	TenantId   string    `gorm:"type:int;"`
}

type TblTemplates struct {
	Id               int       `gorm:"primaryKey;auto_increment;"`
	TemplateName     string    `gorm:"type:varchar(255)"`
	TemplateSlug     string    `gorm:"type:varchar(255)"`
	TemplateModuleId int       `gorm:"type:int"`
	ImageName        string    `gorm:"type:varchar(255)"`
	ImagePath        string    `gorm:"type:varchar(255)"`
	DeployLink       string    `gorm:"type:varchar(255)"`
	GithubLink       string    `gorm:"type:varchar(255)"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:int"`
	IsActive         int       `gorm:"type:int"`
	IsDeleted        int       `gorm:"type:int;DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:int;DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:int"`
	TenantId         int       `gorm:"type:int;"`
	ChannelId        int       `gorm:"type:int;"`
}
type TblBlocks struct {
    Id               int       `gorm:"primaryKey;autoIncrement;type:int"`
    Title            string    `gorm:"type:varchar(255)"`
    SlugName         string    `gorm:"type:varchar(255)"`
    ChannelSlugname  string    `gorm:"type:varchar(255)"`
    BlockDescription string    `gorm:"type:text"`
    BlockContent     string    `gorm:"type:text"`
    BlockCss         string    `gorm:"type:text"`
    IconImage        string    `gorm:"type:varchar(255)"`
    CoverImage       string    `gorm:"type:varchar(255)"`
    Prime            int       `gorm:"type:int"`
    IsActive         int       `gorm:"type:int"`
    TenantId         string    `gorm:"type:varchar(255)"`
    CreatedOn        time.Time `gorm:"type:datetime;default:NULL"`
    CreatedBy        int       `gorm:"type:int"`
    ModifiedBy       int       `gorm:"type:int"`
    ModifiedOn       time.Time `gorm:"type:datetime;default:NULL"`
    DeletedOn        time.Time `gorm:"type:datetime;default:NULL"`
    DeletedBy        int       `gorm:"type:int;default:NULL"`
    IsDeleted        int       `gorm:"type:int;default:0"`
    ChannelId        int       `gorm:"type:int"`
}


type TblBlockTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;"`
	BlockId   int       `gorm:"type:int"`
	TagId     int       `gorm:"type:int"`
	TagName   string    `gorm:"type:varchar(255)"`
	TenantId  string    `gorm:"type:int"`
	CreatedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:int"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:int;DEFAULT:0"`
}

type TblBlockMstrTags struct {
	Id        int       `gorm:"primaryKey;auto_increment"`
	TagTitle  string    `gorm:"type:varchar(255)"`
	TenantId  string    `gorm:"type:int"`
	CreatedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:int"`
}

type TblBlockCollections struct {
	Id        int       `gorm:"primaryKey;auto_increment;"`
	UserId    int       `gorm:"type:int"`
	BlockId   int       `gorm:"type:int"`
	TenantId  string    `gorm:"type:int"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:int;DEFAULT:0"`
}

type TblMstrTenant struct {
	Id            int       `gorm:"primaryKey;auto_increment;"`
	TenantId      int       `gorm:"type:int"`
	S3StoragePath string    `gorm:"type:varchar(255)"`
	IsDeleted     int       `gorm:"type:int;DEFAULT:0"`
	DeletedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy     int       `gorm:"type:int;DEFAULT:NULL"`
}

type TblWebhooks struct {
	Id            int                    `gorm:"primaryKey;auto_increment;"`
	WebhookName   string                 `gorm:"type:varchar(255)"`
	PayloadType   string                 `gorm:"type:varchar(255)"`
	RequestMethod string                 `gorm:"type:varchar(255)"`
	RequestUrl    string                 `gorm:"type:varchar(255)"`
	EventType     string                 `gorm:"type:varchar(255)"`
	IsActive      int                    `gorm:"type:int;DEFAULT:NULL"`
	Headers       map[string]interface{} `gorm:"type:json"`
	Fields        map[string]interface{} `gorm:"type:json"`
	TenantId      int                    `gorm:"type:int;DEFAULT:NULL"`
	CreatedBy     int                    `gorm:"type:int"`
	CreatedOn     time.Time              `gorm:"type:datetime"`
	ModifiedBy    int                    `gorm:"type:int;DEFAULT:NULL"`
	ModifiedOn    time.Time              `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted     int                    `gorm:"type:int;DEFAULT:0"`
	DeletedBy     int                    `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn     time.Time              `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblTemplateModules struct {
	Id                 int       `gorm:"primaryKey;auto_increment;"`
	TemplateModuleName string    `gorm:"type:varchar(255)"`
	TemplateModuleSlug string    `gorm:"type:varchar(255)"`
	Description        string    `gorm:"type:varchar(600)"`
	IsActive           int       `gorm:"type:int;DEFAULT:1"`
	CreatedBy          int       `gorm:"type:int"`
	CreatedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId           int       `gorm:"type:int"`
	ChannelId          int       `gorm:"type:int"`
	IsDeleted          int       `gorm:"type:int;DEFAULT:0"`
	DeletedBy          int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
}

type TblApps struct {
	Id            int       `gorm:"primaryKey;auto_increment"`
	Title         string    `gorm:"type:varchar(255)"`
	Description   string    `gorm:"type:varchar(255)"`
	FieldJsonPath string    `gorm:"type:varchar(255)"`
	IconName      string    `gorm:"type:varchar(255)"`
	IconPath      string    `gorm:"type:varchar(255)"`
	CreatedBy     int       `gorm:"type:varchar(255)"`
	CreatedOn     time.Time `gorm:"type:datetime"`
	ModifiedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsActive      int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted     int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId      int       `gorm:"type:integer;DEFAULT:NULL"`
}
type TblAiPrompt struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	AppId        int       `gorm:"type:varchar(255)"`
	MasterId     int       `gorm:"type:integer;DEFAULT:NULL"`
	ChildId      int       `gorm:"type:integer;DEFAULT:NULL"`
	PromptLevel  int       `gorm:"type:integer;DEFAULT:NULL"`
	LanguageId   int       `gorm:"type:integer;DEFAULT:NULL"`
	SystemPrompt string    `gorm:"type:varchar(255)"`
	UserPrompt   string    `gorm:"type:text"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	ModifiedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsActive     int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId     int       `gorm:"type:integer;DEFAULT:NULL"`
}

func MigrationTables() {

	err := controllers.DB.AutoMigrate(

		TblAccessControls{},
		TblAccessControlPages{},
		TblAccessControlUserGroups{},
		TblCategories{},
		TblChannelCategories{},
		TblChannelEntries{},
		TblChannelEntryFields{},
		TblChannels{},
		TblEmailTemplates{},
		TblFieldGroups{},
		TblFieldOptions{},
		TblFieldTypes{},
		TblFields{},
		TblLanguages{},
		TblMemberGroups{},
		TblMemberNotesHighlights{},
		TblMemberProfiles{},
		TblMembers{},
		TblModulePermissions{},
		TblModules{},
		TblRoles{},
		TblRolePermissions{},
		TblRoleUsers{},
		TblUsers{},
		TblGroupFields{},
		TblUserPersonalizes{},
		TblStorageTypes{},
		TblEmailConfigurations{},
		TblGeneralSettings{},
		TblMemberSettings{},
		TblGraphqlSettings{},
		TblTimezones{},
		TblForms{},
		TblFormResponse{},
		// TblFormFieldOptions{},
		// TblFormFields{},
		// TblFormMasterFields{},
		// TblFormRegistrations{},
		// TblFormValues{},
		TblPageTypes{},
		TblTemplates{},
		TblBlocks{},
		TblBlockCollections{},
		TblBlockMstrTags{},
		TblBlockTags{},
		TblMstrTenant{},
		TblTemplateModules{},
		TblWebhooks{},
		TblApps{},
		TblAiPrompt{},
	)

	if err != nil {

		panic(err)

	}

}
