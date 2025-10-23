package models

import (
	"time"
)


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
	TenantId    int       `gorm:"type:integer"`
}


func GetCountryLIst()([]TblCountrie,error){

	var CountryList []TblCountrie
	if err := DB.Table("tbl_countries").Find(&CountryList).Error; err != nil {

		return []TblCountrie{},err
	}

	return CountryList,nil
}
