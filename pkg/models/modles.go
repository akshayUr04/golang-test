package models

type ResponceModels struct {
	Event           string              `json:"event"`
	EventType       string              `json:"event_type"`
	AppId           string              `josn:"app_id"`
	UserId          string              `josn:"user_id"`
	MessageId       string              `josn:"message_id"`
	PageTitle       string              `josn:"page_title"`
	PageUrl         string              `josn:"page_url"`
	BrowserLanguage string              `josn:"browser_language"`
	ScreenSize      string              `josn:"screen_size"`
	Attributes      []map[string]Extras `josn:"attributes"`
	Traits          []map[string]Extras `josn:"traits"`
}

type Extras struct {
	Value string `json:"vakues"`
	Type  string `json:"type"`
}

// for k := range key {
// 	var i = 1
// 	prev := fmt.Sprintf("atrv%v", i)
// 	fmt.Println("prev:", prev)
// 	if strings.HasPrefix(k, prev) {
// 		fmt.Println("prevk:", k)
// 		prefixV := fmt.Sprintf("atrv %v", i)
// 		prefixT := fmt.Sprintf("atrt %v", i)

// 		fmt.Println("prefv", prefixV)
// 		fmt.Println("preft", prefixT)

// 		resp.Attributes[key[k].(string)] = models.Extras{
// 			Value: key[prefixV].(string),
// 			Type:  key[prefixT].(string),
// 		}
// 		i++
// 	}
// }
