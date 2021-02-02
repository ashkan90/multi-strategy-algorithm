package algorithms

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	"log"
)

type SMS struct {
	URI     string
	Params  driver.Pair
	Headers driver.Pair
}

func (s *SMS) Send() {
	dsms := driver.NewSMS(s.URI, s.Headers)
	dsms.Serialize(s.Params)
	state := dsms.Send(driver.METHODPOST)

	if state.HasError() {
		log.Println("SMSStrategy -> Something went wrong... Look at this! ", dsms.Response)
		return
	}

	log.Println("SMSStrategy -> Response -> ", dsms.Response)
}
