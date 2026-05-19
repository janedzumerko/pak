package gitx

import "testing"

func TestPassthrough_GitVersion(t *testing.T) {
	code, err := Passthrough([]string{"--version"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
}

func TestPassthrough_InvalidCommand(t *testing.T) {
	code, err := Passthrough([]string{"not-a-real-command"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code == 0 {
		t.Errorf("expected non-zero exit code for invalid git command")
	}
}
