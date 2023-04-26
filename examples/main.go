package main

import (
	"fmt"

	"github.com/pgaijin66/dea"
)

func main() {
	validEmail := "nawang81@gmail.com"
	invalidEmail := "salima9536@soombo.com"

	res1, err := dea.IsDisposableEmail(validEmail)
	if err != nil {
		fmt.Printf("%v", err)
	}

	res2, _ := dea.IsDisposableEmail(invalidEmail)

	// Check if email provided is not a disposable email address
	if res1 {
		fmt.Printf("%v is a disposable email address \n", validEmail)
	}

	// Check if email provided is  a disposable email address
	if res2 {
		fmt.Printf("%v is a disposable email address \n", invalidEmail)
	}

}
