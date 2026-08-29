package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Payload string

type Method string

type TblWebhooks struct {
	Id            int
	WebhookName   string
	PayloadType   Payload `gorm:"DEFAULT:NULL"`
	RequestMethod Method
	RequestUrl    string
	EventType     string
	IsActive      int
	Headers       json.RawMessage `gorm:"DEFAULT:NULL"`
	Fields        json.RawMessage `gorm:"DEFAULT:NULL"`
	TenantId      string
	CreatedBy     int
	CreatedOn     time.Time
	ModifiedBy    int         `gorm:"DEFAULT:NULL"`
	ModifiedOn    time.Time   `gorm:"DEFAULT:NULL"`
	IsDeleted     int         `gorm:"DEFAULT:0"`
	DeletedBy     int         `gorm:"DEFAULT:NULL"`
	DeletedOn     time.Time   `gorm:"DEFAULT:NULL"`
	ParseHeaders  interface{} `gorm:"-"`
}

type WebhookFilter struct {
	Id        int
	Limit     int
	Offset    int
	Keyword   string
	Title     string
	EventType string
	IsActive  int
	Order     int
	Sort      string
	TenantId  string
	List      bool
	FetchOne  bool
	Count     bool
}

func CreateWebhook(webhook *TblWebhooks) error {

	if err := DB.Create(&webhook).Error; err != nil {
		return err
	}

	return nil
}

func (wf WebhookFilter) FetchWebhook(webhooks *[]TblWebhooks, webhook *TblWebhooks, count *int64) error {

	query := DB.Debug().Model(TblWebhooks{}).Where("is_deleted=0")

	if wf.Id != 0 {

		query = query.Where("id=?", wf.Id)
	}

	if wf.TenantId > "" {

		query = query.Where("tenant_id=?", wf.TenantId)
	}

	if wf.IsActive > -1 {

		query = query.Where("is_active=?", wf.IsActive)
	}

	if wf.Keyword != "" {

		query = query.Where("TRIM(LOWER(webhook_name)) LIKE TRIM(LOWER(?))", "%"+wf.Keyword+"%")
	}

	if wf.EventType!=""{

		query = query.Where("event_type=?",wf.EventType)
	}

	if wf.Title != "" {

		query = query.Where("webhook_name=?", wf.Title)
	}

	if wf.FetchOne {

		if err := query.First(&webhook).Error; err != nil {

			return err
		}
	}

	if wf.Count {

		if err := query.Count(count).Error; err != nil {

			return err
		}
	}

	if wf.List {

		if wf.Sort != "" {

			var order string

			if wf.Order == 1 {

				order = wf.Sort + " desc"

			} else {

				order = wf.Sort
			}

			query = query.Order(order)

		} else {

			query = query.Order("id desc")
		}

		if wf.Limit > 0 {

			query = query.Limit(wf.Limit)
		}

		if wf.Offset > -1 {

			query = query.Offset(wf.Offset)
		}

		if err := query.Find(&webhooks).Error; err != nil {

			return err
		}
	}

	return nil
}

func UpdateWebhook(webhook *TblWebhooks) error {

	// tx := DB.Begin()

	// if err := tx.Error; err != nil {

	// 	return err
	// }

	// if err := tx.Debug().Where("is_deleted=0 and tenant_id=?", webhook.TenantId).Updates(webhook).Error; err != nil {

	// 	tx.Rollback()

	// 	return err
	// }

	// if err := tx.Where("is_deleted=0 and id=? and tenant_id=?",webhook.Id,webhook.TenantId).First(&webhook).Error; err != nil {

	// 	tx.Rollback()

	// 	return err
	// }

	// if err := tx.Commit().Error; err != nil {

	// 	tx.Rollback()

	// 	return err
	// }

	pipeline := func(tx *gorm.DB) error {

		switch {

		case webhook.IsActive == 0 || webhook.PayloadType == "":

			mapData := map[string]interface{}{"id": webhook.Id, "webhook_name": webhook.WebhookName, "payload_type": webhook.PayloadType, "request_method": webhook.RequestMethod, "request_url": webhook.RequestUrl, "event_type": webhook.EventType, "is_active": webhook.IsActive, "tenant_id": webhook.TenantId, "modified_on": webhook.ModifiedOn, "modified_by": webhook.ModifiedBy, "headers": webhook.Headers, "fields": webhook.Fields}

			if err := tx.Model(TblWebhooks{}).Where("is_deleted=0 and id=? and tenant_id=?", webhook.Id, webhook.TenantId).UpdateColumns(mapData).Error; err != nil {

				return err
			}

		default:

			if err := tx.Where("is_deleted=0 and tenant_id=?", webhook.TenantId).Updates(webhook).Error; err != nil {

				return err
			}

		}

		if err := tx.Where("is_deleted=0 and tenant_id=?", webhook.TenantId).First(&webhook).Error; err != nil {

			return err
		}

		return nil
	}

	if err := DB.Debug().Transaction(pipeline); err != nil {

		return err
	}

	return nil
}

func DeleteWebhook(webhook TblWebhooks) error {

	if err := DB.Debug().Where("is_deleted=0 and tenant_id=?", webhook.TenantId).Updates(webhook).Error; err != nil {

		return err
	}

	return nil
}
