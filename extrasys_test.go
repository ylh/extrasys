// Adapted from bazil.org/fuse/fs testing packages

package extrasys_test

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type testProc struct {
	fn      func()       // should always exit instead of returning
	cleanup func() error // for instance, delete coredumps from testing pledge
	success bool         // whether zero-exit means success or failure
}

var (
	testProcs = map[string]testProc{}
	procName  = ""
)

const (
	optName = "extrasys.internal.procname"
)

func init() {
	flag.StringVar(&procName, optName, "", "internal use only")
}

// TestCmd generates a proper command that, when executed, runs the test
// corresponding to the given key.
func TestCmd(procName string) (*exec.Cmd, error) {
	exe, err := filepath.Abs(os.Args[0])
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(exe, "-"+optName+"="+procName)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd, nil
}

// ExitsCorrectly is a comprehensive, one-line-of-use wrapper for testing
// a testProc with a key. 
func ExitsCorrectly(procName string, t *testing.T) {
	s := testProcs[procName]
	c, err := testCmd(procName)
	defer s.cleanup()
	if err != nil {
		t.Fatalf("Failed to construct command for %s", procName)
	}
	if (c.Run() == nil) != s.success {
		result := "succeed"
		if !s.success {
			result = "fail"
		}
		t.Fatalf("Process did not %s when it was supposed to", result)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	if procName != "" {
		testProcs[procName].fn()
	}
	os.Exit(m.Run())
}
