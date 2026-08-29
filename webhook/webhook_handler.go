package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"spurt-cms/models"
	"strings"

	"errors"
	"net/http"

	"net/url"
)

type EventType string

const (
	EntryCreate   EventType = "entry.create"
	EntryPublish  EventType = "entry.publish"
	EntryUnpubish EventType = "entry.unpublish"
	EntryDelete   EventType = "entry.delete"
	MaxTurns                = 2
)

type WebhookClient struct {
	TenantId  string
	EventType EventType
	Data      map[string]interface{}
}

type WebhookTrigger struct {
	Url         string
	Method      string
	Headers     []interface{}
	PayloadType string
	Payload     []Payload
	CustomData  map[string]interface{}
}

type Payload struct {
	FieldName  string `json:"wh_field_name"`
	FieldValue string `json:"wh_field_value"`
}

func (wc *WebhookClient) FetchWebhooks() ([]models.TblWebhooks, error) {

	model := models.WebhookFilter{TenantId: wc.TenantId, IsActive: 1, List: true, EventType: string(wc.EventType)}

	webhooks := []models.TblWebhooks{}

	err := model.FetchWebhook(&webhooks, nil, nil)

	if err != nil {

		return []models.TblWebhooks{}, err
	}

	fmt.Println(webhooks)

	return webhooks, nil
}

func (wc *WebhookClient) HandleWebhooks() (map[int]interface{}, bool) {

	var (
		err    error
		WhResp = make(map[int]interface{})
	)

	webhooks, err := wc.FetchWebhooks()

	if err != nil {

		return map[int]interface{}{}, false
	}

	for _, wh := range webhooks {

		var (
			whErr       error
			headers     []interface{}
			fields      []Payload
			payloadType string
			turns       = 1
		)

		if wh.Headers != nil {

			whErr = json.Unmarshal(wh.Headers, &headers)

			if whErr != nil {

				WhResp[wh.Id] = nil

				err = whErr

				continue
			}
		}

		if wh.Fields != nil {

			whErr = json.Unmarshal(wh.Fields, &fields)

			if whErr != nil {

				WhResp[wh.Id] = nil

				err = whErr

				continue
			}
		}

		if wh.PayloadType != "" {

			payloadType = string(wh.PayloadType)
		}

		trgrInputs := WebhookTrigger{
			Url:         wh.RequestUrl,
			Method:      string(wh.RequestMethod),
			Headers:     headers,
			PayloadType: payloadType,
			Payload:     fields,
			CustomData:  wc.Data,
		}

	trgrLoop:

		for turns <= MaxTurns {

			fmt.Println("turns: ", turns)

			resp, trgrErr := TriggerWebhooks(trgrInputs)

			fmt.Printf("finalresp: %v\n", err)

			switch {

			case trgrErr == nil:

				fmt.Println("1-->case")

				WhResp[wh.Id] = resp.Body

				break trgrLoop

			case trgrErr != nil && turns == MaxTurns:

				fmt.Println("2-->case")

				WhResp[wh.Id] = nil

				err = trgrErr

				break trgrLoop

			default:

				turns++
			}

		}

	}

	return WhResp, err == nil
}

func TriggerWebhooks(trigger WebhookTrigger) (*http.Response, error) {

	fmt.Println("trigger", trigger)

	client := http.Client{}

	// var client http.Client

	var (
		req    *http.Request
		err    error
		method string = strings.ToUpper(trigger.Method)
	)

	switch method {

	case "GET", "DELETE":

		if len(trigger.Payload) > 0 {

			queryParams := url.Values{}

			for _, field := range trigger.Payload {

				for k2, v2 := range trigger.CustomData {

					if field.FieldValue == k2 {

						queryParams.Set(field.FieldName, v2.(string))
					}
				}
			}

			parseUrl, _ := url.Parse(trigger.Url)

			queries := queryParams.Encode()

			parseUrl.RawQuery = queries

			req, err = http.NewRequest(method, parseUrl.String(), nil)

			if err != nil {

				return &http.Response{}, err
			}

		} else {

			req, err = http.NewRequest(method, trigger.Url, nil)

			if err != nil {

				return &http.Response{}, err
			}
		}

	case "POST", "PUT":

		if len(trigger.Payload) > 0 {

			switch trigger.PayloadType {

			case "form data":

				formData := url.Values{}

				for _, field := range trigger.Payload {

					for k2, v2 := range trigger.CustomData {

						if field.FieldValue == k2 {

							formData.Add(field.FieldName, v2.(string))
						}
					}
				}

				body := formData.Encode()

				req, err = http.NewRequest(method, trigger.Url, strings.NewReader(body))

				if err != nil {

					return &http.Response{}, err
				}

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			case "json":

				mapData := map[string]interface{}{}

				for _, field := range trigger.Payload {

					for k2, v2 := range trigger.CustomData {

						if field.FieldValue == k2 {

							mapData[field.FieldName] = v2
						}
					}
				}

				jsonByte, err := json.Marshal(mapData)

				fmt.Printf("json body %v\n", string(jsonByte))

				if err != nil {

					return &http.Response{}, err
				}

				req, err = http.NewRequest(trigger.Method, trigger.Url, bytes.NewBuffer(jsonByte))

				fmt.Printf("req body: %v\n", req.Body)

				if err != nil {

					return &http.Response{}, err
				}

				req.Header.Set("Content-Type", "application/json")

			case "xml":

				var xmlData string

				for _, field := range trigger.Payload {

					for k2, v2 := range trigger.CustomData {

						if field.FieldValue == k2 {

							xmlData += fmt.Sprintf("<%s>%s</%s>", field.FieldName, v2, field.FieldName)
						}
					}
				}

				fmt.Printf("xml %s\n", xmlData)

				req, err = http.NewRequest(method, trigger.Url, strings.NewReader(xmlData))

				if err != nil {

					return &http.Response{}, err
				}

				req.Header.Set("Content-Type", "application/xml")

			}

		} else {

			req, err = http.NewRequest(method, trigger.Url, nil)

			if err != nil {

				return &http.Response{}, err
			}

		}

	}

	if trigger.Headers != nil {

		for _, header := range trigger.Headers {

			for k, v := range header.(map[string]interface{}) {

				req.Header.Set(k, v.(string))
			}
		}
	}

	resp, err := client.Do(req)

	if err != nil {

		return &http.Response{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		return &http.Response{}, errors.New("error response")
	}

	return resp, nil
}
