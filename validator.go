package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Value is the definition of a parameter that you would like to perform validation against.
type Value struct {
	Result interface{}
	Name   string
	Input  string
	Rules  []Rule
}

// Rule is a function that defines logic you would expect a Value to pass.
// The parameters passed in to this function are their respective
// values as they were set in the Value struct.
type Rule func(name string, input string) error

// Validate checks if an array of values passes their specified rules
func Validate(values []*Value) error {
	// Going through all values
	for _, value := range values {

		// Going through all rules for each value
		for _, rule := range value.Rules {
			// Verifying rule passes
			err := rule(value.Name, value.Input)
			if err != nil {
				return err
			}
		}

		// Sticking the value of result into result
		if value.Input != "" {
			switch i := (value.Result).(type) {
			default:
				panic(fmt.Sprintf("go-carrot/validator cannot handle a Value with Result of type %v", reflect.TypeOf(i)))
			case *string:
				*i = value.Input
			case *float32:
				res, err := strconv.ParseFloat(value.Input, 32)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a float32"))
				}
				*i = float32(res)
			case *float64:
				res, err := strconv.ParseFloat(value.Input, 64)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a float64"))
				}
				*i = float64(res)
			case *bool:
				res, err := strconv.ParseBool(value.Input)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a bool"))
				}
				*i = res
			case *int:
				res, err := strconv.ParseInt(value.Input, 10, 0)
				if err != nil {
					return errors.New(invalidParam(value.Name, "an int"))
				}
				*i = int(res)
			case *int8:
				res, err := strconv.ParseInt(value.Input, 10, 8)
				if err != nil {
					return errors.New(invalidParam(value.Name, "an int8"))
				}
				*i = int8(res)
			case *int16:
				res, err := strconv.ParseInt(value.Input, 10, 16)
				if err != nil {
					return errors.New(invalidParam(value.Name, "an int16"))
				}
				*i = int16(res)
			case *int32:
				res, err := strconv.ParseInt(value.Input, 10, 32)
				if err != nil {
					return errors.New(invalidParam(value.Name, "an int32"))
				}
				*i = int32(res)
			case *int64:
				res, err := strconv.ParseInt(value.Input, 10, 64)
				if err != nil {
					return errors.New(invalidParam(value.Name, "an int64"))
				}
				*i = int64(res)
			case *uint:
				res, err := strconv.ParseUint(value.Input, 10, 0)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a uint"))
				}
				*i = uint(res)
			case *uint8:
				res, err := strconv.ParseUint(value.Input, 10, 8)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a uint8"))
				}
				*i = uint8(res)
			case *uint16:
				res, err := strconv.ParseUint(value.Input, 10, 16)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a uint16"))
				}
				*i = uint16(res)
			case *uint32:
				res, err := strconv.ParseUint(value.Input, 10, 32)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a uint32"))
				}
				*i = uint32(res)
			case *uint64:
				res, err := strconv.ParseUint(value.Input, 10, 64)
				if err != nil {
					return errors.New(invalidParam(value.Name, "a uint64"))
				}
				*i = uint64(res)
			}
		}
	}
	return nil
}

func invalidParam(name string, mustBe string) string {
	return fmt.Sprintf("Invalid `%v` parameter, `%v` must be %v", name, name, mustBe)
}
