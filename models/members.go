package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

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

func GetMemberByEmail(email string, tenantid string) (member TblMembers, err error) {

	query := DB.Table("tbl_members").Where("is_deleted = 0 AND email = ?", email)

	if tenantid != "" {
		query = query.Where("tenant_id = ?", tenantid)
	}

	if err := query.First(&member).Error; err != nil {
		return TblMembers{}, err
	}

	return member, nil
}

func GetMemberById(id int, tenantid string) (member TblMembers, err error) {

	query := DB.Table("tbl_members").Where("is_deleted = 0 AND id = ?", id)

	if tenantid != "" {
		query = query.Where("tenant_id = ?", tenantid)
	}

	if err := query.First(&member).Error; err != nil {
		return TblMembers{}, err
	}

	return member, nil
}

func CreateMember(members *TblMembers) (member TblMembers, err error) {

	if err := DB.Table("tbl_members").Create(&members).Error; err != nil {

		return TblMembers{}, err

	}

	return *members, nil
}

func UpdateMemberOtp(member TblMembers, tenantid string) error {

	result := DB.Table("tbl_members").Where("id = ? and  tenant_id=?", member.Id, member.TenantId).UpdateColumns(map[string]interface{}{"modified_on": member.ModifiedOn, "modified_by": member.Id, "otp": member.Otp, "otp_expiry": member.OtpExpiry})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CheckMemberLogin(Email string, Password string, userid int, tenantid string) (string, bool, bool, error) {

	var members TblMembers

	err := DB.Table("tbl_members").Where("tbl_members.is_deleted = 0 AND tbl_members.email = ? AND tbl_members.created_by = ? AND tbl_members.tenant_id = ?", Email, userid, tenantid).First(&members).Error

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

	token, err := CreateToken(members)

	if err != nil {

		return "", false, true, err
	}

	return token, true, true, nil

}

func CreateToken(members TblMembers) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["member_id"] = members.Id
	atClaims["user_id"] = members.CreatedBy
	atClaims["tenant_id"] = members.TenantId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
