package parser

import "encoding/json"

type RawJsonMap map[string](*json.RawMessage)

type JsonParser interface {
	GetAny(string) interface{}
	GetString(string) string
	GetInt64(string) int64
	GetFloat64(string) float64
	GetBool(string) bool
}

func UnmarshalToMap(data []byte) (RawJsonMap, error) {
	obj := RawJsonMap{}

	err := json.Unmarshal(data, &obj)
	return obj, err
}

func (this *RawJsonMap) GetAny(key string) (interface{}, error) {
	var val interface{}

	var deRef []byte = *((*this)[key])
	err := json.Unmarshal(deRef, &val)
	return val, err
}

func (this *RawJsonMap) GetString(key string) (string, error) {
	str, err := this.GetAny(key)
	return str.(string), err
}

func (this *RawJsonMap) GetInt64(key string) (int64, error) {
	str, err := this.GetAny(key)
	return str.(int64), err
}

func (this *RawJsonMap) GetFloat64(key string) (float64, error) {
	str, err := this.GetAny(key)
	return str.(float64), err
}

func (this *RawJsonMap) GetBool(key string) (bool, error) {
	str, err := this.GetAny(key)
	return str.(bool), err
}
