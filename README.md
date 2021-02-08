Multi strategy algorithm explains how multiple algorithm works together within same code-base.
<hr>

## Example Usage

```go

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
  
```
TODO
-[] cli üzerinden strateji oluşturulacak.