package validator_test

import (
	"testing"
	"time"

	v "github.com/go-carrot/validator"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

// TestNullInt tests handling of a null.Int as the result
func TestNullInt(t *testing.T) {
	// Test success case
	var id null.Int
	err := v.Validate([]*v.Value{
		{Result: &id, Name: "id", Input: "12", TypeHandler: v.NullIntHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(12), id.Int64)
	assert.True(t, id.Valid)

	// Test null case
	var nullId null.Int
	err = v.Validate([]*v.Value{
		{Result: &nullId, Name: "id", Input: "", TypeHandler: v.NullIntHandler},
	})
	assert.False(t, nullId.Valid)

	// Test failure case
	var failureId null.Int
	err = v.Validate([]*v.Value{
		{Result: &failureId, Name: "id", Input: "12a", TypeHandler: v.NullIntHandler},
	})
	assert.NotNil(t, err)
	assert.False(t, failureId.Valid)
}

// TestNullString tests handling of a null.String as the result
func TestNullString(t *testing.T) {
	// Test success case
	var slug null.String
	err := v.Validate([]*v.Value{
		{Result: &slug, Name: "slug", Input: "hello", TypeHandler: v.NullStringHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, "hello", slug.String)
	assert.True(t, slug.Valid)

	// Test null case
	var nullSlug null.String
	err = v.Validate([]*v.Value{
		{Result: &nullSlug, Name: "slug", Input: "", TypeHandler: v.NullStringHandler},
	})
	assert.False(t, nullSlug.Valid)
}

// TestNullFloat tests handling of a null.Float as the result
func TestNullFloat(t *testing.T) {
	// Test success case
	var id null.Float
	err := v.Validate([]*v.Value{
		{Result: &id, Name: "id", Input: "12.8", TypeHandler: v.NullFloatHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(12.8), id.Float64)
	assert.True(t, id.Valid)

	// Test null case
	var nullId null.Float
	err = v.Validate([]*v.Value{
		{Result: &nullId, Name: "id", Input: "", TypeHandler: v.NullFloatHandler},
	})
	assert.False(t, nullId.Valid)

	// Test failure case
	var failureId null.Float
	err = v.Validate([]*v.Value{
		{Result: &failureId, Name: "id", Input: "12.8a", TypeHandler: v.NullFloatHandler},
	})
	assert.NotNil(t, err)
	assert.False(t, failureId.Valid)
}

// TestNullBool tests handling of a null.Bool as the result
func TestNullBool(t *testing.T) {
	// Test success case
	var someBool null.Bool
	err := v.Validate([]*v.Value{
		{Result: &someBool, Name: "some_bool", Input: "true", TypeHandler: v.NullBoolHandler},
	})
	assert.Nil(t, err)
	assert.True(t, someBool.Bool)
	assert.True(t, someBool.Valid)

	// Test null case
	var nullBool null.Bool
	err = v.Validate([]*v.Value{
		{Result: &nullBool, Name: "some_bool", Input: "", TypeHandler: v.NullBoolHandler},
	})
	assert.False(t, nullBool.Valid)

	// Test failure case
	var someOtherBool null.Bool
	err = v.Validate([]*v.Value{
		{Result: &someOtherBool, Name: "some_other_bool", Input: "12.8a", TypeHandler: v.NullBoolHandler},
	})
	assert.NotNil(t, err)
	assert.False(t, someOtherBool.Valid)
}

// TestNullTime tests handling of a null.Time as the result
func TestNullTime(t *testing.T) {
	// Test success case
	var successTime null.Time
	err := v.Validate([]*v.Value{
		{Result: &successTime, Name: "time", Input: "2012-11-01T22:08:41+00:00", TypeHandler: v.NullTimeHandler},
	})
	assert.Nil(t, err)
	assert.Equal(t, successTime.Time.Year(), 2012)
	assert.Equal(t, successTime.Time.Month(), time.November)
	assert.Equal(t, successTime.Time.Day(), 1)

	// Test null case
	var nullTime null.Time
	err = v.Validate([]*v.Value{
		{Result: &nullTime, Name: "time", Input: "", TypeHandler: v.NullTimeHandler},
	})
	assert.False(t, nullTime.Valid)

	// Test failure case
	var errorTime null.Time
	err = v.Validate([]*v.Value{
		{Result: &errorTime, Name: "time", Input: "abcd", TypeHandler: v.NullTimeHandler},
	})
	assert.NotNil(t, err)
}
