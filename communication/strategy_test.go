package communication

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/algorithms"
	_strategy "github.com/ashkan90/multi-strategy-algorithm/communication/strategy"
	"log"
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
		URI: "https://api.sendinblue.com/v3/smtp/email",
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
			"api-key":      "",
		},
		Params: map[string]interface{}{
			"sender": map[string]interface{}{
				"name":  "sender name",
				"email": "info@yourdomain.com",
			},
			"to": []map[string]interface{}{
				{
					"email": "your@test.com",
					"name":  "Test Surname",
				},
			},
			"subject":     "Hello world",
			"htmlContent": "<html><head></head><body><p>Hello,</p>This is my first transactional email sent from Sendinblue.</p></body></html>",
		},
	}
	strategy.SetAlgorithm(emailStrategy)

	strategy.Send()
}

func TestEmailStrategy(t *testing.T) {
	emailStrategy := &algorithms.Email{
		URI: "https://api.sendinblue.com/v3/smtp/email",
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
			"api-key":      "",
		},
		Params: map[string]interface{}{
			"sender": map[string]interface{}{
				"name":  "sender name",
				"email": "info@yourdomain.com",
			},
			"to": []map[string]interface{}{
				{
					"email": "your@test.com",
					"name":  "Test Surname",
				},
			},
			"templateId": 1,
			//"subject":"Hello world",
			//"htmlContent":"<html><head></head><body><p>Hello,</p>This is my first transactional email sent from Sendinblue.</p></body></html>",
		},
	}
	strategy := _strategy.NewStrategy(emailStrategy)

	strategy.Before(nil)
	strategy.After(AfterHandler)

	strategy.Send()
}

func BeforeHandler(algorithm _strategy.Algorithm) {
	log.Println("BeforeHandler -> Incoming Algorithm ->", algorithm)
}

func AfterHandler(algorithm _strategy.Algorithm, response interface{}) {
	log.Println("AfterHandler -> Incoming Response ->", response)
}
