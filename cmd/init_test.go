package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/schaermu/gopress/internal/test"
	"github.com/stretchr/testify/assert"
)

var testInitName string = "__TEST_SLIDES__"

func cleanup() {
	if _, err := os.Stat(testInitName); err == nil {
		os.RemoveAll(testInitName)
	}
}

func Test_initCmd(t *testing.T) {
	// test cases for default behavior
	var cases test.TDataProvider = test.TDataProvider{
		// no options
		{
			Command: "init",
			AssertError: func(err error) bool {
				return assert.EqualErrorf(t, err, "requires at least 1 arg(s), only received 0", "init should return an error without args")
			},
			Cleanup: cleanup,
		},
		// arg
		{
			Command: fmt.Sprintf("init %s", testInitName),
			Assert: func(output string) bool {
				return assert.DirExistsf(t, testInitName, "init should create folder %q", testInitName) &&
					assert.FileExists(t, testInitName+"/.gopress.yaml", "init should create config")
			},
			Cleanup: cleanup,
		},
		// flag
		{
			Command: fmt.Sprintf("init %s -t foobar", testInitName),
			Assert: func(output string) bool {
				return assert.DirExistsf(t, testInitName, "init should create folder %q", testInitName) &&
					test.FileContains(t, testInitName+"/.gopress.yaml", "template: foobar", "init should respect custom template")
			},
			Cleanup: cleanup,
		},
	}

	test.RunTestCasesForCmd(t, cases, createInitCommand)
	cleanup()

}
