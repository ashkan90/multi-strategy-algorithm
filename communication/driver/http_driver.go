package driver

import (
	"encoding/json"
	"github.com/ashkan90/multi-strategy-algorithm/communication/driver/utils/reflection"
	"github.com/ashkan90/multi-strategy-algorithm/http"
)

const (
	METHODGET    = "GET"
	METHODPOST   = "POST"
	METHODPUT    = "PUT"
	METHODPATCH  = "PATCH"
	METHODDELETE = "DELETE"
)

type Pair map[string]interface{}

type IHTTPDriver interface {
	URI() string
	Params() Pair
	Headers() Pair
}

type IHTTPResponseState interface {
	HasError() bool
}

type HTTPDriver struct {
	calle    IHTTPDriver
	method   string
	hasError bool
}

func (p *Pair) Append(val interface{}) {
	if val == nil {
		return
	}

	var newPair = CastToPair(val)
	for key, value := range newPair {
		if value != nil {
			map[string]interface{}(*p)[key] = value
		}
	}
}

func (driver *HTTPDriver) Method(m string) *HTTPDriver {
	driver.method = m
	return driver
}

func (driver *HTTPDriver) Do(response interface{}) IHTTPResponseState {
	err := http.Init(driver.calle.URI()).
		Headers(driver.calle.Headers()).
		Values(driver.calle.Params()).
		SendMethod(driver.method, response)

	if err != nil {
		driver.hasError = true
		b, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})

		_ = json.Unmarshal(b, response)
	}

	return driver
}

func (driver *HTTPDriver) HasError() bool {
	return driver.hasError
}

func Call(calle IHTTPDriver) *HTTPDriver {
	return &HTTPDriver{calle: calle, method: "POST"}
}

func CastToPair(val interface{}) Pair {
	return Pair(reflection.Cast(val))
}
