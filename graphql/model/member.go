package model

import "fmt"

type MembersListReq struct {
	Limit    int
	Offset   int
	Keyword  string
	TenantId int
	Count    bool
}

func (model ModelConfig) MembersList(inputs MembersListReq) (members []Members, count int64, err error) {

	fmt.Println(model.DB, inputs.Count, inputs.Limit, "dbvalue")
	query := model.DB.Debug().Table("tbl_members").Where("tbl_members.is_deleted = 0").Order("id desc")

	if inputs.TenantId != -1 {

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
