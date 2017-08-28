<a href="https://engineering.carrot.is/"><p align="center"><img src="https://cloud.githubusercontent.com/assets/2105067/24525319/d3d26516-1567-11e7-9506-7611b3287d53.png" alt="Go Carrot" width="350px" align="center;" /></p></a>
# Validator

[![Build Status](https://travis-ci.org/go-carrot/validator.svg?branch=master)](https://travis-ci.org/go-carrot/validator) [![codecov](https://codecov.io/gh/go-carrot/validator/branch/master/graph/badge.svg)](https://codecov.io/gh/go-carrot/validator) [![Go Report Card](https://goreportcard.com/badge/github.com/go-carrot/validator)](https://goreportcard.com/report/github.com/go-carrot/validator) [![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/go-carrot/validator)

Validator is a library that performs flexible string validation.

## Sample Usage

Before jumping into the details of the library, let's check out some usage:

```go
// Set the result values (The Input string gets parsed and stuffed into these)
var id int
var name string

// Run the validation
err := Validate([]*Value{
    {Result: &id, Name: "id", Input: "100", Rules: []Rule{MaxVal(10)}},
    {Result: &name, Name: "name", Input: "Brandon", Rules: []Rule{MaxLength(20)}},
})

// Check for any validation errors
//
// This error is thrown if one of the rules fail, or if the Input
// fails to be converted into it's Result type
if err != nil {
    // TODO, handle error (possibly HTTP 403)
    fmt.Println(err)
    return
}

// TODO, handle success - `id` and `name` are set at this point
```

> Note, the Rule implementations (MaxVal, MaxLength, etc.) aren't included in this library.  You'll either have to build them out yourself or check out [go-carrot/rules](https://github.com/go-carrot/rules) for some prebuilt ones.

## Values

Let's check out the Value struct.

```go
type Value struct {
    Result      interface{}
    Default     string
    Name        string
    Input       string
    Rules       []Rule
    TypeHandler TypeHandler
}
```

#### Result

Result must be a pointer to the variable you want to store the parsed input in.

By default, supported types for this are `*string`, `*float32`, `*float64`, `*bool`, `*int`, `*int8`, `*int16`, `*int32`, `*int64`, `*uint`, `*uint8`, `*uint16`, `*uint32`, `*uint64`.

For the default supported types, it is expected that the value of the `Input` parameter can be parsed into the decided type using their respective [strconv](https://golang.org/pkg/strconv/) function, else an error will be thrown by  [the Validate function](#the-validate-function) when it is called.

If you need to use another type, `TypeHandler` must also be set to the Value struct.

#### Default

This is the optional default value of `Input` that will be set, if the value of `Input` ends up being an empty string.

#### Name

Name should be a string that is available to all `Rule` functions.  This will be used to provide more user friendly error messaging.

#### Input

Input is the actual value that you would like to run validations against.  Because this library was built with validating HTTP requests in mind, this value must be a string.

#### Rules

This is a slice of rules that you require a particular value to pass.

This is optional, and can be not set if you don't have any rules for your value to pass.  The value will still go through the type check if the Input is a non-empty string.

#### TypeHandler

TypeHandler is a function that defines how the input string is parsed.

For basic types, it's not necessary to implement your own TypeHandler, as they have already been implemented and will be attached to Values automatically.

## Rules

A Rule is a very simple type of function:

```go
type Rule func(name string, input string) error
```

This function should throw an error in the event that the input does not match the criteria.

A simple rule can be implemented directly:

```go
func IsSet(name string, value string) error {
    if value == "" {
        return errors.New("Error, missing " + name)
    }
    return nil
}
```

You can have rules that take custom parameters by implementing a function that returns a rule:

```go
func MaxVal(maxValue int) Rule {
    return func(name string, input string) error {
        myval, _ := strconv.Atoi(input)
        if myval > maxValue {
            return errors.New(fmt.Sprintf("The value of %v may not be greater than %v", name, maxValue))
        }
        return nil
    }
}
```

Both of these strategies should feel very fluent in use:

```go
&Value{Result: &id, Name: "id", Input: "100", Rules: []Rule{IsSet, MaxVal(10)}},
```

> You won't find any prebuilt rules in [go-carrot/validator](https://github.com/go-carrot/validator).  If you're looking for those check out the [go-carrot/rules](https://github.com/go-carrot/rules) repository.

## TypeHandlers

A TypeHandler is a function that follows the following definition:

```go
type TypeHandler func(input string, value *Value) error
```

A TypeHandler is responsible for:

- Validation of the string input
- Converting string input to the desired type
- Passing the converted type into the `value.Result`

TypeHandlers are best explained by example.  This is a TypeHandler for a `*null.Int`

```go
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
```

## The Validate Function

The validate function is the function that will actually perform your input validation.  This function will throw an error if any of your values fail validation.

```go
func Validate(values []*Value) error
```

The easiest way to call this validate function is to simply inline the `[]*Value` parameter, as displayed below:

```go
err := Validate([]*Value{
    {Result: &id, Name: "id", Input: "100", Rules: []Rule{IsSet, MaxVal(10)}},
    {Result: &name, Name: "name", Input: "Brandon", Rules: []Rule{IsSet, MaxLength(20)}},
})
```

## License

[MIT](LICENSE.md)
