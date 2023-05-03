package passgen_test

import (
	"os"
	"testing"

	"github.com/qba73/passgen"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"passgen": passgen.Main,
	}))
}

func TestPassgenCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
