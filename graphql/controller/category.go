package controller

import (
	"context"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryList(ctx context.Context, categoryFilter *model.CategoryFilter, commonFilter *model.Filter) (*model.CategoryDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.CategoryDetails{}, info.ErrGinCtx
	}

	var (
		limit, offset, categoryGroupId, hierarchyLevel int
		checkEntriesPresence, excludeGroup, excludeParent, exactLevelOnly bool
		categoryGroupSlug,channelSlug   string
		finalCategoriesList []model.Category
	)

	if commonFilter != nil {

		if commonFilter.Limit.IsSet() {
			limit = *commonFilter.Limit.Value()

		}
		if commonFilter.Offset.IsSet() {
			offset = *commonFilter.Offset.Value()

		}
	}

	if categoryFilter != nil {

		if categoryFilter.CategoryGroupID.IsSet() {
			categoryGroupId = *categoryFilter.CategoryGroupID.Value()
		}

		if categoryFilter.HierarchyLevel.IsSet() {
			hierarchyLevel = *categoryFilter.HierarchyLevel.Value()
		}

		if categoryFilter.CheckEntriesPresence.IsSet() && *categoryFilter.CheckEntriesPresence.Value(){
			checkEntriesPresence = *categoryFilter.CheckEntriesPresence.Value()
		}

		if categoryFilter.ExcludeGroup.IsSet() && *categoryFilter.ExcludeGroup.Value() {
			excludeGroup = *categoryFilter.ExcludeGroup.Value()
		}

		if categoryFilter.ExcludeParent.IsSet() && *categoryFilter.ExcludeParent.Value() {
			excludeParent = *categoryFilter.ExcludeParent.Value()
		}

		if categoryFilter.CategoryGroupSlug.IsSet() {
			categoryGroupSlug = *categoryFilter.CategoryGroupSlug.Value()
		}

		if categoryFilter.ExactLevelOnly.IsSet() && *categoryFilter.ExactLevelOnly.Value(){
			exactLevelOnly = *categoryFilter.ExactLevelOnly.Value()
		}

		if categoryFilter.ChannelSlug.IsSet() && *categoryFilter.ChannelSlug.Value() != ""{

			channelSlug = *categoryFilter.ChannelSlug.Value()
		}

	}

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.CategoryDetails{}, info.ErrFetchTenantDetails
	}

	categoryListLocal, count, err := CategoryInstance.CategoryList(limit, offset, categoryGroupId, hierarchyLevel,tenantDetails.TenantId, checkEntriesPresence, excludeGroup, excludeParent,exactLevelOnly,channelSlug,categoryGroupSlug)
	
	if err != nil {

		if err == gorm.ErrRecordNotFound {

			return &model.CategoryDetails{}, nil

		}

		ErrorLog.Printf("%v", info.ErrFetchCategoryDetails)

		c.AbortWithStatus(500)

		return &model.CategoryDetails{}, info.ErrFetchCategoryDetails
	}

	for _, categoryList := range categoryListLocal {

		cateModifiedOn := categoryList.ModifiedOn

		cateModifiedBy := categoryList.ModifiedBy

		var localCategoryList model.Category

		localCategoryList.CategoryName = categoryList.CategoryName
		localCategoryList.CategorySlug = categoryList.CategorySlug
		localCategoryList.CreatedBy = categoryList.CreatedBy
		localCategoryList.CreatedOn = categoryList.CreatedOn
		localCategoryList.Description = categoryList.Description
		localCategoryList.ID = categoryList.Id
		localCategoryList.ImagePath = categoryList.ImagePath
		localCategoryList.ModifiedBy = &cateModifiedBy
		localCategoryList.ModifiedOn = &cateModifiedOn
		localCategoryList.ParentID = categoryList.ParentId
		localCategoryList.TenantID = categoryList.TenantId

		finalCategoriesList = append(finalCategoriesList, localCategoryList)
	}

	return &model.CategoryDetails{Categorylist: finalCategoriesList, Count: count}, nil

}
