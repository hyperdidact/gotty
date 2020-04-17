package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd("hi")
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi", string(out))
	}
}
