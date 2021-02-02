package reflection

import (
	"log"
	"reflect"
	"testing"
)

func TestCast(t *testing.T) {
	var assertion []struct {
		Value  interface{}
		Expect interface{}
	}
	var assertExpect = Conversion{
		"URI":    "uri",
		"Params": "params",
	}

	assertion = append(assertion, struct {
		Value  interface{}
		Expect interface{}
	}{Value: map[string]interface{}{
		"URI":    "uri",
		"Params": "params",
	}, Expect: assertExpect})

	assertion = append(assertion, struct {
		Value  interface{}
		Expect interface{}
	}{Value: struct {
		URI    string
		Params interface{}
	}{
		URI:    "uri",
		Params: "params",
	}, Expect: assertExpect})

	assertion = append(assertion, struct {
		Value  interface{}
		Expect interface{}
	}{Value: &struct {
		URI    string
		Params interface{}
	}{
		URI:    "uri",
		Params: "params",
	}, Expect: assertExpect})

	// Assertion fail tests.
	//assertion = append(assertion, struct {
	//	Value  interface{}
	//	Expect interface{}
	//}{Value: 1, Expect: assertExpect})
	//assertion = append(assertion, struct {
	//	Value  interface{}
	//	Expect interface{}
	//}{Value: "1", Expect: assertExpect})

	for _, assert := range assertion {
		got := Cast(assert.Value)
		if !reflect.DeepEqual(got, assert.Expect) {
			t.Errorf("Expected value is: %v, but got: %v", assert.Expect, got)
			t.FailNow()
		}
	}

	log.Println("Everything looks fine.")

}
