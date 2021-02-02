package algorithms

import (
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	"log"
)

type Email struct {
	URI     string
	Params  driver.Pair
	Headers driver.Pair
}

func (e *Email) Send() {
	smtp := driver.NewSMTP(e.URI, e.Headers)
	smtp.Serialize(e.Params)
	state := smtp.Send(driver.METHODGET)

	if state.HasError() {
		log.Println("Something went wrong... Look at this! ", smtp.Response)
		return
	}

	log.Println("EmailStrategy -> Response -> ", smtp.Response)
}
