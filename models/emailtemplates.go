package models

import (
	"spurt-cms/config"
	"time"
	"fmt"
)

type Filter struct {
	Keyword     string
	ChannelName string
	Status      string
	Title       string
	Category    string
}

type TblEmailTemplate struct {
	Id                  int
	TemplateName        string
	TemplateSubject     string
	TemplateMessage     string
	TemplateDescription string
	ModuleId            int
	CreatedOn           time.Time
	CreatedBy           int
	ModifiedOn          time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"DEFAULT:NULL"`
	DeletedOn           time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy           int       `gorm:"DEFAULT:NULL"`
	IsDeleted           int       `gorm:"DEFAULT:0"`
	IsActive            int
	IsDefault           int
	DateString          string `gorm:"-"`
	TemplateSlug        string
	TenantId            int `gorm:"DEFAULT:1"`
}

type TblImportfileLogs struct {
	Id             int
	FileName       string
	TotalFields    int
	MappedFields   int
	UnmappedFields int
	CreatedOn      time.Time
	CreatedBy      int
	Status         int
	DateString     string `gorm:"-"`
	TenantId       int    `gorm:"DEFAULT:1"`
}

var DB = config.SetupDB()

func GetTemplateList(templates *[]TblEmailTemplate, limit, offset int, filter Filter, flag bool, tenantid int) ([]TblEmailTemplate, int64) {

	var Total_templates int64

	var emptyslice []TblEmailTemplate

	query := DB.Debug().Table("tbl_email_templates").Where("tbl_email_templates.is_deleted=0 and tbl_email_templates.module_id=0 and tbl_email_templates.tenant_id=?",tenantid).Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_email_templates.template_name)) ILIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%")

	}

	if flag {

		query.Find(&templates)

		return *templates, 0

	}

	if limit != 0 && !flag {

		query.Offset(offset).Limit(limit).Find(&templates)

		return *templates, 0

	} else {

		query.Find(&templates).Count(&Total_templates)

		return emptyslice, Total_templates
	}

}

func GetTempdetail(temp *TblEmailTemplate, id int, _ int) error {

	if err := DB.Table("tbl_email_templates").Where("is_deleted =0 and id=?", id).First(&temp).Error; err != nil {

		return err
	}
	return nil
}

func UpdateTemplate(template *TblEmailTemplate, _ int) error {

	if err := DB.Table("tbl_email_templates").Where("is_deleted=0 and id=?", template.Id).Updates(TblEmailTemplate{TemplateName: template.TemplateName, TemplateSubject: template.TemplateSubject, TemplateMessage: template.TemplateMessage, Id: template.Id, ModifiedOn: template.ModifiedOn, ModifiedBy: template.ModifiedBy, TemplateDescription: template.TemplateDescription}).Error; err != nil {

		return err
	}

	return nil
}

func TemplateStatus(id int, status int, userid int, _ int) error {

	var Tempstatus TblEmailTemplate

	Tempstatus.ModifiedBy = userid

	Tempstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := TemplateIsActive(&Tempstatus, id, status,-1)

	if err1 != nil {

		return err1
	}

	return nil

}

/*update template isactive*/
func TemplateIsActive(template *TblEmailTemplate, id int, val int,_ int) error {

	if err := DB.Table("tbl_email_templates").Where("id=? and is_deleted = 0", id).UpdateColumns(map[string]interface{}{"is_active": val, "modified_by": template.ModifiedBy, "modified_on": template.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func GetTemplates(template *TblEmailTemplate, key string, tenantid int) error {
	fmt.Println("tenantsss",tenantid)
if tenantid<=0{

	if err := DB.Debug().Table("tbl_email_templates").Where("template_slug=? and is_deleted = 0 ", key).First(&template).Error; err != nil {

		return err
	}
}else{
	if err := DB.Debug().Table("tbl_email_templates").Where("template_slug=? and is_deleted = 0 and tenant_id=?", key,tenantid).First(&template).Error; err != nil {

		return err
	}
}
	return nil
}

func CreateFileLog(logdata *TblImportfileLogs, tenantid int) error {

	if err := DB.Debug().Table("tbl_importfile_logs").Where("tenant_id = ?", tenantid).Create(&logdata).Error; err != nil {

		return err
	}

	return nil

}
func GetFileLogs(logdata *[]TblImportfileLogs, tenantid int) ([]TblImportfileLogs, error) {

	if err := DB.Debug().Table("tbl_importfile_logs").Where("tenant_id = ?", tenantid).Select("tbl_importfile_logs.*").Order("id desc").Find(&logdata).Error; err != nil {

		return *logdata, err
	}

	return *logdata, nil

}

func GetTemplatesByModuleId(moduleid int, tenantid int) (templates *[]TblEmailTemplate, error bool) {
	fmt.Println("moduleid",moduleid)

	if err := DB.Debug().Table("tbl_email_templates").Where("is_deleted = 0 and module_id=? and tenant_id=?", moduleid,tenantid).Order("id asc").Find(&templates).Error; err != nil {

		return &[]TblEmailTemplate{}, false
	}

	return templates, true
}
