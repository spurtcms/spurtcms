package controller

import (
	"context"
	"spurt-cms/graphql/info"
	"spurt-cms/graphql/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CourseList Api

func CourseList(ctx context.Context, tenantId *string, status *int, categoryid *int) (*model.CourseLists, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.CourseLists{}, info.ErrGinCtx
	}

	var tenantid string

	if *tenantId != "" {

		tenantid = *tenantId
	}

	var Status int

	if *status != 0 {

		Status = *status
	}

	var CategoryId int

	if *categoryid != 0 {

		CategoryId = *categoryid
	}

	courselist, count, err := CoursesInstance.CourseList(Status, tenantid, CategoryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(400)
			return nil, err
		}
		ErrorLog.Printf("%v", err)
		c.AbortWithStatus(500)
		return nil, err
	}

	var finalCourseList []model.Course

	for _, val := range courselist {

		value := model.Course{
			ID:          val.Id,
			Title:       val.Title,
			Description: val.Description,
			ImageName:   val.ImageName,
			ImagePath:   val.ImagePath,
			Category:    val.CategoryId,
		}

		finalCourseList = append(finalCourseList, value)

	}

	list := model.CourseLists{
		Courselist: finalCourseList,
		Count:      int(count),
	}

	return &list, nil

}

//CourseDetails Api

func CourseDetails(ctx context.Context, tenantId *string, courseId *int) (*model.CourseDetails, error) {

	c, ok := ctx.Value(GinContext).(*gin.Context)

	if !ok {

		ErrorLog.Printf("%v", info.ErrGinCtx)

		return &model.CourseDetails{}, info.ErrGinCtx
	}

	var tenantid string

	if *tenantId != "" {

		tenantid = *tenantId
	}

	var CourseId int

	if *courseId != 0 {

		CourseId = *courseId
	}

	sectionlist, serr := CoursesInstance.SectionList(CourseId, tenantid)
	if serr != nil {
		if serr == gorm.ErrRecordNotFound {
			c.AbortWithStatus(400)
			return nil, serr
		}
		ErrorLog.Printf("%v", serr)
		c.AbortWithStatus(500)
		return nil, serr
	}

	var FinalSectionList []model.Sections

	for _, sectionValue := range sectionlist {

		SectionList := model.Sections{
			ID:         sectionValue.Id,
			CourseID:   sectionValue.CourseId,
			Title:      sectionValue.Title,
			Content:    sectionValue.Content,
			OrderIndex: sectionValue.OrderIndex,
			TenantID:   sectionValue.TenantId,
		}

		FinalSectionList = append(FinalSectionList, SectionList)
	}

	lessonlist, lerr := CoursesInstance.LessonList(CourseId, tenantid)
	if lerr != nil {
		if lerr == gorm.ErrRecordNotFound {
			c.AbortWithStatus(400)
			return nil, lerr
		}
		ErrorLog.Printf("%v", lerr)
		c.AbortWithStatus(500)
		return nil, lerr
	}

	var FinalLessonList []model.Lesson

	for _, lessonValue := range lessonlist {

		LessonList := model.Lesson{
			ID:         lessonValue.Id,
			CourseID:   lessonValue.CourseId,
			SectionID:  lessonValue.SectionId,
			Title:      lessonValue.Title,
			Content:    lessonValue.Content,
			EmbedLink:  lessonValue.EmbedLink,
			FileName:   lessonValue.FileName,
			FilePath:   lessonValue.FilePath,
			LessonType: lessonValue.LessonType,
			OrderIndex: lessonValue.OrderIndex,
			TenantID:   lessonValue.TenantId,
		}

		FinalLessonList = append(FinalLessonList, LessonList)
	}

	return &model.CourseDetails{Lesson: FinalLessonList, Section: FinalSectionList}, nil

}
