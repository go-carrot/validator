package validator_test

import (
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
		&v.Value{Result: &myCat, Name: "cat", Input: "{ 'name': 'rae' }"},
	})

	// The Validate function should panic, and we never hit this line
	assert.Fail(t, "Line of code should be unreachable")
}

// TestString tests handling of a string as the result
func TestString(t *testing.T) {
	// Test success case
	var successId string
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, "20", successId)

	// Test empty case
	var emptyId string
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, "", emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId string
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, "", emptyIsSetId)
}

// TestFloat32 tests handing of a float32 as the result
func TestFloat32(t *testing.T) {
	// Test success case
	var successId float32
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20.2"},
	})
	assert.Nil(t, err)
	assert.Equal(t, float32(20.2), successId)

	// Test another success case
	var anotherSuccessId float32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &anotherSuccessId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, float32(20.0), anotherSuccessId)

	// Test empty case
	var emptyId float32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, float32(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId float32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, float32(0), emptyIsSetId)

	// Test parse failure case
	var parseFailureId float32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &parseFailureId, Name: "id", Input: "20a", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, float32(0), parseFailureId)
}

// TestFloat64 tests handing of a float64 as the result
func TestFloat64(t *testing.T) {
	// Test success case
	var successId float64
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20.2"},
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(20.2), successId)

	// Test another success case
	var anotherSuccessId float64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &anotherSuccessId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(20.0), anotherSuccessId)

	// Test empty case
	var emptyId float64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId float64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), emptyIsSetId)

	// Test parse failure case
	var parseFailureId float64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &parseFailureId, Name: "id", Input: "20a", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), parseFailureId)
}

// TestBool tests handing of a bool as the result
func TestBool(t *testing.T) {
	// Test success case
	var successBool bool
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successBool, Name: "bool", Input: "true"},
	})
	assert.Nil(t, err)
	assert.Equal(t, true, successBool)

	// Test empty case
	var emptyBool bool
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyBool, Name: "bool", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, false, emptyBool)

	// Test empty case with IsSet rule
	var emptyIsSetBool bool
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetBool, Name: "bool", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, false, emptyIsSetBool)

	// Test failure case
	var failureBool bool
	err = v.Validate([]*v.Value{
		&v.Value{Result: &failureBool, Name: "bool", Input: "not-a-bool"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, false, failureBool)
}

// TestInt tests handing of an int as the result
func TestInt(t *testing.T) {
	// Test success case
	var successId int
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, int(20), successId)

	// Test empty case
	var emptyId int
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, int(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId int
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int(0), emptyIsSetId)

	// Test overflow case
	var overflowId int
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "9223372036854775808"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int(0), overflowId)

	// Test string case
	var stringId int
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int(0), stringId)
}

// TestInt8 tests handing of an int8 as the result
func TestInt8(t *testing.T) {
	// Test success case
	var successId int8
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, int8(20), successId)

	// Test empty case
	var emptyId int8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, int8(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId int8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int8(0), emptyIsSetId)

	// Test overflow case
	var overflowId int8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "128"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int8(0), overflowId)

	// Test string case
	var stringId int8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int8(0), stringId)
}

// TestInt16 tests handing of an int16 as the result
func TestInt16(t *testing.T) {
	// Test success case
	var successId int16
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, int16(20), successId)

	// Test empty case
	var emptyId int16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, int16(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId int16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int16(0), emptyIsSetId)

	// Test overflow case
	var overflowId int16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "32768"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int16(0), overflowId)

	// Test string case
	var stringId int16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int16(0), stringId)
}

// TestInt32 tests handling of an int32 as the result
func TestInt32(t *testing.T) {
	// Test success case
	var successId int32
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, int32(20), successId)

	// Test empty case
	var emptyId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, int32(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), emptyIsSetId)

	// Test successful case with MaxVal rule
	var maxValPassId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &maxValPassId, Name: "id", Input: "100", Rules: []v.Rule{MaxVal(100)}},
	})
	assert.Nil(t, err)
	assert.Equal(t, int32(100), maxValPassId)

	// Test failure case with MaxVal rule
	var maxValFailureId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &maxValFailureId, Name: "id", Input: "101", Rules: []v.Rule{MaxVal(100)}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), maxValFailureId)

	// Test string case
	var stringId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "hello world"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), stringId)

	// Test float case
	var floatId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &floatId, Name: "id", Input: "20.1"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), floatId)

	// Test overflow case
	var overflowId int32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "2147483648"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), overflowId)
}

// TestInt64 tests handing of an int64 as the result
func TestInt64(t *testing.T) {
	// Test success case
	var successId int64
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(20), successId)

	// Test empty case
	var emptyId int64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId int64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), emptyIsSetId)

	// Test overflow case
	var overflowId int64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "9223372036854775808"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), overflowId)

	// Test string case
	var stringId int64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), stringId)
}

// TestUint tests handing of a uint as the result
func TestUint(t *testing.T) {
	// Test success case
	var successId uint
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint(20), successId)

	// Test empty case
	var emptyId uint
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId uint
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), emptyIsSetId)

	// Test overflow case
	var overflowId uint
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "18446744073709551616"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), overflowId)

	// Test string case
	var stringId uint
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), stringId)
}

// TestUint8 tests handing of an uint8 as the result
func TestUint8(t *testing.T) {
	// Test success case
	var successId uint8
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint8(20), successId)

	// Test empty case
	var emptyId uint8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint8(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId uint8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint8(0), emptyIsSetId)

	// Test overflow case
	var overflowId uint8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "256"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint8(0), overflowId)

	// Test string case
	var stringId uint8
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint8(0), stringId)
}

// TestUint16 tests handing of an uint16 as the result
func TestUint16(t *testing.T) {
	// Test success case
	var successId uint16
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint16(20), successId)

	// Test empty case
	var emptyId uint16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint16(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId uint16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint16(0), emptyIsSetId)

	// Test overflow case
	var overflowId uint16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "65536"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint16(0), overflowId)

	// Test string case
	var stringId uint16
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint16(0), stringId)
}

// TestUint32 tests handing of an uint32 as the result
func TestUint32(t *testing.T) {
	// Test success case
	var successId uint32
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint32(20), successId)

	// Test empty case
	var emptyId uint32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId uint32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint32(0), emptyIsSetId)

	// Test overflow case
	var overflowId uint32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "4294967296"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint32(0), overflowId)

	// Test string case
	var stringId uint32
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint32(0), stringId)
}

// TestUint64 tests handing of an uint64 as the result
func TestUint64(t *testing.T) {
	// Test success case
	var successId uint64
	err := v.Validate([]*v.Value{
		&v.Value{Result: &successId, Name: "id", Input: "20"},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint64(20), successId)

	// Test empty case
	var emptyId uint64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyId, Name: "id", Input: ""},
	})
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), emptyId)

	// Test empty case with IsSet rule
	var emptyIsSetId uint64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &emptyIsSetId, Name: "id", Input: "", Rules: []v.Rule{IsSet}},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), emptyIsSetId)

	// Test overflow case
	var overflowId uint64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &overflowId, Name: "id", Input: "18446744073709551616"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), overflowId)

	// Test string case
	var stringId uint64
	err = v.Validate([]*v.Value{
		&v.Value{Result: &stringId, Name: "id", Input: "Hello World"},
	})
	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), stringId)
}
