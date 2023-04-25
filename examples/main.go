package main

import (
	"fmt"

	"github.com/pgaijin66/dea"
)

func main() {
	validEmail := "demo@gmail.com"
	invalidEmail := "demo@2wc.info"

	res1, _ := dea.IsDisposableEmail(validEmail)
	res2, _ := dea.IsDisposableEmail(invalidEmail)

	// Check if email provided is not a disposable email address
	if !res1 {
		fmt.Printf("%v is a not disposable email address \n", validEmail)
	}

	// Check if email provided is  a disposable email address
	if res2 {
		fmt.Printf("%v is a disposable email address \n", invalidEmail)
	}

}
