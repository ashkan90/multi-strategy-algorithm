package algorithms

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	"github.com/ashkan90/multi-strategy-algorithm/communication/strategy"
	"log"
)

type Email struct {
	URI     string
	Params  driver.Pair
	Headers driver.Pair
	before  strategy.BeforeHandler
	after   strategy.AfterHandler
}

func (e *Email) Send() {
	e.before(e)

	smtp := driver.NewSMTP(e.URI, e.Headers)
	smtp.Serialize(e.Params)
	state := smtp.Send(driver.METHODPOST)

	if state.HasError() {
		log.Println("Something went wrong... Look at this! ", smtp.Response)
		return
	}

	e.after(smtp.Response)

	log.Println("EmailStrategy -> Response -> ", smtp.Response)
}

func (e *Email) Before(handler strategy.BeforeHandler) {
	if handler == nil {
		handler = func(algorithm strategy.Algorithm) {}
	}
	e.before = handler
}
func (e *Email) After(handler strategy.AfterHandler) {
	if handler == nil {
		handler = func(response interface{}) {}
	}
	e.after = handler
}
