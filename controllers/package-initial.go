package controllers

import (
	"fmt"
	"os"
	"sync"

	newauth "github.com/spurtcms/auth"
	cat "github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	memaccess "github.com/spurtcms/member-access"
	"github.com/spurtcms/team"
	role "github.com/spurtcms/team-roles"
)

var (
	NewAuth        *newauth.Auth
	NewTeam        *team.Teams
	NewRole        *role.PermissionConfig
	ChannelConfig  *chn.Channel
	CategoryConfig *cat.Categories
	MemberConfig   *mem.Member
)

var lock = &sync.Mutex{}

// AuthCofing
func AuthConfig() *newauth.Auth {
	if NewAuth == nil {
		lock.Lock()
		defer lock.Unlock()
		if NewAuth == nil {
			fmt.Println("Creating Auth instance now.")
			NewAuth = newauth.AuthSetup(newauth.Config{
				SecretKey: os.Getenv("JWT_SECRET"),
				DB:        DB,
			})
		} else {
			fmt.Println("Single Auth already created.")
		}
	} else {
		fmt.Println("Single Auth already created.")
	}

	return NewAuth
}

// TeamConfig
func GetTeamInstance() *team.Teams {
	if NewTeam == nil {
		lock.Lock()
		defer lock.Unlock()
		if NewTeam == nil {
			fmt.Println("Creating Teams instance now.")
			NewTeam = team.TeamSetup(team.Config{
				DB:               DB,
				AuthEnable:       true,
				PermissionEnable: true,
				Auth:             NewAuth,
			})
		} else {
			fmt.Println("Teams instance already created.")
		}
	} else {
		fmt.Println("Teams instance already created.")
	}
	return NewTeam
}

// roles and permissions config
func GetRoleInstance() *role.PermissionConfig {
	if NewRole == nil {
		lock.Lock()
		defer lock.Unlock()
		if NewRole == nil {
			fmt.Println("Creating Roles instance now.")
			NewRole = role.RoleSetup(role.Config{
				DB:               DB,
				AuthEnable:       true,
				PermissionEnable: true,
				Authenticate:     NewAuth,
			})
		} else {
			fmt.Println("Roles instance already created.")
		}
	} else {
		fmt.Println("Roles instance already created.")
	}
	return NewRole
}

// channel config
func GetChannelInstance() *chn.Channel {
	AuthConfig()
	if ChannelConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if ChannelConfig == nil {
			fmt.Println("Creating Channel instance now.")
			ChannelConfig = chn.ChannelSetup(chn.Config{
				DB:               DB,
				AuthEnable:       true,
				PermissionEnable: true,
				Auth:             NewAuth,
			})
		} else {
			fmt.Println("Channel instance already created.")
		}
	} else {
		fmt.Println("Channel instance already created.")
	}
	return ChannelConfig
}

// Categoryconfig
func GetCategoryInstance() *cat.Categories {
	AuthConfig()
	if CategoryConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if CategoryConfig == nil {
			fmt.Println("Creating category instance now.")
			CategoryConfig = cat.CategoriesSetup(cat.Config{
				DB:               db,
				AuthEnable:       true,
				PermissionEnable: true,
				Auth:             NewAuth,
			})
		} else {
			fmt.Println("category instance already created.")
		}
	} else {
		fmt.Println("category instance already created.")
	}

	return CategoryConfig
}

// member config
func GetMemberInstance() *mem.Member {
	AuthConfig()
	if MemberConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if MemberConfig == nil {
			fmt.Println("Creating Member instance now.")
			MemberConfig = mem.MemberSetup(mem.Config{
				DB:               db,
				AuthEnable:       true,
				PermissionEnable: true,
				Auth:             NewAuth,
			})
		} else {
			fmt.Println("Member instance already created.")
		}
	} else {
		fmt.Println("Member instance already created.")
	}
	return MemberConfig
}

var MemberaccessConfig *memaccess.AccessControl

/*content access Config*/
func GetMemberaccessInstance() *memaccess.AccessControl {
	AuthConfig()
	if MemberaccessConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if MemberaccessConfig == nil {
			fmt.Println("Creating Member instance now.")
			MemberaccessConfig = memaccess.AccessSetup(memaccess.Config{
				DB:               db,
				AuthEnable:       true,
				PermissionEnable: true,
				Auth:             NewAuth,
			})
		} else {
			fmt.Println("Member instance already created.")
		}
	} else {
		fmt.Println("Member instance already created.")
	}
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
}
