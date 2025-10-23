package model

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type MembersListReq struct {
	Limit    int
	Offset   int
	Keyword  string
	TenantId string
	Count    bool
}

type TblMemberGroup struct {
	Id          int
	Name        string
	Slug        string
	Description string
	IsActive    int
	IsDeleted   int
	CreatedOn   time.Time
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"default:null"`
	ModifiedBy  int       `gorm:"default:null"`
	DeletedOn   time.Time `gorm:"default:null"`
	DeletedBy   int       `gorm:"default:null"`
	TenantId    string
}

func (model ModelConfig) MembersList(inputs MembersListReq) (members []Members, count int64, err error) {

	fmt.Println(model.DB, inputs.Count, inputs.Limit, "dbvalue")
	query := model.DB.Debug().Table("tbl_members").Where("tbl_members.is_deleted = 0").Order("id desc")

	if inputs.TenantId != "" {

		query = query.Where("tbl_members.tenant_id=?", inputs.TenantId)

	}

	if inputs.Keyword != "" {

		query = query.Where("LOWER(TRIM(first_name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%")
	}

	if inputs.Count {

		err = query.Count(&count).Error

		if err != nil {

			return []Members{}, 0, err
		}
	}

	if inputs.Limit != 0 {

		query = query.Limit(inputs.Limit)
	}

	if inputs.Offset != -1 {

		query = query.Offset(inputs.Offset)
	}

	err = query.Find(&members).Error

	if err != nil {

		return []Members{}, 0, err
	}

	fmt.Println(members, "mmembersss")
	return members, count, nil
}
func (model ModelConfig) CheckMemberLogin(Email string, Password string, userid int, tenantid string) (string, bool, bool, error) {

	var members Members

	err := model.DB.Table("tbl_members").Where("tbl_members.is_deleted = 0 AND tbl_members.email = ? AND tbl_members.created_by = ? AND tbl_members.tenant_id = ?", Email, userid, tenantid).First(&members).Error

	if err != nil {
		return "", false, false, err
	}

	if members.Email == "" {
		return "", false, false, err
	}

	passerr := bcrypt.CompareHashAndPassword([]byte(members.Password), []byte(Password))

	if passerr != nil || passerr == bcrypt.ErrMismatchedHashAndPassword {

		return "", false, true, passerr
	}

	token, err := model.CreateToken(members)

	if err != nil {

		return "", false, true, err
	}

	return token, true, true, nil

}

func (model ModelConfig) CreateToken(members Members) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["member_id"] = members.ID
	atClaims["user_id"] = members.CreatedBy
	atClaims["tenant_id"] = members.TenantID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (model ModelConfig) GetmemberGroup(slug string, tenantid string) (membergrb TblMemberGroup,err  error) {

	err1 := model.DB.Table("tbl_member_groups").Where("tbl_member_groups.is_deleted = 0 and slug=? and tbl_member_groups.tenant_id = ?",slug, tenantid).First(&membergrb).Error

	if err1 != nil {
		return TblMemberGroup{},nil
	}
	return TblMemberGroup{}, nil
}


