package controllers

import (
	"os"

	newauth "github.com/spurtcms/auth"
	cat "github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	memaccess "github.com/spurtcms/member-access"
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
}
