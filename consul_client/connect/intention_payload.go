package connect

type IntentionPayload struct {
	SourceNS        string                 `json:"SourceNS"`
	SourceName      string                 `json:"SourceName"`
	DestinationNS   string                 `json:"DestinationNS"`
	DestinationName string                 `json:"DestinationName"`
	SourceType      string                 `json:"SourceType"`
	Action          string                 `json:"Action"`
	Meta            map[string]interface{} `json:"Meta"`
	Precedence      int                    `json:"Precedence"`
	CreateIndex     int                    `json:"CreateIndex"`
	ModifyIndex     int                    `json:"ModifyIndex"`
	Description     string                 `json:"Description"`
}

type IntentionResponse struct {
	ID string `json:"ID"`
	Allowed bool `json:"Allowed"`
}
