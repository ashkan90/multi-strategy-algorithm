package algorithms

import (
	"encoding/json"
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	json_util "github.com/ashkan90/multi-strategy-algorithm/communication/driver/utils/json"
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
	if e.before != nil {
		e.before(e)
	}

	smtp := driver.NewSMTP(e.URI, e.Headers)
	smtp.Serialize(e.Params)
	state := smtp.Send(driver.METHODPOST)

	if state.HasError() {
		log.Println("Something went wrong... Look at this! ", smtp.Response)
		return
	}

	if e.after != nil {
		e.after(e, smtp.Response)
	}

	log.Println("EmailStrategy -> Response -> ", smtp.Response)
}

func (e *Email) Before(handler strategy.BeforeHandler) {
	e.before = handler
}
func (e *Email) After(handler strategy.AfterHandler) {
	e.after = handler
}

func (e *Email) Deserialize() json_util.Scanner {
	var data, err = json.Marshal(e)
	if err != nil {
		return json_util.Jsonify(data)
	}

	return json_util.Jsonify(data)
}
