package controller

import (
	"context"
	"fmt"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"
	"spurt-cms/graphql/scalars"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/channels"
	"gorm.io/gorm"
)

func ChannelList(ctx context.Context, filter *model.Filter, sort *model.Sort) (*model.ChannelDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.ChannelDetails{}, info.ErrGinCtx
	}

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.ChannelDetails{}, info.ErrFetchTenantDetails
	}

	var (
		limit, offset, order int
		keyword, sortBy      string
		isActive             bool
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

		if filter.IsActive.IsSet() && filter.IsActive.Value() != nil {

			isActive = *filter.IsActive.Value()
		}
	}

	if sort != nil {

		if sort.SortBy.IsSet() && sort.SortBy.Value() != nil {

			sortBy = *sort.SortBy.Value()
		}

		if sort.Order.IsSet() && sort.Order.Value() != nil {

			order = *sort.Order.Value()
		}
	}

	input := channels.Channels{
		Limit:        limit,
		Offset:       offset,
		Keyword:      keyword,
		IsActive:     isActive,
		SortBy:       sortBy,
		SortingOrder: order,
		TenantId:     tenantDetails.TenantId,
		Count:        true,
		// ChannelEntries: true,
		// AuthorDetail:   true,
		// EntriesCount:   true,
	}

	channels, count, err := ChannelConfigWP.ListChannel(input)

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			c.AbortWithStatus(400)

			return &model.ChannelDetails{}, err
		}

		ErrorLog.Printf("%v", err)

		c.AbortWithStatus(500)

		return &model.ChannelDetails{}, err
	}

	var channelList []model.Channel

	for _, channel := range channels {

		modifiedOn := channel.ModifiedOn

		modifiedBy := channel.ModifiedBy

		channel := model.Channel{
			ID:                 channel.Id,
			ChannelName:        channel.ChannelName,
			ChannelDescription: channel.ChannelDescription,
			SlugName:           channel.SlugName,
			FieldGroupID:       channel.FieldGroupId,
			IsActive:           channel.IsActive,
			CreatedOn:          channel.CreatedOn,
			CreatedBy:          channel.CreatedBy,
			IsDeleted:          channel.IsDeleted,
			ModifiedOn:         &modifiedOn,
			ModifiedBy:         &modifiedBy,
			TenantID:           channel.TenantId,
		}

		channelList = append(channelList, channel)
	}

	return &model.ChannelDetails{Channellist: channelList, Count: count}, nil

}

func ChannelDetails(ctx context.Context, channelId *int, channelSlug *string, status *bool) (*model.Channel, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.Channel{}, info.ErrGinCtx
	}

	var (
		id       int
		slug     string
		isActive bool
	)

	if channelId != nil {

		id = *channelId
	}

	if channelSlug != nil {

		slug = *channelSlug
	}

	if status != nil {

		isActive = *status
	}

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.Channel{}, info.ErrFetchTenantDetails
	}

	inputs := channels.Channels{
		Id:       id,
		Slug:     slug,
		TenantId: tenantDetails.TenantId,
		IsActive: isActive,
		// ChannelEntries: true,
		// AuthorDetail:   true,
	}

	channel, err := ChannelConfigWP.ChannelDetail(inputs)

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			c.AbortWithStatus(400)

			return &model.Channel{}, err
		}

		ErrorLog.Printf("%v", err)

		c.AbortWithStatus(500)

		return &model.Channel{}, err
	}

	channelDetails := model.Channel{
		ID:                 channel.Id,
		ChannelName:        channel.ChannelName,
		ChannelDescription: channel.ChannelDescription,
		SlugName:           channel.SlugName,
		FieldGroupID:       channel.FieldGroupId,
		IsActive:           channel.IsActive,
		CreatedOn:          channel.CreatedOn,
		CreatedBy:          channel.CreatedBy,
		IsDeleted:          channel.IsDeleted,
		ModifiedOn:         &channel.ModifiedOn,
		ModifiedBy:         &channel.ModifiedBy,
		TenantID:           channel.TenantId,
	}

	return &channelDetails, nil
}

func ChannelEntriesList(ctx context.Context, commonFilter *model.Filter, sort *model.Sort, EntryFilter *model.EntriesFilter, additionalData *model.EntriesAdditionalData) (*model.ChannelEntryDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.ChannelEntryDetails{}, info.ErrGinCtx
	}

	type EntryStatus string

	const (
		Draft EntryStatus = "Draft"

		Publish EntryStatus = "Publish"

		Unpublish EntryStatus = "Unpublish"
	)

	var (
		limit, offset                                                                   int = 0, -1
		order                                                                           int
		isActive                                                                        bool
		keyword, title, sortBy                                                          string
		channelId, categoryId                                                           int
		categorySlug                                                                    string
		status                                                                          string
		memberProfFlag, categoriesFlag, fieldsFlg, authorFlag, selectedCategoriesFilter bool
	)

	tenantDetails, err := GetTenantDetails(c)

	if err != nil {

		ErrorLog.Printf("%v", info.ErrFetchTenantDetails)

		c.AbortWithStatus(500)

		return &model.ChannelEntryDetails{}, info.ErrFetchTenantDetails
	}

	if commonFilter != nil {

		if commonFilter.Limit.IsSet() && commonFilter.Limit.Value() != nil {

			limit = *commonFilter.Limit.Value()
		}

		if commonFilter.Offset.IsSet() && commonFilter.Offset.Value() != nil {

			offset = *commonFilter.Offset.Value()
		}

		if commonFilter.IsActive.IsSet() && commonFilter.IsActive.Value() != nil {

			isActive = *commonFilter.IsActive.Value()
		}

		if commonFilter.Keyword.IsSet() && commonFilter.Keyword.Value() != nil {

			keyword = *commonFilter.Keyword.Value()
		}
	}

	if EntryFilter != nil {

		if EntryFilter.ChannelID.IsSet() && EntryFilter.ChannelID.Value() != nil {

			channelId = *EntryFilter.ChannelID.Value()
		}

		if EntryFilter.CategoryID.IsSet() && EntryFilter.CategoryID.Value() != nil {

			categoryId = *EntryFilter.CategoryID.Value()
		}

		if EntryFilter.CategorySlug.IsSet() && EntryFilter.CategorySlug.Value() != nil {

			categorySlug = *EntryFilter.CategorySlug.Value()
		}

		if EntryFilter.GetChildCategories.IsSet() && EntryFilter.GetChildCategories.Value() != nil && !*EntryFilter.GetChildCategories.Value() {

			selectedCategoriesFilter = true
		}

		if EntryFilter.Status.IsSet() && EntryFilter.Status.Value() != nil && *EntryFilter.Status.Value() != "" {

			entryStatus := *EntryFilter.Status.Value()

			switch {

			case entryStatus == string(Draft):

				status = string(Draft)

			case entryStatus == string(Publish):

				status = string(Publish)

			case entryStatus == string(Unpublish):

				status = string(Unpublish)
			}
		}
	}

	if sort != nil {

		if sort.SortBy.IsSet() && sort.SortBy.Value() != nil {

			sortBy = *sort.SortBy.Value()
		}

		if sort.Order.IsSet() && sort.Order.Value() != nil {

			order = *sort.Order.Value()
		}
	}

	if additionalData != nil {

		if additionalData.AuthorDetails.IsSet() && additionalData.AuthorDetails.Value() != nil {

			authorFlag = *additionalData.AuthorDetails.Value()
		}

		if additionalData.MemberProfile.IsSet() && additionalData.MemberProfile.Value() != nil {

			memberProfFlag = *additionalData.MemberProfile.Value()
		}

		if additionalData.Categories.IsSet() && additionalData.Categories.Value() != nil {

			categoriesFlag = *additionalData.Categories.Value()
		}

		if additionalData.AdditionalFields.IsSet() && additionalData.AdditionalFields.Value() != nil {

			fieldsFlg = *additionalData.AdditionalFields.Value()
		}
	}

	// fmt.Printf("limit: %v,offset: %v,isActive: %v,order: %v,sort: %v,title: %v,keyword: %v,channelid: %v,categoryid: %v,categorySlug: %v,status: %v,memflag: %v,categoryFlag: %v,authorflag: %v,fieldsFlag: %v\n", limit, offset, isActive, order, sortBy, title, keyword, channelId, categoryId, categorySlug, status, memberProfFlag, categoriesFlag, authorFlag, fieldsFlg)

	inputs := channels.EntriesInputs{
		ChannelId:              channelId,
		Limit:                  limit,
		Offset:                 offset,
		Keyword:                keyword,
		Status:                 status,
		Title:                  title,
		CategoryId:             categoryId,
		CategorySlug:           categorySlug,
		ActiveEntriesonly:      isActive,
		TenantId:               tenantDetails.TenantId,
		GetMemberProfile:       memberProfFlag,
		GetAuthorDetails:       authorFlag,
		GetAdditionalFields:    fieldsFlg,
		GetLinkedCategories:    categoriesFlag,
		SelectedCategoryFilter: selectedCategoriesFilter,
		TotalCount:             true,
		Order:                  order,
		SortBy:                 sortBy,
		SectionFieldTypeId:     12,
		MemberFieldTypeId:      14,
	}

	channelEntries, commonCount, _, err := ChannelConfigWP.FlexibleChannelEntriesList(inputs)

	if err != nil {

		ErrorLog.Printf("%v", err)

		c.AbortWithStatus(500)

		return &model.ChannelEntryDetails{}, err
	}

	finalChannelEntries := make([]model.ChannelEntries, len(channelEntries))

	for i, v := range channelEntries {

		finalChannelEntries[i].ID = v.Id

		finalChannelEntries[i].Title = v.Title

		author := v.Author

		finalChannelEntries[i].Author = &author

		finalChannelEntries[i].ChannelID = v.ChannelId

		finalChannelEntries[i].CategoriesID = v.CategoriesId

		finalChannelEntries[i].IsActive = v.IsActive

		finalChannelEntries[i].CoverImage = v.CoverImage

		createTime := v.CreateTime

		finalChannelEntries[i].CreateTime = &createTime

		finalChannelEntries[i].CreatedBy = v.CreatedBy

		finalChannelEntries[i].CreatedOn = v.CreatedOn

		finalChannelEntries[i].Description = scalars.CustomString(v.Description)

		excerpt := &v.Excerpt

		finalChannelEntries[i].Excerpt = excerpt

		finalChannelEntries[i].FeaturedEntry = v.Feature

		imgAltTag := v.ImageAltTag

		finalChannelEntries[i].ImageAltTag = &imgAltTag

		finalChannelEntries[i].Keyword = v.Keyword

		finalChannelEntries[i].MetaDescription = v.MetaDescription

		finalChannelEntries[i].MetaTitle = v.MetaTitle

		publishTime := v.PublishedTime

		finalChannelEntries[i].PublishedTime = &publishTime

		readTime := v.ReadingTime

		finalChannelEntries[i].ReadingTime = &readTime

		finalChannelEntries[i].RelatedArticles = v.RelatedArticles

		finalChannelEntries[i].ViewCount = v.ViewCount

		finalChannelEntries[i].Slug = v.Slug

		sortOrder := v.SortOrder

		finalChannelEntries[i].SortOrder = &sortOrder

		finalChannelEntries[i].Status = v.Status

		tags := v.Tags

		finalChannelEntries[i].Tags = &tags

		finalChannelEntries[i].ThumbnailImage = v.ThumbnailImage

		finalChannelEntries[i].UserID = v.UserId

		finalChannelEntries[i].TenantID = v.TenantId

		switch {

		case utf8.RuneCountInString(v.Description) > MaxChunkLength:

			chunks := FetchChunkData(v.Description)

			finalChannelEntries[i].ContentChunk = &model.Chunk{Data: chunks, Length: len(chunks)}

		default:

			finalChannelEntries[i].ContentChunk = &model.Chunk{Data: []string{v.Description}, Length: 1}
		}

		if v.AuthorDetail.Id != 0 {

			authorMobile := v.AuthorDetail.MobileNo

			authorActive := v.AuthorDetail.IsActive

			authorProfileImgPath := v.AuthorDetail.ProfileImagePath

			authorModon := v.AuthorDetail.ModifiedOn

			authorModby := v.AuthorDetail.ModifiedBy

			authorDetails := &model.Author{
				ID:               v.AuthorDetail.Id,
				FirstName:        v.AuthorDetail.FirstName,
				LastName:         v.AuthorDetail.LastName,
				Email:            v.AuthorDetail.Email,
				MobileNo:         &authorMobile,
				IsActive:         &authorActive,
				ProfileImagePath: &authorProfileImgPath,
				CreatedOn:        v.AuthorDetail.CreatedOn,
				CreatedBy:        v.AuthorDetail.CreatedBy,
				ModifiedOn:       &authorModon,
				ModifiedBy:       &authorModby,
				TenantID:         v.AuthorDetail.TenantId,
			}

			finalChannelEntries[i].AuthorDetails = authorDetails

		}

		if v.MemberProfiles.Id != 0 {

			profilePage := v.MemberProfiles.ProfilePage

			companyName := v.MemberProfiles.CompanyName

			companyLocation := v.MemberProfiles.CompanyLocation

			companyLogo := v.MemberProfiles.CompanyLogo

			aboutComapny := v.MemberProfiles.About

			seoTitle := v.MemberProfiles.SeoTitle

			seoDesc := v.MemberProfiles.SeoDescription

			seoKey := v.MemberProfiles.SeoKeyword

			linkedin := v.MemberProfiles.Linkedin

			twitter := v.MemberProfiles.Twitter

			website := v.MemberProfiles.Website

			profModon := v.MemberProfiles.ModifiedOn

			profModby := v.MemberProfiles.ModifiedBy

			claimStatus := v.MemberProfiles.ClaimStatus

			claimDate := v.MemberProfiles.ClaimDate

			createdOn := v.MemberProfiles.CreatedOn

			createdBy := v.MemberProfiles.CreatedBy

			memberProfile := &model.MemberProfile{
				ID:              v.MemberProfiles.Id,
				MemberID:        v.MemberProfiles.MemberId,
				ProfileName:     v.MemberProfiles.ProfileName,
				ProfileSlug:     v.MemberProfiles.ProfileSlug,
				ProfilePage:     &profilePage,
				MemberDetails:   v.MemberProfiles.MemberDetails,
				CompanyName:     &companyName,
				CompanyLocation: &companyLocation,
				CompanyLogo:     &companyLogo,
				About:           &aboutComapny,
				SeoTitle:        &seoTitle,
				SeoDescription:  &seoDesc,
				SeoKeyword:      &seoKey,
				Linkedin:        &linkedin,
				Twitter:         &twitter,
				Website:         &website,
				CreatedBy:       &createdBy,
				CreatedOn:       &createdOn,
				ModifiedOn:      &profModon,
				ModifiedBy:      &profModby,
				ClaimStatus:     &claimStatus,
				TenantID:        v.MemberProfiles.TenantId,
				ClaimDate:       &claimDate,
			}

			finalChannelEntries[i].MemberProfile = memberProfile
		}

		if len(v.Categories) > 0 {

			conv_categories := make([][]model.Category, len(v.Categories))

			for cat_index, categories := range v.Categories {

				conv_categoryz := make([]model.Category, len(categories))

				for i, category := range categories {

					categoryModon := category.ModifiedOn

					categoryModBy := category.ModifiedBy

					conv_category := model.Category{
						ID:           category.Id,
						CategoryName: category.CategoryName,
						CategorySlug: category.CategorySlug,
						Description:  category.Description,
						ImagePath:    category.ImagePath,
						CreatedOn:    category.CreatedOn,
						CreatedBy:    category.CreatedBy,
						ModifiedOn:   &categoryModon,
						ModifiedBy:   &categoryModBy,
						ParentID:     category.ParentId,
						TenantID:     category.TenantId,
					}

					conv_categoryz[i] = conv_category

				}

				conv_categories[cat_index] = conv_categoryz
			}

			finalChannelEntries[i].Categories = conv_categories
		}

		conv_sections := make([]model.Section, len(v.Sections))

		if len(v.Sections) > 0 {

			for section_index, section := range v.Sections {

				sectionModon := section.ModifiedOn

				sectionModBy := section.ModifiedBy

				conv_section := model.Section{
					ID:            section.Id,
					SectionName:   section.FieldName,
					SectionTypeID: section.FieldTypeId,
					CreatedOn:     section.CreatedOn,
					CreatedBy:     section.CreatedBy,
					ModifiedOn:    &sectionModon,
					ModifiedBy:    &sectionModBy,
					OrderIndex:    section.OrderIndex,
					TenantID:      section.TenantId,
				}

				conv_sections[section_index] = conv_section

			}
		}

		conv_fields := make([]model.Field, len(v.Fields))

		if len(v.Fields) > 0 {

			for field_index, field := range v.Fields {

				fieldValueModon := field.FieldValue.ModifiedOn

				fieldValueModBy := field.FieldValue.ModifiedBy

				conv_field_value := model.FieldValue{
					ID:         field.FieldValue.FieldId,
					FieldValue: field.FieldValue.FieldValue,
					CreatedOn:  field.FieldValue.CreatedOn,
					CreatedBy:  field.FieldValue.CreatedBy,
					ModifiedOn: &fieldValueModon,
					ModifiedBy: &fieldValueModBy,
					TenantID:   field.FieldValue.TenantId,
				}

				conv_fieldOptions := make([]model.FieldOptions, len(field.FieldOptions))

				for option_index, field_option := range field.FieldOptions {

					optionModOn := field_option.ModifiedOn

					optionModBy := field_option.ModifiedBy

					conv_fieldOption := model.FieldOptions{
						ID:          field_option.Id,
						OptionName:  field_option.OptionName,
						OptionValue: field_option.OptionValue,
						CreatedOn:   field_option.CreatedOn,
						CreatedBy:   field_option.CreatedBy,
						ModifiedOn:  &optionModOn,
						ModifiedBy:  &optionModBy,
						TenantID:    field_option.TenantId,
					}

					conv_fieldOptions[option_index] = conv_fieldOption
				}

				fieldModon := field.ModifiedOn

				fieldModBy := field.ModifiedBy

				fieldDateTime := field.DatetimeFormat

				fieldTime := field.TimeFormat

				fieldSectionParentId := field.SectionParentId

				fieldCharAllowed := field.CharacterAllowed

				conv_field := model.Field{
					ID:               field.Id,
					FieldName:        field.FieldName,
					FieldTypeID:      field.FieldTypeId,
					MandatoryField:   field.MandatoryField,
					OptionExist:      field.OptionExist,
					CreatedOn:        field.CreatedOn,
					CreatedBy:        field.CreatedBy,
					ModifiedOn:       &fieldModon,
					ModifiedBy:       &fieldModBy,
					FieldDesc:        field.FieldDesc,
					OrderIndex:       field.OrderIndex,
					ImagePath:        field.ImagePath,
					DatetimeFormat:   &fieldDateTime,
					TimeFormat:       &fieldTime,
					SectionParentID:  &fieldSectionParentId,
					CharacterAllowed: &fieldCharAllowed,
					FieldTypeName:    field.FieldTypeName,
					FieldValue:       &conv_field_value,
					FieldOptions:     conv_fieldOptions,
					TenantID:         field.TenantId,
				}

				conv_fields[field_index] = conv_field

			}
		}

		finalChannelEntries[i].AdditionalFields = &model.AdditionalFields{Sections: conv_sections, Fields: conv_fields}

	}

	return &model.ChannelEntryDetails{ChannelEntriesList: finalChannelEntries, Count: commonCount}, nil
}

func ChannelEntryDetail(ctx context.Context, id *int, slug *string, additionalData *model.EntriesAdditionalData, channelId *int) (*model.ChannelEntries, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.ChannelEntries{}, info.ErrGinCtx
	}

	if id == nil && slug == nil {

		ErrorLog.Printf("%v", info.ErrReqMandatory)

		c.AbortWithStatus(400)

		return &model.ChannelEntries{}, info.ErrReqMandatory
	}

	var (
		entryId, chanId                                     int
		entrySlug                                           string
		categoryFlag, fieldsFlg, memberProfFlag, authorFlag bool
	)

	if id != nil && *id != 0 {

		entryId = *id
	}

	if slug != nil && *slug != "" {

		entrySlug = *slug
	}

	if channelId != nil && *channelId != 0 {
		chanId = *channelId
	}

	fmt.Println("chaannan", chanId)

	if additionalData != nil {

		if additionalData.AuthorDetails.IsSet() && additionalData.AuthorDetails.Value() != nil {

			authorFlag = *additionalData.AuthorDetails.Value()
		}

		if additionalData.MemberProfile.IsSet() && additionalData.MemberProfile.Value() != nil {

			memberProfFlag = *additionalData.MemberProfile.Value()
		}

		if additionalData.Categories.IsSet() && additionalData.Categories.Value() != nil {

			categoryFlag = *additionalData.Categories.Value()
		}

		if additionalData.AdditionalFields.IsSet() && additionalData.AdditionalFields.Value() != nil {

			fieldsFlg = *additionalData.AdditionalFields.Value()
		}
	}

	tenantData, err := GetTenantDetails(c)

	if err != nil {

		return &model.ChannelEntries{}, err
	}

	inputs := channels.EntriesInputs{
		Id:                  entryId,
		Slug:                entrySlug,
		TenantId:            tenantData.TenantId,
		GetMemberProfile:    memberProfFlag,
		GetLinkedCategories: categoryFlag,
		GetAdditionalFields: fieldsFlg,
		GetAuthorDetails:    authorFlag,
		ChannelId:           chanId,
	}

	channelEntry, _, err := ChannelConfigWP.FetchChannelEntryDetail(inputs, nil)

	fmt.Println("aithorDetails", channelEntry.AuthorDetail.CreatedOn)

	switch {

	case err != nil:

		return &model.ChannelEntries{}, err

	case channelEntry.Id == 0:

		return &model.ChannelEntries{}, info.ErrRecordNotFound

	}

	conv_categories := make([][]model.Category, len(channelEntry.Categories))

	if len(channelEntry.Categories) > 0 {

		for cat_index, categories := range channelEntry.Categories {

			conv_categoryz := make([]model.Category, len(categories))

			for i, category := range categories {

				categoryModon := category.ModifiedOn

				categoryModBy := category.ModifiedBy

				conv_category := model.Category{
					ID:           category.Id,
					CategoryName: category.CategoryName,
					CategorySlug: category.CategorySlug,
					Description:  category.Description,
					ImagePath:    category.ImagePath,
					CreatedOn:    category.CreatedOn,
					CreatedBy:    category.CreatedBy,
					ModifiedOn:   &categoryModon,
					ModifiedBy:   &categoryModBy,
					ParentID:     category.ParentId,
					TenantID:     category.TenantId,
				}

				conv_categoryz[i] = conv_category

			}

			conv_categories[cat_index] = conv_categoryz
		}

	}

	var memberProfile model.MemberProfile

	if channelEntry.MemberProfiles.Id != 0 {

		memberProfile = model.MemberProfile{
			ID:              channelEntry.MemberProfiles.Id,
			MemberID:        channelEntry.MemberProfiles.MemberId,
			ProfileName:     channelEntry.MemberProfiles.ProfileName,
			ProfileSlug:     channelEntry.MemberProfiles.ProfileSlug,
			ProfilePage:     &channelEntry.MemberProfiles.ProfilePage,
			MemberDetails:   channelEntry.MemberProfiles.MemberDetails,
			CompanyName:     &channelEntry.MemberProfiles.CompanyName,
			CompanyLocation: &channelEntry.MemberProfiles.CompanyLocation,
			CompanyLogo:     &channelEntry.MemberProfiles.CompanyLogo,
			About:           &channelEntry.MemberProfiles.About,
			SeoTitle:        &channelEntry.MemberProfiles.SeoTitle,
			SeoDescription:  &channelEntry.MemberProfiles.SeoDescription,
			SeoKeyword:      &channelEntry.MemberProfiles.SeoKeyword,
			Linkedin:        &channelEntry.MemberProfiles.Linkedin,
			Twitter:         &channelEntry.MemberProfiles.Twitter,
			Website:         &channelEntry.MemberProfiles.Website,
			CreatedBy:       &channelEntry.MemberProfiles.CreatedBy,
			CreatedOn:       &channelEntry.MemberProfiles.CreatedOn,
			ModifiedOn:      &channelEntry.MemberProfiles.ModifiedOn,
			ModifiedBy:      &channelEntry.MemberProfiles.ModifiedBy,
			ClaimStatus:     &channelEntry.MemberProfiles.ClaimStatus,
			TenantID:        channelEntry.MemberProfiles.TenantId,
			ClaimDate:       &channelEntry.MemberProfiles.ClaimDate,
		}
	}

	var authorDetails model.Author

	if channelEntry.AuthorDetail.Id != 0 {

		authorDetails = model.Author{
			ID:               channelEntry.AuthorDetail.Id,
			FirstName:        channelEntry.AuthorDetail.FirstName,
			LastName:         channelEntry.AuthorDetail.LastName,
			Email:            channelEntry.AuthorDetail.Email,
			MobileNo:         &channelEntry.AuthorDetail.MobileNo,
			IsActive:         &channelEntry.AuthorDetail.IsActive,
			ProfileImagePath: &channelEntry.AuthorDetail.ProfileImagePath,
			CreatedOn:        channelEntry.AuthorDetail.CreatedOn,
			CreatedBy:        channelEntry.AuthorDetail.CreatedBy,
			ModifiedOn:       &channelEntry.AuthorDetail.ModifiedOn,
			ModifiedBy:       &channelEntry.AuthorDetail.ModifiedBy,
			TenantID:         channelEntry.AuthorDetail.TenantId,
		}

	}

	conv_sections := make([]model.Section, len(channelEntry.Sections))

	if len(channelEntry.Sections) > 0 {

		for section_index, section := range channelEntry.Sections {

			sectionModon := section.ModifiedOn

			sectionModBy := section.ModifiedBy

			conv_section := model.Section{
				ID:            section.Id,
				SectionName:   section.FieldName,
				SectionTypeID: section.FieldTypeId,
				CreatedOn:     section.CreatedOn,
				CreatedBy:     section.CreatedBy,
				ModifiedOn:    &sectionModon,
				ModifiedBy:    &sectionModBy,
				OrderIndex:    section.OrderIndex,
				TenantID:      section.TenantId,
			}

			conv_sections[section_index] = conv_section

		}

	}

	conv_fields := make([]model.Field, len(channelEntry.Fields))

	if len(channelEntry.Fields) > 0 {

		for field_index, field := range channelEntry.Fields {

			fieldValueModon := field.FieldValue.ModifiedOn

			fieldValueModBy := field.FieldValue.ModifiedBy

			conv_field_value := model.FieldValue{
				ID:         field.FieldValue.FieldId,
				FieldValue: field.FieldValue.FieldValue,
				CreatedOn:  field.FieldValue.CreatedOn,
				CreatedBy:  field.FieldValue.CreatedBy,
				ModifiedOn: &fieldValueModon,
				ModifiedBy: &fieldValueModBy,
				TenantID:   field.FieldValue.TenantId,
			}

			conv_fieldOptions := make([]model.FieldOptions, len(field.FieldOptions))

			for option_index, field_option := range field.FieldOptions {

				optionModOn := field_option.ModifiedOn

				optionModBy := field_option.ModifiedBy

				conv_fieldOption := model.FieldOptions{
					ID:          field_option.Id,
					OptionName:  field_option.OptionName,
					OptionValue: field_option.OptionValue,
					CreatedOn:   field_option.CreatedOn,
					CreatedBy:   field_option.CreatedBy,
					ModifiedOn:  &optionModOn,
					ModifiedBy:  &optionModBy,
					TenantID:    field_option.TenantId,
				}

				conv_fieldOptions[option_index] = conv_fieldOption
			}

			fieldModon := field.ModifiedOn

			fieldModBy := field.ModifiedBy

			fieldDateTime := field.DatetimeFormat

			fieldTime := field.TimeFormat

			fieldSectionParentId := field.SectionParentId

			fieldCharAllowed := field.CharacterAllowed

			conv_field := model.Field{
				ID:               field.Id,
				FieldName:        field.FieldName,
				FieldTypeID:      field.FieldTypeId,
				MandatoryField:   field.MandatoryField,
				OptionExist:      field.OptionExist,
				CreatedOn:        field.CreatedOn,
				CreatedBy:        field.CreatedBy,
				ModifiedOn:       &fieldModon,
				ModifiedBy:       &fieldModBy,
				FieldDesc:        field.FieldDesc,
				OrderIndex:       field.OrderIndex,
				ImagePath:        field.ImagePath,
				DatetimeFormat:   &fieldDateTime,
				TimeFormat:       &fieldTime,
				SectionParentID:  &fieldSectionParentId,
				CharacterAllowed: &fieldCharAllowed,
				FieldTypeName:    field.FieldTypeName,
				FieldValue:       &conv_field_value,
				FieldOptions:     conv_fieldOptions,
				TenantID:         field.TenantId,
			}

			conv_fields[field_index] = conv_field

		}
	}

	additionalFields := model.AdditionalFields{Sections: conv_sections, Fields: conv_fields}

	// desc := scalars.CustomString(channelEntry.Description)

	// decsByte,_ := desc.MarshalJSON()

	convChannelEntry := model.ChannelEntries{
		ID:               channelEntry.Id,
		Title:            channelEntry.Title,
		Slug:             channelEntry.Slug,
		Description:      scalars.CustomString(channelEntry.Description),
		UserID:           channelEntry.UserId,
		ChannelID:        channelEntry.ChannelId,
		Status:           channelEntry.Status,
		IsActive:         channelEntry.IsActive,
		CreatedOn:        channelEntry.CreatedOn,
		CreatedBy:        channelEntry.CreatedBy,
		ModifiedBy:       &channelEntry.ModifiedBy,
		ModifiedOn:       &channelEntry.ModifiedOn,
		CoverImage:       channelEntry.CoverImage,
		ThumbnailImage:   channelEntry.ThumbnailImage,
		MetaTitle:        channelEntry.MetaTitle,
		MetaDescription:  channelEntry.MetaDescription,
		Keyword:          channelEntry.Keyword,
		CategoriesID:     channelEntry.CategoriesId,
		RelatedArticles:  channelEntry.RelatedArticles,
		FeaturedEntry:    channelEntry.Feature,
		ViewCount:        channelEntry.ViewCount,
		Author:           &channelEntry.Author,
		SortOrder:        &channelEntry.SortOrder,
		CreateTime:       &channelEntry.CreateTime,
		PublishedTime:    &channelEntry.PublishedTime,
		ReadingTime:      &channelEntry.ReadingTime,
		Tags:             &channelEntry.Tags,
		Excerpt:          &channelEntry.Excerpt,
		ImageAltTag:      &channelEntry.ImageAltTag,
		TenantID:         channelEntry.TenantId,
		Categories:       conv_categories,
		MemberProfile:    &memberProfile,
		AdditionalFields: &additionalFields,
		AuthorDetails:    &authorDetails,
	}

	switch {

	case utf8.RuneCountInString(channelEntry.Description) > MaxChunkLength:

		chunks := FetchChunkData(channelEntry.Description)

		convChannelEntry.ContentChunk = &model.Chunk{Data: chunks, Length: len(chunks)}

	default:

		convChannelEntry.ContentChunk = &model.Chunk{Data: []string{channelEntry.Description}, Length: 1}
	}

	fmt.Println("aouthoorrr", convChannelEntry.AuthorDetails.CreatedOn)

	return &convChannelEntry, nil
}

func UpdateEntryViewCount(ctx context.Context, id *int, slug *string) (*model.CountUpdate, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.CountUpdate{Count: 0, Status: false}, info.ErrGinCtx
	}

	if id == nil && slug == nil {

		ErrorLog.Printf("%v", info.ErrReqMandatory)

		c.AbortWithStatus(400)

		return &model.CountUpdate{Count: 0, Status: false}, info.ErrReqMandatory
	}

	var (
		entryId   int
		entrySlug string
	)

	if id != nil && *id != 0 {

		entryId = *id
	}

	if slug != nil && *slug != "" {

		entrySlug = *slug
	}

	tenantData, err := GetTenantDetails(c)

	if err != nil {

		return &model.CountUpdate{Count: 0, Status: false}, err
	}

	viewCount, err := ChannelConfigWP.UpdateChannelEntryViewCount(entryId, entrySlug, tenantData.TenantId)

	if err != nil {

		switch {

		case fmt.Sprintf("%v", err) == fmt.Sprintf("%v", info.ErrNoRowsAffected):

			return &model.CountUpdate{Count: 0, Status: false}, info.ErrUpdateViewCount

		default:

			return &model.CountUpdate{Count: 0, Status: false}, err
		}
	}

	return &model.CountUpdate{Count: viewCount, Status: true}, nil
}
