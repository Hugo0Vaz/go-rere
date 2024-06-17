package pkg

import (
	"testing"
)

func TestValidRunCommand(t *testing.T) {
	command := "echo hello"

	result, _ := RunCommand(command)

	if result == nil {
		t.Fatal(`Esperava resultado recebeu nil`)
	}
}
