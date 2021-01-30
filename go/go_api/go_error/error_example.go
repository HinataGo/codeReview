package go_error

import (
	"fmt"
	"testing"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func NewErr(text string) error {
	return &errorString{text}
}

var errType = NewErr("err")

func TestErr(t *testing.T) {
	if errType == NewErr("err") {
		fmt.Println("same")
	}
}
