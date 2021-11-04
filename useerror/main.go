package main

import (
	"fmt"
	"javier/errors"
)

func main() {
	customError := errors.Set("Internal")
	fmt.Printf("customError type: %T\n", customError)
	isInternal := errors.Get(customError, "Internal")
	fmt.Printf("IsInternal: %v", isInternal)
}
