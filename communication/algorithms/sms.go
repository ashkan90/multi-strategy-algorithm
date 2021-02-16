package algorithms

import (
	"encoding/json"
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver"
	json_util "github.com/ashkan90/multi-strategy-algorithm/communication/driver/utils/json"
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
	if s.before != nil {
		s.before(s)
	}

	dsms := driver.NewSMS(s.URI, s.Headers)
	dsms.Serialize(s.Params)
	state := dsms.Send(driver.METHODPOST)

	if state.HasError() {
		log.Println("SMSStrategy -> Something went wrong... Look at this! ", dsms.Response)
		return
	}

	if s.after != nil {
		s.after(s, dsms.Response)
	}
	log.Println("SMSStrategy -> Response -> ", dsms.Response)
}

func (s *SMS) Before(handler strategy.BeforeHandler) {
	s.before = handler
}
func (s *SMS) After(handler strategy.AfterHandler) {
	s.after = handler
}

func (s *SMS) Deserialize() json_util.Scanner {
	var data, err = json.Marshal(s)
	if err != nil {
		return json_util.Jsonify(data)
	}

	return json_util.Jsonify(data)
}
