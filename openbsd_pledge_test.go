// +build openbsd
// +build 386 amd64 arm

package extrasys_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"ylh.io/extrasys"
)

func init() {
	testProcs["pledge"] = testProc{
		func() {
			fmt.Println(extrasys.Pledge("", nil))
			os.Exit(0)
		},
		func() error {
			files, err := ioutil.ReadDir(".")
			if err != nil {
				return err
			}
			for _, file := range files {
				if filepath.Ext(file.Name()) == ".core" {
					if err := os.Remove(file.Name()); err != nil {
						return err
					}
				}
			}
			return nil
		},
		false,
	}
}

func TestPledge(t *testing.T) {
	ExitsCorrectly("pledge", t)
}
