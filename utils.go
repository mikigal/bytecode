package bytecode

import (
	"fmt"
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
