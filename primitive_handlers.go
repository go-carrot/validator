package validator

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func stringHandler(input string, value *Value) error {
	*value.Result.(*string) = input
	return nil
}

func float32Handler(input string, value *Value) error {
	res, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a float32"))
	}
	*value.Result.(*float32) = float32(res)
	return nil
}

func float64Handler(input string, value *Value) error {
	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a float64"))
	}
	*value.Result.(*float64) = float64(res)
	return nil
}

func boolHandler(input string, value *Value) error {
	res, err := strconv.ParseBool(input)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a bool"))
	}
	*value.Result.(*bool) = res
	return nil
}

func intHandler(input string, value *Value) error {
	res, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int"))
	}
	*value.Result.(*int) = int(res)
	return nil
}

func int8Handler(input string, value *Value) error {
	res, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int8"))
	}
	*value.Result.(*int8) = int8(res)
	return nil
}

func int16Handler(input string, value *Value) error {
	res, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int16"))
	}
	*value.Result.(*int16) = int16(res)
	return nil
}

func int32Handler(input string, value *Value) error {
	res, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int32"))
	}
	*value.Result.(*int32) = int32(res)
	return nil
}

func int64Handler(input string, value *Value) error {
	res, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int64"))
	}
	*value.Result.(*int64) = int64(res)
	return nil
}

func uintHandler(input string, value *Value) error {
	res, err := strconv.ParseUint(input, 10, 0)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a uint"))
	}
	*value.Result.(*uint) = uint(res)
	return nil
}

func uint8Handler(input string, value *Value) error {
	res, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a uint8"))
	}
	*value.Result.(*uint8) = uint8(res)
	return nil
}

func uint16Handler(input string, value *Value) error {
	res, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a uint16"))
	}
	*value.Result.(*uint16) = uint16(res)
	return nil
}

func uint32Handler(input string, value *Value) error {
	res, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a uint32"))
	}
	*value.Result.(*uint32) = uint32(res)
	return nil
}

func uint64Handler(input string, value *Value) error {
	res, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a uint64"))
	}
	*value.Result.(*uint64) = uint64(res)
	return nil
}

func timeHandler(input string, value *Value) error {
	res, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an RFC 3339 date-time (2006-01-02T15:04:05Z07:00)"))
	}
	*value.Result.(*time.Time) = res
	return nil
}

func invalidParam(name string, mustBe string) string {
	return fmt.Sprintf("Invalid `%v` parameter, `%v` must be %v", name, name, mustBe)
}
