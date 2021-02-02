package communication

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/algorithms"
	_strategy "github.com/ashkan90/multi-strategy-algorithm/communication/strategy"
	"testing"
)

func TestNewStrategy(t *testing.T) {
	smsStrategy := &algorithms.SMS{
		URI: "http://your.sms.provider",
		Params: map[string]interface{}{
			"messages": map[string]interface{}{},
		},
		Headers: map[string]interface{}{
			"Content-Type":  "application/json",
			"Authorization": "",
		},
	}
	strategy := _strategy.NewStrategy(smsStrategy)
	strategy.Send()

	emailStrategy := &algorithms.Email{
		URI: "http://your.custom.smtp_api.provider",
		Params: map[string]interface{}{
			//"key":   "value",
			//"other": "value",
		},
	}
	strategy.SetAlgorithm(emailStrategy)

	strategy.Send()
}
