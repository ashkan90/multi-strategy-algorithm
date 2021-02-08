package reflection

import (
	"reflect"
)

type Conversion map[string]interface{}

func Cast(v interface{}) Conversion {
	source := reflect.ValueOf(v)

	if !source.IsValid() {
		return Conversion{}
	}

	return ConvertToMap(source)
}

func ConvertToMap(v reflect.Value) (cv Conversion) {
	cv = make(Conversion)

	switch v.Kind() {
	case reflect.Map:
		for _map := v.MapRange(); _map.Next(); {
			cv[_map.Key().String()] = _map.Value().Interface()
		}

	case reflect.Struct, reflect.Ptr:
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		for i := 0; i < v.NumField(); i += 1 {
			var field = v.Type().Field(i)
			var key = field.Name

			if js := field.Tag.Get("json"); js != "" {
				key = js
			}

			if valField := v.Field(i); valField.Kind() == reflect.Ptr && valField.IsNil() {
				continue
			}

			cv[key] = v.Field(i).Interface()

		}

	default:
		return
	}

	return
}
