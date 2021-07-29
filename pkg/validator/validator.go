package validator

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

type validator interface {
	Lt(mark string) string
	Le(mark string) string
	Eq(mark string) string
	Ne(mark string) string
	Ge(mark string) string
	Gt(mark string) string
	NotEmpty() string
}

type Validator struct{}

var _ validator = (*Validator)(nil)

func NewValidator() Validator {
	return Validator{}
}

// Lt Less than the input parameter (<)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Lt(mark string) string {
	return "lt=" + mark
}

// Le is less than or equal to the input parameter (<=)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Le(mark string) string {
	return "le=" + mark
}

// Eq Equal to input (==)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Eq(mark string) string {
	return "eq=" + mark
}

// Ne Not equal to input (!=)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Ne(mark string) string {
	return "ne=" + mark
}

// Ge Greater than or equal to input parameters (>=)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Ge(mark string) string {
	return "ge=" + mark
}

// Gt Greater than the input parameter (>)
// If it is string, array or slice, it is length comparison.
// If it is int, uint or float, it is numeric comparison.
func (v Validator) Gt(mark string) string {
	return "gt=" + mark
}

// NotEmpty not empty
func (v Validator) NotEmpty() string {
	return "notEmpty"
}

// Verify ...
func Verify(st interface{}, ruleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if len(ruleMap[tagVal.Name]) > 0 {
			for _, v := range ruleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Tag.Get("json") + " value cannot be empty")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Tag.Get("json") + " length or value does not fit the range")
					}
				}
			}
		}
	}
	return nil
}

// compareVerify Length and number verification method Automatic verification according to type
func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

// isBlank non empty check
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// compare used to compare two values
func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	log.Println("val.kind: ", val.Kind())
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}
