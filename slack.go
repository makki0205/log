package log

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/makki0205/config"
)

var SlackURL = config.Env("slack_url")

func SendSlack(msg string) {
	payload := map[string]interface{}{
		"text": msg,
	}
	str, err := json.Marshal(payload)
	if err != nil {
		return
	}
	values := url.Values{}
	values.Set("payload", string(str))

	req, _ := http.NewRequest(
		"POST",
		SlackURL,
		strings.NewReader(values.Encode()),
	)

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	client.Do(req)
}

func SendSlackWithChan(msg, channel string) {
	payload := map[string]interface{}{
		"text": msg,
		"channel": channel,
	}
	str, err := json.Marshal(payload)
	if err != nil {
		return
	}
	values := url.Values{}
	values.Set("payload", string(str))

	req, _ := http.NewRequest(
		"POST",
		SlackURL,
		strings.NewReader(values.Encode()),
	)

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	client.Do(req)
}
