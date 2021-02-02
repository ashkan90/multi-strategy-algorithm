package driver

import "log"

type SMSDriver struct {
	Base
}

func NewSMS(uri string, headers Pair) *SMSDriver {
	return &SMSDriver{
		Base: Base{
			BaseURI:     uri,
			HTTPHeaders: headers,
		},
	}
}

func (driver *SMSDriver) URI() string {
	return driver.BaseURI
}

func (driver *SMSDriver) Params() Pair {
	return driver.Parameters
}

func (driver *SMSDriver) Headers() Pair {
	log.Println("SMSDriver -> Headers ->", driver.HTTPHeaders)
	return driver.HTTPHeaders
}

func (driver *SMSDriver) Serialize(params Pair) {
	driver.Parameters = params
}

func (driver *SMSDriver) Send(method string) IHTTPResponseState {
	return Call(driver).Method(method).Do(&driver.Response)
}