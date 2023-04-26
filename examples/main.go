package main

import (
	"fmt"

	"github.com/pgaijin66/dea"
)

func main() {
	validEmail := "salima9536@soombo.com"

	_, err := dea.IsDisposableEmail(validEmail)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
