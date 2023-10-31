package javagen

import (
	"fmt"
	"testing"
)

func TestJavaCommand(t *testing.T) {

	err := JavaCommand(nil, nil)
	if err != nil {
		fmt.Println(err)
	}
}
