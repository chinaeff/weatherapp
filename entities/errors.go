package entities

import (
	"fmt"
)

func CommonError(msg string, err error) {
	if err != nil {
		fmt.Printf("%v: %v\n", msg, err)
		return
	}
}
