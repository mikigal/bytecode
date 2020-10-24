package bytecode

import (
	"fmt"
	"log"
)

const debugMode = false

func Ternary(expresion bool, whenTrue interface{}, whenFalse interface{}) interface{} {
	if expresion {
		return whenTrue
	} else {
		return whenFalse
	}
}

func Log(message string, values ...interface{}) {
	fmt.Printf(message + "\n", values...)
}

func Debug(message string, values ...interface{}) {
	if debugMode {
		fmt.Printf(message + "\n", values...)
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}