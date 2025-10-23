package postgres

import (
	"spurt-cms/controllers"
	"time"

	"gorm.io/datatypes"
)

type TblRoles struct {
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
	TenantId    string    `gorm:"type:character varying"`
}

type TblUsers struct {
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
	StorageType       string    `gorm:"type:character varying"`
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
	TenantId          string    `gorm:"type:character varying"`
	Otp               int       `gorm:"DEFAULT:NULL"`
	OtpExpiry         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	S3FolderName      string    `gorm:"type:character varying"`
	ArticleCount      int       `gorm:"type:integer"`
	TotalCount        int       `gorm:"type:integer"`
	GoTemplateDefault int       `gorm:"type:integer"`
	Subdomain         string    `gorm:"type:character varying"`
	UsageMode         string    `gorm:"type:character varying"`
	ChannelId         int       `gorm:"type:integer"`
	Country           string    `gorm:"type:character varying"`
	ChooseType        string    `gorm:"type:character varying"`
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
	TenantId     string    `gorm:"type:character varying"`
}

type TblChannels struct {
	Id                 int       `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelName        string    `gorm:"type:character varying"`
	ChannelUniqueId    string    `gorm:"type:character varying"`
	ChannelDescription string    `gorm:"type:character varying"`
	SlugName           string    `gorm:"type:character varying"`
	FieldGroupId       int       `gorm:"type:integer"`
	IsActive           int       `gorm:"type:integer"`
	IsDeleted          int       `gorm:"type:integer"`
	CreatedOn          time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy          int       `gorm:"type:integer"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	ChannelType        string    `gorm:"type:character varying"`
	CollectionCount    int       `gorm:"type:integer"`
	CloneCount         int       `gorm:"type:integer"`
	TenantId           string    `gorm:"type:character varying"`
	ImagePath          string    `gorm:"type:character varying"`
}

type TblMemberGroups struct {
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
	TenantId    string    `gorm:"type:character varying"`
}

type TblMembers struct {
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
	StorageType      string    `gorm:"type:character varying"`
	TenantId         string    `gorm:"type:character varying"`
}

type TblAccessControls struct {
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
	TenantId          string    `gorm:"type:character varying"`
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
	TenantId                 string    `gorm:"type:character varying"`
}

type TblAccessControlUserGroups struct {
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
	TenantId        string    `gorm:"type:character varying"`
}

type TblEmailTemplates struct {
	Id                  int       `gorm:"primaryKey;auto_increment;type:serial"`
	TemplateName        string    `gorm:"type:character varying"`
	TemplateSlug        string    `gorm:"type:character varying"`
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
	TenantId            string    `gorm:"type:character varying"`
}

type TblFields struct {
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
	TenantId         string    `gorm:"type:character varying"`
}

type TblFieldGroups struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	GroupName  string    `gorm:"type:character varying"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblFieldOptions struct {
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
	TenantId    string    `gorm:"type:character varying"`
	OrderIndex  int       `gorm:"DEFAULT:NULL"`
}

type TblFieldTypes struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	TypeName   string    `gorm:"type:character varying"`
	TypeSlug   string    `gorm:"type:character varying"`
	IsActive   int       `gorm:"type:integer"`
	IsDeleted  int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblLanguages struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	LanguageName string    `gorm:"type:character varying"`
	LanguageCode string    `gorm:"type:character varying"`
	ImagePath    string    `gorm:"type:character varying"`
	StorageType  string    `gorm:"type:character varying"`
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
	TenantId     string    `gorm:"type:character varying"`
}

type TblModules struct {
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
	TenantId             string    `gorm:"type:character varying"`
}

type TblModulePermissions struct {
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
	TenantId             string    `gorm:"type:character varying"`
	ChannelId            int       `gorm:"column:channel_id"`
}

type TblRolePermissions struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	TenantId     string    `gorm:"type:character varying"`
}

type TblChannelCategories struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId  int       `gorm:"type:integer"`
	CategoryId string    `gorm:"type:character varying"`
	CreatedAt  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblGroupFields struct {
	Id        int    `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId int    `gorm:"type:integer"`
	FieldId   int    `gorm:"type:integer"`
	TenantId  string `gorm:"type:character varying"`
}

type TblUserPersonalizes struct {
	Id                  int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId              int       `gorm:"type:integer"`
	MenuBackgroundColor string    `gorm:"type:character varying"`
	FontColor           string    `gorm:"type:character varying"`
	IconColor           string    `gorm:"type:character varying"`
	LogoPath            string    `gorm:"type:character varying"`
	ExpandLogoPath      string    `gorm:"type:character varying"`
	CreatedOn           time.Time `gorm:"type:timestamp without time zone"`
	ModifiedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId            string    `gorm:"type:character varying"`
}

type TblRoleUsers struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId     int       `gorm:"type:integer"`
	UserId     int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblChannelEntries struct {
	Id                 int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title              string    `gorm:"type:character varying"`
	Slug               string    `gorm:"type:character varying"`
	Description        string    `gorm:"type:character varying"`
	UserId             int       `gorm:"type:integer"`
	ChannelId          int       `gorm:"type:integer"`
	Status             int       `gorm:"type:integer"` //0-draft 1-publish 2-unpublish
	CoverImage         string    `gorm:"type:character varying"`
	ThumbnailImage     string    `gorm:"type:character varying"`
	MetaTitle          string    `gorm:"type:character varying"`
	MetaDescription    string    `gorm:"type:character varying"`
	Keyword            string    `gorm:"type:character varying"`
	CategoriesId       string    `gorm:"type:character varying"`
	MemebrshipLevelIds string    `gorm:"type:character varying"`
	RelatedArticles    string    `gorm:"type:character varying"`
	Feature            int       `gorm:"DEFAULT:0"`
	ViewCount          int       `gorm:"DEFAULT:0"`
	CreateTime         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	PublishedTime      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ImageAltTag        string    `gorm:"type:character varying"`
	Author             string    `gorm:"type:character varying"`
	SortOrder          int       `gorm:"type:integer"`
	Excerpt            string    `gorm:"type:character varying"`
	ReadingTime        int       `gorm:"type:integer"`
	Tags               string    `gorm:"type:character varying"`
	CreatedOn          time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy          int       `gorm:"type:integer"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive           int       `gorm:"type:integer"`
	IsDeleted          int       `gorm:"DEFAULT:0"`
	DeletedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy          int       `gorm:"DEFAULT:NULL"`
	TenantId           string    `gorm:"type:character varying"`
	Uuid               string    `gorm:"type:character varying"`
	ParentId           int       `gorm:"type:integer"`
	OrderIndex         int       `gorm:"type:integer"`
	MembergroupId      string    `gorm:"type:character varying"`
	CtaId              int       `gorm:"type:integer"`
	LanguageId         int       `gorm:"type:integer"`
	EntryReferenceId   int       `gorm:"type:integer"`
}

type TblChannelEntryFields struct {
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
	TenantId       string    `gorm:"type:character varying"`
}

type TblMemberProfiles struct {
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
	StorageType     string            `gorm:"type:character varying"`
	ClaimDate       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId        string            `gorm:"type:character varying"`
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
	TenantId                string            `gorm:"type:character varying"`
}

type TblStorageTypes struct {
	Id           int               `gorm:"primaryKey;auto_increment;type:serial"`
	Local        string            `gorm:"type:character varying"`
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:character varying"`
	TenantId     string            `gorm:"type:character varying"`
}

type TblEmailConfigurations struct {
	Id           int               `gorm:"primaryKey;auto_increment;type:serial"`
	SmtpConfig   datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string            `gorm:"type:character varying"`
	TenantId     string            `gorm:"type:character varying"`
}

type TblGeneralSettings struct {
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
	TenantId       string    `gorm:"type:character varying"`
	StorageType    string    `gorm:"type:character varying"`
}
type TblMemberSettings struct {
	Id                int       `gorm:"primaryKey;auto_increment;type:serial"`
	AllowRegistration int       `gorm:"type:integer"`
	MemberLogin       string    `gorm:"type:character varying"`
	ModifiedBy        int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	NotificationUsers string    `gorm:"type:character varying"`
	TenantId          string    `gorm:"type:character varying"`
}

type TblGraphqlSettings struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	TokenName   string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Duration    string    `gorm:"type:character varying"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
	Token       string    `gorm:"type:character varying"`
	IsDefault   int       `gorm:"type:integer;DEFAULT:0"`
	ExpiryTime  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId    string    `gorm:"type:character varying"`
}

type TblTimezones struct {
	Id       int    `gorm:"primaryKey;auto_increment;type:serial"`
	Timezone string `gorm:"type:character varying"`
	TenantId string `gorm:"type:character varying"`
}

type TblForms struct {
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
	TenantId             string    `gorm:"type:character varying"`
	Uuid                 string    `gorm:"type:character varying"`
	FormImagePath        string    `gorm:"type:character varying"`
	FormDescription      string    `gorm:"type:character varying"`
	ChannelId            string    `gorm:"type:character varying"`
	ChannelName          string    `gorm:"type:character varying"`
	FormPreviewImagepath string    `gorm:"type:character varying"`
	FormPreviewImagename string    `gorm:"type:character varying"`
}

type TblFormResponse struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId       int       `gorm:"type:integer;"`
	FormResponse string    `gorm:"type:character varying"`
	UserId       int       `gorm:"type:integer;"`
	IsActive     int       `gorm:"type:integer"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId     string    `gorm:"type:character varying"`
	EntryId      int       `gorm:"type:integer;DEFAULT:0"`
}

type TblFormRegistrations struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name       string    `gorm:"type:character varying"`
	EmailId    string    `gorm:"type:character varying"`
	MobileNo   string    `gorm:"type:character varying"`
	FormId     int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblFormFields struct {
	Id            int       `gorm:"primaryKey;auto_increment;type:serial"`
	FormId        int       `gorm:"type:integer"`
	MasterFieldId int       `gorm:"type:integer"`
	FieldName     string    `gorm:"type:character varying"`
	Mandatory     int       `gorm:"type:integer"`
	OptionExists  int       `gorm:"type:integer"`
	CreatedBy     int       `gorm:"type:integer"`
	CreatedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted     int       `gorm:"type:integer;DEFAULT:0"`
	TenantId      string    `gorm:"type:character varying"`
}

type TblFormValues struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldId    int       `gorm:"type:integer"`
	FieldValue string    `gorm:"type:character varying"`
	RegisterId int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblFormMasterFields struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName        string    `gorm:"type:character varying"`
	FieldDescription string    `gorm:"type:character varying"`
	IconPath         string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	CreatedBy        int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	TenantId         string    `gorm:"type:character varying"`
}

type TblFormFieldOptions struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldId    int       `gorm:"type:integer"`
	Options    string    `gorm:"type:character varying"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	TenantId   string    `gorm:"type:character varying"`
}
type TblPageTypes struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title      string    `gorm:"type:character varying"`
	Slug       string    `gorm:"type:character varying"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
	TenantId   string    `gorm:"type:character varying"`
}

type TblTemplates struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	TemplateName     string    `gorm:"type:character varying"`
	TemplateSlug     string    `gorm:"type:character varying"`
	TemplateModuleId int       `gorm:"type:integer"`
	ImageName        string    `gorm:"type:character varying"`
	ImagePath        string    `gorm:"type:character varying"`
	DeployLink       string    `gorm:"type:character varying"`
	GithubLink       string    `gorm:"type:character varying"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	IsActive         int       `gorm:"type:integer"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer"`
	TenantId         string    `gorm:"type:character varying"`
	Preview          string    `gorm:"type:character varying"`
	SlugName         string    `gorm:"type:character varying"`
}
type TblBlocks struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	SlugName         string    `gorm:"type:character varying"`
	ChannelSlugname  string    `gorm:"type:character varying"`
	BlockDescription string    `gorm:"type:text"`
	BlockContent     string    `gorm:"type:text"`
	BlockCss         string    `gorm:"type:text"`
	IconImage        string    `gorm:"type:character varying"`
	CoverImage       string    `gorm:"type:character varying"`
	Prime            int       `gorm:"type:integer"`
	IsActive         int       `gorm:"type:integer"`
	TenantId         string    `gorm:"type:character varying"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedBy       int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	ChannelId        int       `gorm:"type:character varying"`
}

type TblBlockTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	BlockId   int       `gorm:"type:integer"`
	TagId     int       `gorm:"type:integer"`
	TagName   string    `gorm:"type:character varying"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}

type TblBlockMstrTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	TagTitle  string    `gorm:"type:character varying"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
}

type TblBlockCollections struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId    int       `gorm:"type:integer"`
	BlockId   int       `gorm:"type:integer"`
	TenantId  string    `gorm:"type:character varying"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}

type TblMstrTenant struct {
	Id            int       `gorm:"primaryKey;auto_increment;type:serial"`
	TenantId      string    `gorm:"type:character varying"`
	S3StoragePath string    `gorm:"type:character varying"`
	IsDeleted     int       `gorm:"type:integer;DEFAULT:0"`
	DeletedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
}

type TblTemplateModules struct {
	Id                 int       `gorm:"primaryKey;auto_increment;type:serial"`
	TemplateModuleName string    `gorm:"type:character varying"`
	TemplateModuleSlug string    `gorm:"type:character varying"`
	Description        string    `gorm:"type:character varying"`
	IsActive           int       `gorm:"type:integer;DEFAULT:1"`
	CreatedBy          int       `gorm:"type:integer"`
	CreatedOn          time.Time `gorm:"type:timestamp without time zone"`
	ChannelId          int       `gorm:"type:integer"`
	TenantId           string    `gorm:"type:character varying"`
	IsDeleted          int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	SlugName           string    `gorm:"type:character varying"`
}
type TblApps struct {
	Id            int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title         string    `gorm:"type:character varying"`
	Description   string    `gorm:"type:character varying"`
	FieldJsonPath string    `gorm:"type:character varying"`
	IconName      string    `gorm:"type:character varying"`
	IconPath      string    `gorm:"type:character varying"`
	CreatedBy     int       `gorm:"type:integer"`
	CreatedOn     time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsActive      int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted     int       `gorm:"type:integer;DEFAULT:0"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId      string    `gorm:"type:character varying"`
	ChannelSlug   string    `gorm:"type:character varying"`
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
	TenantId     string    `gorm:"type:character varying"`
}

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
	TenantId    string    `gorm:"type:character varying"`
}

type TblCountrie struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	CountryCode string    `gorm:"type:character varying"`
	CountryName string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	TenantId    string    `gorm:"type:character varying"`
}

type TblSavedEntries struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	EntryId   int       `gorm:"type:integer"`
	UserId    int       `gorm:"type:integer"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone"`
	IsDeleted int       `gorm:"type:integer"`
}

type TblMenus struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	SlugName    string    `gorm:"type:character varying"`
	TenantId    string    `gorm:"type:character varying"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	Status      int       ` gorm:"type:integer"`
	UrlPath     string    `gorm:"type:character varying"`
	ParentId    int       `gorm:"type:integer"`
	Type        string    `gorm:"type:character varying"`
	TypeId      int       `gorm:"type:integer"`
	WebsiteId   int       `gorm:"type:integer"`
	ListingsIds string    `gorm:"type:character varying"`
}

type TblGoTemplates struct {
	Id                  int       `gorm:"primaryKey;autoIncrement;type:serial"`
	TemplateName        string    `gorm:"type:character varying"`
	TemplateImage       string    `gorm:"type:character varying"`
	TemplateDescription string    `gorm:"type:character varying"`
	IsDeleted           int       `gorm:"type:integer"`
	TenantId            string    `gorm:"type:character varying"`
	CreatedOn           time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy           int       `gorm:"type:integer"`
	ChannelSlugName     string    `gorm:"type:character varying"`
	TemplateModuleName  string    `gorm:"type:character varying"`
}

type TblGoTemplateSeo struct {
	Id               int    `gorm:"primaryKey;auto_increment;type:serial"`
	PageTitle        string `gorm:"type:character varying"`
	PageDescription  string `gorm:"type:character varying"`
	PageKeyword      string `gorm:"type:character varying"`
	StoreTitle       string `gorm:"type:character varying"`
	StoreDescription string `gorm:"type:character varying"`
	StoreKeyword     string `gorm:"type:character varying"`
	SiteMapName      string `gorm:"type:character varying"`
	SiteMapPath      string `gorm:"type:character varying"`
	TenantId         string `gorm:"type:character varying"`
	WebsiteId        int    `gorm:"type:integer"`
}

type TblGoTemplateSettings struct {
	Id              int    `gorm:"primaryKey;auto_increment;type:serial"`
	SiteName        string `gorm:"type:character varying"`
	SiteLogo        string `gorm:"type:character varying"`
	SiteLogoPath    string `gorm:"type:character varying"`
	SiteFavIcon     string `gorm:"type:character varying"`
	SiteFavIconPath string `gorm:"type:character varying"`
	WebsiteUrl      string `gorm:"type:character varying"`
	TenantId        string `gorm:"type:character varying"`
	WebsiteId       int    `gorm:"type:integer"`
}

type TblTemplatePages struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name            string    `gorm:"type:character varying"`
	Slug            string    `gorm:"type:character varying"`
	PageDescription string    `gorm:"type:character varying"`
	TenantId        string    `gorm:"type:character varying"`
	IsDeleted       int       `gorm:"type:integer"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy       int       `gorm:"type:integer"`
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL;type:integer"`
	Status          int       `gorm:"type:integer"`
	MetaTitle       string    `gorm:"type:character varying"`
	MetaDescription string    `gorm:"type:character varying"`
	MetaKeywords    string    `gorm:"type:character varying"`
	MetaSlug        string    `gorm:"type:character varying"`
	WebsiteId       int       `gorm:"type:integer"`
}

type TblWebsite struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name         string    `gorm:"type:character varying"`
	ChannelNames string    `gorm:"type:character varying"`
	TemplateId   int       `gorm:"type:integer"`
	TenantId     string    `gorm:"type:character varying"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL;type:integer"`
	Status       int       `gorm:"type:integer"`
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
		TblPageTypes{},
		TblTemplates{},
		TblBlocks{},
		TblBlockCollections{},
		TblBlockMstrTags{},
		TblBlockTags{},
		TblMstrTenant{},
		TblTemplateModules{},
		TblApps{},
		TblAiPrompt{},
		TblAiSettingsModule{},
		TblCountrie{},
		TblSavedEntries{},
		TblMenus{},
		TblGoTemplates{},
		TblGoTemplateSeo{},
		TblGoTemplateSettings{},
		TblTemplatePages{},
		TblWebsite{},
	)

	if err != nil {

		panic(err)

	}

}
