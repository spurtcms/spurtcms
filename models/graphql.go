package models

import (
	"time"

	"gorm.io/gorm"
)

type TblGraphqlSettings struct {
	Id          int
	TokenName   string
	Description string
	Duration    string
	CreatedBy   int
	CreatedOn   time.Time
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	Token       string    
	IsDefault   int      `gorm:"DEFAULT:0"`
	DateString  string `gorm:"-"`
	ExpiryTime  time.Time
	TenantId    int
}

func GetListOfTokens(limit int, offset int, keyword string, tenantid int) (grpahqlsett []TblGraphqlSettings, count int64, err error) {

	query := DB.Debug().Table("tbl_graphql_settings").Where("is_deleted = 0 and tenant_id = ?", tenantid)

	var listQuery *gorm.DB
	var countQuery *gorm.DB

	if keyword != "" {
		query = query.Where("lower(trim(token_name)) like lower(trim(?)) or lower(trim(description)) like lower(trim(?)) or lower(trim(duration)) like lower(trim(?))", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

		listQuery = query.Limit(limit).Offset(offset).Order("id desc").Find(&grpahqlsett)
		if listQuery.Error != nil {

			return grpahqlsett, count, listQuery.Error
		}

		countQuery = DB.Debug().Table("tbl_graphql_settings").Where("is_deleted = 0 and tenant_id = ?", tenantid).Where("lower(trim(token_name)) like lower(trim(?)) or lower(trim(description)) like lower(trim(?)) or lower(trim(duration)) like lower(trim(?))", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Count(&count)
		if countQuery.Error != nil {
			return []TblGraphqlSettings{}, -1, countQuery.Error
		}

	} else {

		listQuery = query.Limit(limit).Offset(offset).Order("is_default desc,id desc").Find(&grpahqlsett)
		if listQuery.Error != nil {

			return grpahqlsett, count, listQuery.Error
		}

		countQuery = DB.Table("tbl_graphql_settings").Where("is_deleted = 0 and tenant_id = ?", tenantid).Count(&count)
		if countQuery.Error != nil {
			return []TblGraphqlSettings{}, -1, countQuery.Error
		}

	}

	return grpahqlsett, count, nil
}

func CreateApiToken(graphqlapi TblGraphqlSettings) error {

	query := DB.Debug().Table("tbl_graphql_settings")

	if graphqlapi.ExpiryTime.IsZero() {

		query.Omit("expiry_time")
	}

	query.Create(&graphqlapi)

	if err := query.Error; err != nil {

		return err
	}

	return nil

}

func GetTokenDetailsById(id int, tenantid int) (graphql TblGraphqlSettings, err error) {

	if err := DB.Table("tbl_graphql_settings").Where("id=? and tenant_id = ?", id, tenantid).First(&graphql).Error; err != nil {

		return graphql, err
	}

	return graphql, nil
}

func UpdateApiToken(graphqlapi map[string]interface{}, id int, tenantid int) error {

	query := DB.Debug().Table("tbl_graphql_settings").Where("id=? and tenant_id = ?", id, tenantid).UpdateColumns(graphqlapi)

	if err := query.Error; err != nil {

		return err
	}

	return nil
}

func DeleteApiToken(id int, deletedby int, deletedon time.Time, tenantid int) error {

	if err := DB.Debug().Table("tbl_graphql_settings").Where("id=? and tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

func MultiDeleteGraphqltoken(graph *TblGraphqlSettings, ids []int, tenantid int) error {

	if err := DB.Table("tbl_graphql_settings").Where("id in (?) and tenant_id = ?", ids, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": graph.IsDeleted, "deleted_by": graph.DeletedBy, "deleted_on": graph.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}
