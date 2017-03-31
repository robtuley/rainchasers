package request

import (
	"encoding/json"
	"errors"
)

func ParseFloat(raw json.RawMessage) (float32, error) {
	var f float32
	var a []float32

	err := json.Unmarshal(raw, &f)
	if err != nil {
		err = json.Unmarshal(raw, &a)
		if err == nil {
			if len(a) > 0 {
				f = a[0]
			} else {
				err = errors.New("empty array")
			}
		}
	}
	return f, err
}

func ParseString(raw json.RawMessage) (string, error) {
	var s string
	var a []string

	err := json.Unmarshal(raw, &s)
	if err != nil {
		err = json.Unmarshal(raw, &a)
		if err == nil {
			if len(a) > 0 {
				s = a[0]
			} else {
				err = errors.New("empty array")
			}
		}
	}
	return s, err
}
