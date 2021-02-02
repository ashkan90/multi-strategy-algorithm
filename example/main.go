package main

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/algorithms"
	_strategy "github.com/ashkan90/multi-strategy-algorithm/communication/strategy"
)

// You can specify your parameters inside of Params key.
// So with that, you would be able to integrate with
// your custom sms provider.
// There's a bit of boilerplate code that you've to write your own
// but these algorithms holding same code base in this case
// you can debug code much more easily.
// ...
// Even you can make your basic Struct to adopt it here as Params.
func main() {

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
