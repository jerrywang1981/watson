package assistant

import "testing"

func TestAssistant(t *testing.T) {
	bot := NewSession(&WAConfig{
		ApiKey:      "your key",
		ApiUrl:      "your url",
		AssistantId: "your assistant id",
		Version:     "2020-04-01",
	})
	defer bot.Close()
	msg := "your message"
	output := bot.Send(msg)
	t.Logf("%+v", output)
}
