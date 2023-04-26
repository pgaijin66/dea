package dea

import (
	"reflect"
	"testing"
)

func TestIsDisposableEmail(t *testing.T) {
	tests := []struct {
		name       string
		inputEmail string
		want       bool
	}{
		{name: "Check against valid email", inputEmail: "demo@gmail.com", want: false},
		{name: "Check against invalid email", inputEmail: "demo@027168.com", want: true},
		{name: "Check against random string", inputEmail: "asdsadasdkjh@asdjkhasd.com", want: true},
		{name: "Check against empty strng", inputEmail: "", want: true},
		{name: "Check against invalid email format", inputEmail: "@2or.com", want: true},
		{name: "Check against empty email format", inputEmail: "12345", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := IsDisposableEmail(tc.inputEmail)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected: %#v, got: %#v", tc.name, tc.want, got)
			}
		})
	}

}
