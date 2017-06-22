package validator

import (
	"errors"
	"fmt"
	"reflect"
)

// Value is the definition of a parameter that you would like to perform validation against.
type Value struct {
	Result      interface{}
	Default     string
	Name        string
	Input       string
	Rules       []Rule
	TypeHandler TypeHandler
}

// TypeHandler is a function that is responsible for
// are responsible for validating the input
// matches the type, and also to stuff the result into the *value.Result
type TypeHandler func(input string, value *Value) error

// Rule is a function that defines logic you would expect a Value to pass.
// The parameters passed in to this function are their respective
// values as they were set in the Value struct.
type Rule func(name string, input string) error

// Validate checks if an array of values passes their specified rules
func Validate(values []*Value) error {
	// Going through all values
	for _, value := range values {
		// Setting default, if value string isn't set
		resolvedInput := value.Input
		if resolvedInput == "" {
			resolvedInput = value.Default
		}

		// Going through all rules for each value
		for _, rule := range value.Rules {
			// Verifying rule passes
			err := rule(value.Name, resolvedInput)
			if err != nil {
				return err
			}
		}

		// Set primitive type handlers
		if value.TypeHandler == nil {
			err := applyTypeHandler(value)
			if err != nil {
				panic(err.Error())
			}
		}

		// Validate against type
		if value.Input != "" {
			err := value.TypeHandler(resolvedInput, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func applyTypeHandler(value *Value) error {
	switch i := (value.Result).(type) {
	default:
		return errors.New(fmt.Sprintf("go-carrot/validator cannot by default handle a Value with Result of type %v.  Must set a custom TypeHandler for %v.", reflect.TypeOf(i), value.Name))
	case *string:
		value.TypeHandler = stringHandler
	case *float32:
		value.TypeHandler = float32Handler
	case *float64:
		value.TypeHandler = float64Handler
	case *bool:
		value.TypeHandler = boolHandler
	case *int:
		value.TypeHandler = intHandler
	case *int8:
		value.TypeHandler = int8Handler
	case *int16:
		value.TypeHandler = int16Handler
	case *int32:
		value.TypeHandler = int32Handler
	case *int64:
		value.TypeHandler = int64Handler
	case *uint:
		value.TypeHandler = uintHandler
	case *uint8:
		value.TypeHandler = uint8Handler
	case *uint16:
		value.TypeHandler = uint16Handler
	case *uint32:
		value.TypeHandler = uint32Handler
	case *uint64:
		value.TypeHandler = uint64Handler
	}
	return nil
}
