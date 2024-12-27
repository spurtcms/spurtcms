package controller

import (
	"context"
	"errors"
	"net/http"
	"spurt-cms/controllers"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/member"
	"gorm.io/gorm"
)

func MemberRegister(ctx context.Context, memberData *model.MemberDetails, arguments *model.MemberArguments) (bool, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		c.AbortWithStatus(500)

		return false, info.ErrGinCtx
	}

	tenantDetails := c.GetStringMap("tenant")
	tenantId := tenantDetails["TenantId"].(int)

	var (
		memberSettings member.TblMemberSetting
		err            error
	)

	memberSettings, err = MemberInstance.GetMemberSettings(tenantId)
	if err != nil {

		return false, err
	}

	if memberSettings.AllowRegistration == 0 {

		return false, info.ErrMemberRegisterPerm
	}

	var memberDetails member.TblMember

	if memberData.Mobile.IsSet() {

		memberDetails.MobileNo = *memberData.Mobile.Value()
	}

	if memberData.LastName.IsSet() {

		memberDetails.LastName = *memberData.LastName.Value()
	}

	if memberData.Password.IsSet() {

		hashPass := controllers.HashingPassword(*memberData.Password.Value())

		memberDetails.Password = hashPass
	}

	if memberData.Username.IsSet() {

		memberDetails.Username = *memberData.Username.Value()

		isExist, err := MemberInstance.CheckNameInMember(0, *memberData.Username.Value(), tenantId)
		if isExist || err == nil {
			err = errors.New("member already exists")

			c.AbortWithError(422, err)

			return isExist, err
		}
	}

	if memberData.Email != "" {

		memberDetails.Email = memberData.Email

		isExist, err := MemberInstance.CheckEmailInMember(0, memberData.Email, tenantId)

		if isExist || err == nil {

			err = errors.New("member already exists")

			c.AbortWithError(http.StatusBadRequest, err)

			return isExist, err
		}
	}

	memberDetails.FirstName = memberData.FirstName
	memberDetails.Username = strings.ToLower(memberData.FirstName)
	memberDetails.IsActive = 1

	// MemberInstance.CreateMember()

	return true, nil
}

func MembersList(ctx context.Context, filter *model.Filter) (*model.MembersDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.MembersDetails{}, info.ErrGinCtx
	}

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.MembersDetails{}, info.ErrFetchTenantDetails
	}

	var (
		limit, offset int
		keyword       string
	)

	if filter != nil {

		if filter.Limit.IsSet() && filter.Limit.Value() != nil {

			limit = *filter.Limit.Value()
		}

		if filter.Offset.IsSet() && filter.Offset.Value() != nil {

			offset = *filter.Offset.Value()
		}

		if filter.Keyword.IsSet() && filter.Keyword.Value() != nil {

			keyword = *filter.Keyword.Value()
		}

	}

	input := model.MembersListReq{
		Limit:    limit,
		Offset:   offset,
		Keyword:  keyword,
		TenantId: tenantDetails.TenantId,
		Count:    true,
	}

	members, count, err := model.Model.MembersList(input)

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			c.AbortWithStatus(400)

			return &model.MembersDetails{}, err
		}

		ErrorLog.Printf("%v", err)

		c.AbortWithStatus(500)

		return &model.MembersDetails{}, err
	}

	var memberlist []model.Members

	for _, member := range members {

		modifiedOn := member.ModifiedOn

		modifiedBy := member.ModifiedBy

		member := model.Members{
			ID:         member.ID,
			FirstName:  member.FirstName,
			LastName:   member.LastName,
			Email:      member.Email,
			Mobile:     member.Mobile,
			IsActive:   member.IsActive,
			CreatedOn:  member.CreatedOn,
			CreatedBy:  member.CreatedBy,
			IsDeleted:  member.IsDeleted,
			ModifiedOn: modifiedOn,
			ModifiedBy: modifiedBy,
			TenantID:   member.TenantID,
		}

		memberlist = append(memberlist, member)
	}

	return &model.MembersDetails{MembersList: memberlist, Count: int(count)}, nil
}
