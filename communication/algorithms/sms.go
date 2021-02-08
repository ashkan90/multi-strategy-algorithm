package algorithms

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	"github.com/ashkan90/multi-strategy-algorithm/communication/strategy"
	"log"
)

type SMS struct {
	URI     string
	Params  driver.Pair
	Headers driver.Pair

	before strategy.BeforeHandler
	after  strategy.AfterHandler
}

func (s *SMS) Send() {
	s.before(s)

	dsms := driver.NewSMS(s.URI, s.Headers)
	dsms.Serialize(s.Params)
	state := dsms.Send(driver.METHODPOST)

	if state.HasError() {
		log.Println("SMSStrategy -> Something went wrong... Look at this! ", dsms.Response)
		return
	}

	s.after(dsms.Response)
	log.Println("SMSStrategy -> Response -> ", dsms.Response)
}

func (s *SMS) Before(handler strategy.BeforeHandler) {
	if handler == nil {
		handler = func(algorithm strategy.Algorithm) {}
	}
	s.before = handler
}
func (s *SMS) After(handler strategy.AfterHandler) {
	if handler == nil {
		handler = func(response interface{}) {}
	}
	s.after = handler
}
