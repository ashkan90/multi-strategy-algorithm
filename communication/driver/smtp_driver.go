package driver

type SMTPDriver struct {
	Base
}

func NewSMTP(uri string, headers Pair) *SMTPDriver {
	return &SMTPDriver{
		Base: Base{
			BaseURI:     uri,
			HTTPHeaders: headers,
		},
	}
}

func (driver *SMTPDriver) URI() string {
	return driver.BaseURI
}

func (driver *SMTPDriver) Params() Pair {
	return driver.Parameters
}

func (driver *SMTPDriver) Headers() Pair {
	return driver.HTTPHeaders
}

func (driver *SMTPDriver) Serialize(params Pair) {
	driver.Parameters = params
}

func (driver *SMTPDriver) Send(method string) IHTTPResponseState {
	return Call(driver).Method(method).Do(&driver.Response)
}