package main

import (
	"os"
	"testing"
)

func TestParseArgsPass(t *testing.T) {
	backupArgs := os.Args
	definedArgs := [4]string{"init", "list", "add", "help"}
	for _, arg := range definedArgs {
		os.Args = append(os.Args, arg)
		got := parseArgs()
		expect := arg
		if got != expect {
			t.Fatalf("got: %v\nwant: %v", got, expect)
		}
		os.Args = backupArgs
	}
	os.Args = backupArgs
}

func TestParseArgsDefault(t *testing.T) {
	backupArgs := os.Args
	got := parseArgs()
	expect := "default"
	if got != expect {
		t.Fatalf("got: %v\nwant: %v", got, expect)
	}
	os.Args = backupArgs
}

func TestParseArgsTooManyArgs(t *testing.T) {
	oldExit := osExit
	defer func() { osExit = oldExit }()

	var status int
	exit := func(code int) {
		status = code
	}
	osExit = exit

	backupArgs := os.Args
	os.Args = append(os.Args, "aaa")
	os.Args = append(os.Args, "bbb")
	parseArgs()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
	os.Args = backupArgs
}
