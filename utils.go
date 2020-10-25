package bytecode

import (
	"fmt"
	"reflect"
)

func countParameters(descriptor string) uint16 {
	parameters := 0
	foundReference := false

	for _, char := range descriptor {
		if char == '(' || char == '[' {
			continue
		}

		if char == ')' {
			break
		}

		if char == 'L' {
			foundReference = true
		}

		if foundReference {
			if char == ';' {
				foundReference = false
				parameters++
			}

			continue
		}

		parameters++
	}

	return uint16(parameters)
}

func AsJavaType(value interface{}) JavaType {
	if _, ok := value.(int); ok {
		return Type_Char
	}

	switch value.(type) {
	case int8:
		return Type_Byte
	case int16:
		return Type_Short
	case int:
		return Type_Int
	case int32:
		return Type_Int
	case int64:
		return Type_Long
	case float32:
		return Type_Float
	case float64:
		return Type_Double
	case bool:
		return Type_Boolean
	case string:
		return Type_Reference
	}

	panic("could not get JavaType of " + reflect.TypeOf(value).Name())
}

func Ternary(expresion bool, whenTrue interface{}, whenFalse interface{}) interface{} {
	if expresion {
		return whenTrue
	} else {
		return whenFalse
	}
}

func Log(message string, values ...interface{}) {
	fmt.Printf(message+"\n", values...)
}
