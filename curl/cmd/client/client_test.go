package client

import (
	"os/exec"
	"testing"
)

func TestClient(t *testing.T) {
	inputs := []string{"http://httpbin.org/get", "http://httpbin.org/"}

	for _, input := range inputs {
		cmd := exec.Command("../../bin/gocurl", "-v", input)
		err := cmd.Run()
		if err != nil {
			t.Errorf("CLI failed to hit %s: %v", input, err)
			return
		}
	}
}
