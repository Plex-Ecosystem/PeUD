package v1

import (
	"errors"
	"fmt"
)

type InternalServerError struct {
	Code     int
	Function string
	Error    string
}

// from https://stackoverflow.com/a/37214476/12204515
type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" || asString == "null" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

func (bit *ConvertibleBoolean) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		switch src.(int64) {
		case 0:
			*bit = false
		case 1:
			*bit = true
		default:
			return errors.New("incompatible type for ConvertibleBoolean")
		}
	}
	return nil
}
