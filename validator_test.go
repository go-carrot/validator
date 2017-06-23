package validator_test

import (
	"database/sql"
	"errors"
	"fmt"
	v "github.com/go-carrot/validator"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// IsSet is a rule that makes sure the value passed in isn't an empty string
func IsSet(name string, value string) error {
	if value == "" {
		return errors.New(fmt.Sprintf("Error, missing %v", name))
	}
	return nil
}

// MaxVal rule to test against ints
func MaxVal(maxValue int) v.Rule {
	return func(name string, input string) error {
		myval, _ := strconv.Atoi(input)
		if myval > maxValue {
			return errors.New(fmt.Sprintf("The value of %v may not be greater than %v", name, maxValue))
		}
		return nil
	}
}

// TestUnknownType tests handling a Cat, which is a type that this library
// knows nothing about (and will cause a panic)
func TestUnknownType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// Do nothing, just don't crash!
		}
	}()

	// Testing an unknown type
	type Cat struct{ name string }
	var myCat Cat
	v.Validate([]*v.Value{
		{Result: &myCat, Name: "cat", Input: "{ 'name': 'rae' }"},
	})

	// The Validate function should panic, and we never hit this line
	assert.Fail(t, "Line of code should be unreachable")
}

// TestCustomTypeHandler tests that we can create a new type handler
// and use it as expected
func TestCustomTypeHandler(t *testing.T) {
	// Create type handler
	var nullInt64TypeHandler = func(input string, value *v.Value) error {
		// Get int64
		res, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("Invalid parameter, must be an int64")
		}

		// Update nullInt
		nullInt := value.Result.(*sql.NullInt64)
		(*nullInt).Int64 = int64(res)
		(*nullInt).Valid = true
		return nil
	}

	// Test valid case
	var id sql.NullInt64
	err := v.Validate([]*v.Value{
		{Result: &id, Name: "id", Input: "42", TypeHandler: nullInt64TypeHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(42), id.Int64)
	assert.Equal(t, true, id.Valid)

	// Test empty case
	var emptyId sql.NullInt64
	err = v.Validate([]*v.Value{
		{Result: &emptyId, Name: "id", Input: "", TypeHandler: nullInt64TypeHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), emptyId.Int64)
	assert.Equal(t, false, emptyId.Valid)

	// Test error case
	var errorId sql.NullInt64
	err = v.Validate([]*v.Value{
		{Result: &errorId, Name: "id", Input: "abcd", TypeHandler: nullInt64TypeHandler},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), errorId.Int64)
	assert.Equal(t, false, errorId.Valid)
}
