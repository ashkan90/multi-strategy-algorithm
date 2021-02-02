package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReplicateHttp struct {
	to      string
	headers map[string]interface{}
	values  map[string]interface{}
}

func Init(to string) *ReplicateHttp {
	return &ReplicateHttp{
		to: to,
	}
}

func (rh *ReplicateHttp) Headers(h map[string]interface{}) *ReplicateHttp {
	rh.headers = h

	return rh
}

func (rh *ReplicateHttp) Values(v map[string]interface{}) *ReplicateHttp {
	rh.values = v

	return rh
}

func (rh *ReplicateHttp) Send() {}
func (rh *ReplicateHttp) SendScan(out interface{}) (err error) {
	err = rh.SendMethod("POST", out)
	return
}
func (rh *ReplicateHttp) SendMethod(method string, out interface{}) (err error) {

	data, err := json.Marshal(rh.values)
	if err != nil {
		return
	}

	req, err := http.NewRequest(method, rh.to, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	for h, v := range rh.headers {
		req.Header.Set(h, v.(string))
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return
	}

	return
}

