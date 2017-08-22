package validator

import (
	"errors"
	"strconv"
	"time"

	"gopkg.in/guregu/null.v3"
)

// NullIntHandler is a TypeHandler for null.Int
func NullIntHandler(input string, value *Value) error {
	// Cast
	nullInt := value.Result.(*null.Int)

	// Check for empty
	if len(input) == 0 {
		(*nullInt).Valid = false
		return nil
	}

	// Get int64
	res, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an int64"))
	}

	// Update null.Int
	(*nullInt).Int64 = int64(res)
	(*nullInt).Valid = true
	return nil
}

// NullStringHandler is a TypeHandler for null.String
func NullStringHandler(input string, value *Value) error {
	nullString := value.Result.(*null.String)
	(*nullString).String = input
	(*nullString).Valid = len(input) != 0
	return nil
}

// NullFloatHandler is a TypeHandler for null.Float
func NullFloatHandler(input string, value *Value) error {
	// Cast
	nullFloat := value.Result.(*null.Float)

	// Check for empty
	if len(input) == 0 {
		(*nullFloat).Valid = false
		return nil
	}

	// Get float64
	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a float64"))
	}

	// Update null.Float
	(*nullFloat).Float64 = res
	(*nullFloat).Valid = true
	return nil
}

// NullBoolHandler is a TypeHandler for null.Bool
func NullBoolHandler(input string, value *Value) error {
	// Cast
	nullBool := value.Result.(*null.Bool)

	// Check for empty
	if len(input) == 0 {
		(*nullBool).Valid = false
		return nil
	}

	// Get bool
	res, err := strconv.ParseBool(input)
	if err != nil {
		return errors.New(invalidParam(value.Name, "a bool"))
	}

	// Update null.Bool
	(*nullBool).Bool = res
	(*nullBool).Valid = true
	return nil
}

// NullTimeHandler is a TypeHandler for null.Time
func NullTimeHandler(input string, value *Value) error {
	// Cast
	nullTime := value.Result.(*null.Time)

	// Check for empty
	if len(input) == 0 {
		(*nullTime).Valid = false
		return nil
	}

	// Get time.Time
	res, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return errors.New(invalidParam(value.Name, "an RFC 3339 date-time (2006-01-02T15:04:05Z07:00)"))
	}

	// Update null.Time
	(*nullTime).Time = res
	(*nullTime).Valid = true
	return nil
}
