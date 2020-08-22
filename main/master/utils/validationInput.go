package utils

import (
	"errors"
)

func ValidateInputNotNil(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("Data Input Cannot Empty")
		case 0:
			return errors.New("Data Input Cannot 0")
		case nil:
			return errors.New("Input cannot be nil")
		}
	}
	return nil
}

func ValidateInputLenCharacter(min, max int, data ...interface{}) error {
	for _, value := range data {
		if len(value.(string)) >= min && len(value.(string)) <= max {
			return errors.New("Input cannot more 20 character")
		}
	}
	return nil
}
