package utils

import (
	"fmt"
)

func LogError(err *error) {
	fmt.Println("Error:", *err)
}

func LogWarn(err *error) {
	fmt.Println("Error:", *err)
}
