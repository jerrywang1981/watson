package assistant

import (
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

// Watson assistant configuration
type WAConfig struct {
	ApiKey      string
	ApiUrl      string
	AssistantId string
	Version     string
}

// watson assistant bot
type WABot struct {
	config      *WAConfig
	service     *assistantv2.AssistantV2
	assistantId *string
	sessionId   *string
	UserId      *string
}

type Intent struct {
	Intent     string  `json:"intent"`
	Confidence float32 `json:"confidence"`
}

type Generic struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}

type WAOutput struct {
	Generic []Generic `json:"generic"`
	Intents []Intent  `json:"intents"`
}

type WAResult struct {
	Output WAOutput `json:"output"`
}
