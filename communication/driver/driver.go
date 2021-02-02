package driver

type Base struct {
	BaseURI     string      `json:"base_uri"`
	Parameters  Pair        `json:"parameters"`
	HTTPHeaders Pair        `json:"headers"`
	Response    interface{} `json:"response"`
}

type Driver interface {
	Serialize(params Pair)
	Send(method string) IHTTPResponseState
}
