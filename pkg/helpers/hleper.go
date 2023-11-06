package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-test/pkg/models"
	"net/http"
)

func Worker(ch <-chan map[string]interface{}) {
	var resp models.ResponceModels
	i := 1
	m := 1
	for key := range ch {

		resp.Event = key["ev"].(string)
		resp.EventType = key["et"].(string)
		resp.AppId = key["id"].(string)
		resp.UserId = key["uid"].(string)
		resp.MessageId = key["mid"].(string)
		resp.PageTitle = key["t"].(string)
		resp.PageUrl = key["p"].(string)
		resp.BrowserLanguage = key["l"].(string)
		resp.ScreenSize = key["sc"].(string)

		for j := 0; j < len(key); j++ {
			prevK := fmt.Sprintf("atrk%v", i)
			prevV := fmt.Sprintf("atrv%v", i)
			prevT := fmt.Sprintf("atrt%v", i)

			if value, ok := key[prevK]; ok {
				if _, isString := value.(string); isString {
					m := map[string]models.Extras{}
					m[key[prevK].(string)] = models.Extras{
						Value: key[prevV].(string),
						Type:  key[prevT].(string),
					}
					resp.Attributes = append(resp.Attributes, m)
				}
				i++
			}
		}

		for j := 0; j < len(key); j++ {
			k := fmt.Sprintf("uatrk%v", m)
			v := fmt.Sprintf("uatrv%v", m)
			t := fmt.Sprintf("uatrt%v", m)

			if value, ok := key[k]; ok {
				if _, isString := value.(string); isString {
					m := map[string]models.Extras{}
					m[key[k].(string)] = models.Extras{
						Value: key[v].(string),
						Type:  key[t].(string),
					}
					resp.Traits = append(resp.Traits, m)
				}
				m++
			}
		}

		sendToWebHook(resp)

	}

}

func sendToWebHook(resp models.ResponceModels) {
	url := "https://webhook.site/d47dbe70-81d2-44ff-9928-da3e02abf896"
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	webHookResp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer webHookResp.Body.Close()

	if webHookResp.Status == "200 OK" {
		println("Webhook request was successful.")
	} else {
		println("Webhook request failed with status:", webHookResp.Status)
	}

}
