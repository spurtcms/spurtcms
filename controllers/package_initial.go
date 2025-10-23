package controllers

import (
	"os"

	newauth "github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
	cat "github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	courses "github.com/spurtcms/courses"
	forms "github.com/spurtcms/forms-builders"
	listing "github.com/spurtcms/listing"
	mem "github.com/spurtcms/member"
	memaccess "github.com/spurtcms/member-access"
	memship "github.com/spurtcms/membership"
	"github.com/spurtcms/menu"
	menus "github.com/spurtcms/menu"
	"github.com/spurtcms/team"
	role "github.com/spurtcms/team-roles"
)

var (
	NewAuth            *newauth.Auth
	NewTeam            *team.Teams
	NewTeamWP          *team.Teams
	NewRole            *role.PermissionConfig
	NewRoleWP          *role.PermissionConfig
	ChannelConfig      *chn.Channel
	ChannelConfigWP    *chn.Channel
	CategoryConfig     *cat.Categories
	MemberConfig       *mem.Member
	MemberConfigWP     *mem.Member
	MemberaccessConfig *memaccess.AccessControl
	BlockConfig        *blocks.Block
	FormConfig         *forms.Formbuilders
	MembershipConfig   *memship.Membership
	CoursesConfig      *courses.Courses
	MenuConfig         *menus.Menu
	MenuConfigWp       *menus.Menu
	ListingConfig      *listing.Listing
)

// AuthCofing
func AuthConfig() *newauth.Auth {

	NewAuth = newauth.AuthSetup(newauth.Config{
		SecretKey: os.Getenv("JWT_SECRET"),
		DB:        DB,
		// DataBaseType: os.Getenv("DATABASE_TYPE"),
	})

	return NewAuth
}

// TeamConfig
func GetTeamInstance() *team.Teams {

	NewTeam = team.TeamSetup(team.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewTeam
}

// TeamConfig
func GetTeamInstanceWithoutPermission() *team.Teams {

	NewTeamWP = team.TeamSetup(team.Config{
		DB:               DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewTeamWP
}

// roles and permissions config
func GetRoleInstance() *role.PermissionConfig {

	NewAuth.AuthFlg = true
	NewAuth.PermissionFlg = true

	NewRole = role.RoleSetup(role.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewRole
}

// roles and permissions without permission config
func GetRoleInstanceWithoutPermission() *role.PermissionConfig {

	NewRoleWP = role.RoleSetup(role.Config{
		DB:               DB,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return NewRoleWP
}

// channel config
func GetChannelInstance() *chn.Channel {

	ChannelConfig = chn.ChannelSetup(chn.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return ChannelConfig
}

// channel config without permission
func GetChannelInstanceWithoutPermission() *chn.Channel {

	ChannelConfigWP = chn.ChannelSetup(chn.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return ChannelConfigWP
}

// Categoryconfig
func GetCategoryInstance() *cat.Categories {

	CategoryConfig = cat.CategoriesSetup(cat.Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return CategoryConfig
}

// member config
func GetMemberInstance() *mem.Member {

	MemberConfig = mem.MemberSetup(mem.Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return MemberConfig
}

// member config without permission
func GetMemberInstanceWithoutPermission() *mem.Member {

	MemberConfigWP = mem.MemberSetup(mem.Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return MemberConfigWP
}

/*content access Config*/
func GetMemberaccessInstance() *memaccess.AccessControl {

	MemberaccessConfig = memaccess.AccessSetup(memaccess.Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
		// DataBaseType:     os.Getenv("DATABASE_TYPE"),
	})

	return MemberaccessConfig
}

func GetBlockInstance() *blocks.Block {

	BlockConfig = blocks.BlockSetup(blocks.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return BlockConfig
}

func GetFormInstance() *forms.Formbuilders {

	FormConfig = forms.FormSetup(forms.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return FormConfig
}

func GetMembershipInstance() *memship.Membership {

	MembershipConfig = memship.MembershipSetup(memship.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return MembershipConfig
}

func GetCoursesInstance() *courses.Courses {

	CoursesConfig = courses.CoursesSetup(courses.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return CoursesConfig
}

func GetListingInstance() *listing.Listing {

	ListingConfig = listing.ListingSetup(listing.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return ListingConfig
}

func GetMenuInstrance() *menu.Menu {

	MenuConfig = menu.MenuSetup(menu.Config{

		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             NewAuth,
	})

	return MenuConfig
}
func GetMenuInstranceWithoutPermission() *menu.Menu {

	MenuConfigWp = menu.MenuSetup(menu.Config{

		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	return MenuConfigWp
}
func PackageInitialize() {
	AuthConfig()
	GetTeamInstance()
	GetRoleInstance()
	GetChannelInstance()
	GetCategoryInstance()
	GetMemberInstance()
	GetMemberaccessInstance()
	GetTeamInstanceWithoutPermission()
	GetRoleInstanceWithoutPermission()
	GetMemberInstanceWithoutPermission()
	GetChannelInstanceWithoutPermission()
	GetCoursesInstance()
	GetBlockInstance()
	GetFormInstance()
	GetMenuInstrance()
	GetMenuInstranceWithoutPermission()
	GetListingInstance()
	GetMembershipInstance()
}
