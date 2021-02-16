package strategy

import "github.com/ashkan90/multi-strategy-algorithm/communication/driver/utils/json"

type BeforeHandler func(algorithm Algorithm)
type AfterHandler func(algorithm Algorithm, response interface{})

type Algorithm interface {
	Send()
	Before(handler BeforeHandler)
	After(handler AfterHandler)

	json.Serialization
}
