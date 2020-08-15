package assistant

import (
	"encoding/json"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

// global watson bot
var Bot *WABot

const API_VERSION = "2020-04-01"

func ConnectWA(c *WAConfig) *WABot {
	Bot = NewSession(c)
	return Bot
}

func NewSession(c *WAConfig) *WABot {
	if c.Version == "" {
		c.Version = API_VERSION
	}
	bot := &WABot{
		config: c,
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: c.ApiKey,
	}

	service, serviceErr := assistantv2.
		NewAssistantV2(&assistantv2.AssistantV2Options{
			URL:           c.ApiUrl,
			Version:       c.Version,
			Authenticator: authenticator,
		})

		// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}
	bot.service = service

	assistantID := c.AssistantId
	// Call the assistant CreateSession method
	createSessionResult, _, responseErr := service.
		CreateSession(&assistantv2.CreateSessionOptions{
			AssistantID: core.StringPtr(assistantID),
		})

	if responseErr != nil {
		panic(responseErr)
	}
	sessionID := createSessionResult.SessionID

	bot.sessionId = sessionID
	bot.assistantId = core.StringPtr(assistantID)
	if bot.UserId == nil {
		bot.UserId = core.StringPtr("dummy")
	}
	return bot
}

func (w *WABot) Send(msg string) *WAResult {
	_, response, responseErr := w.service.
		Message(&assistantv2.MessageOptions{
			AssistantID: w.assistantId,
			SessionID:   w.sessionId,
			Input: &assistantv2.MessageInput{
				Text: core.StringPtr(msg),
			},
			Context: &assistantv2.MessageContext{
				Global: &assistantv2.MessageContextGlobal{
					System: &assistantv2.MessageContextGlobalSystem{
						UserID: w.UserId,
					},
				},
			},
		})

	if responseErr != nil {
		panic(responseErr)
	}

	oRes := &WAResult{}
	output, err := json.Marshal(response.GetResult())
	if err != nil {
		return oRes
	}
	json.Unmarshal(output, &oRes)
	return oRes
}

func (w *WABot) Close() {
	if w != nil && w.service != nil {
		_, responseErr := w.service.
			DeleteSession(&assistantv2.DeleteSessionOptions{
				AssistantID: w.assistantId,
				SessionID:   w.sessionId,
			})

		if responseErr != nil {
			panic(responseErr)
		}
	}
	w.service = nil
	w.sessionId = nil
	w.assistantId = nil
	w.UserId = nil
}
