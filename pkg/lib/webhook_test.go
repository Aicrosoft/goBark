package lib

import (
	"testing"

	"github.com/Aicrosoft/goBark/internal/setting"
)

func TestBuildReqURL(t *testing.T) {
	w := GetWebhook(&setting.AppSetting{
		Webhook: setting.WebhookSetting{
			Enabled: true,
			URL:     "https://api.day.app/amF/{{.Title}}/{{.Content}}",
		}})

	ret, err := w.buildReqURL(&setting.MessageSetting{Title: "Test title", Content: "Send OR Request Content."})
	if err != nil {
		t.Error(err)
	}

	expected := "https://api.day.app/amF/Test title/Send OR Request Content."
	if ret != expected {
		t.Errorf("expected %s, got %s", expected, ret)
	}

	t.Log(ret)
}

func TestBuildReqBody(t *testing.T) {
	t.Skip() //only one method can be used.
	w := GetWebhook(&setting.AppSetting{
		Webhook: setting.WebhookSetting{
			Enabled:     true,
			URL:         "https://api.day.app/amF",
			RequestBody: "{ \"title\": \"{{.Title}}\", \"content\": \"{{.Content}}\", \"value\": \"{{.Value}}\" }",
		}})

	ret, err := w.buildReqBody(&setting.MessageSetting{Title: "Test title", Content: "Send OR Request Content.", Value: "ValueOfValue"})
	if err != nil {
		t.Error(err)
	}

	expected := `{ "title": "Test title", "content": "Send OR Request Content.", "value": "ValueOfValue" }`
	if ret != expected {
		t.Errorf("expected %s, got %s", expected, ret)
	}

	t.Log(ret)
}
