package json

import "encoding/json"

type Jsonify []byte

type Serialization interface {
	Deserialize() Scanner
}

type Scanner interface {
	Scan(out interface{})
}

func (j Jsonify) Scan(out interface{}) {
	_ = json.Unmarshal(j, out)
}
