// taken & generalized from https://github.com/KEINOS/Hello-Cobra/blob/main/cmd/cmd_helloExtended_test.go
package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type TDataProvider []struct {
	Command     string // command args of the app
	Assert      func(output string) bool
	AssertError func(err error) bool
	MsgError    string // message to display when the test fails
	Cleanup     func()
}

func convertShellArgsToSlice(t *testing.T, str string, cmdName string) []string {
	t.Helper() // With this call go test prints correct lines of code for failed tests.

	cmdArgs, err := shellwords.Parse(str)
	if err != nil {
		t.Fatalf("args parse error: %+v\n", err)
	}

	if len(cmdArgs) == 0 {
		t.Fatalf("args parse error. Command contains fatal strings: %+v\n", str)
	}

	if cmdArgs[0] != cmdName {
		t.Fatalf("args error: Command should start with %s", cmdName)
	}

	return cmdArgs[1:] // trim the first arg
}

func RunTestCasesForCmd(t *testing.T, cases TDataProvider, createCmdFunc func() *cobra.Command) {
	t.Helper() // With this call go test prints correct lines of code for failed tests.

	var (
		subjectCmd *cobra.Command
		argsTmp    []string
	)

	// Loop test cases
	for _, c := range cases {
		if c.Assert == nil && c.AssertError == nil {
			assert.FailNow(t, "Either the Assert or AssertError function has to be defined for a testcase.")
		}

		subjectCmd = createCmdFunc()
		argsTmp = convertShellArgsToSlice(t, c.Command, subjectCmd.Name())

		// init
		subjectCmd.SetArgs(argsTmp)

		capturedOutput := capturer.CaptureOutput(func() {
			// run command
			if err := subjectCmd.Execute(); err != nil {
				if c.AssertError != nil {
					c.AssertError(err)
				} else {
					assert.FailNowf(t, fmt.Sprintf("Failed to execute '%s.Execute()'.", subjectCmd.Name()), "Error msg: %v", err)
				}
			}
		})

		// assert
		if c.Assert != nil {
			c.Assert(capturedOutput)
		}

		// clean up
		if c.Cleanup != nil {
			c.Cleanup()
		}
	}
}

func FileContains(t *testing.T, path string, needle string, msgAndArgs ...interface{}) bool {
	t.Helper()

	fileExists := assert.FileExists(t, path, msgAndArgs...)
	if !fileExists {
		return assert.Fail(t, "file could not be checked for content")
	}

	file, err := os.Open(path)
	if err != nil {
		return assert.Fail(t, fmt.Sprintf("could not open file %q: %v", path, err), msgAndArgs...)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return assert.Fail(t, fmt.Sprintf("could not read file %q: %v", path, err), msgAndArgs...)
	}
	fileContents := buf.String()

	return assert.Containsf(t, fileContents, needle, "could not find %q in %q", needle, fileContents)
}
