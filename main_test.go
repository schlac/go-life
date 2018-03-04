package main

import (
	"os/exec"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Skip()
	if version == "private" {
		t.Error("No custom version set." +
			"Use '-ldflags \"-X main.version=$(VERSION)\"' to set a version.")
	}
}

func TestFmtDone(t *testing.T) {
	cmd := exec.Command("go", "fmt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fail()
		panic(err)
	}
	if len(out) != 0 {
		t.Errorf("Files not formated, run 'go fmt': %s", out)
	}
}
