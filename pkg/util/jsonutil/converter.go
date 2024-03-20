package jsonutil

import (
	"encoding/json"

	"github.com/maxzycon/rs-informasi-be/pkg/errors"
)

func ConvertToJSONString(i interface{}) (jsonString string, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		err = errors.ErrBadRequest
	}
	jsonString = string(b)
	return
}

func ConvertFromJSONSting(jsonString string, i interface{}) (err error) {
	err = json.Unmarshal([]byte(jsonString), i)
	if err != nil {
		err = errors.ErrBadRequest
	}
	return
}
