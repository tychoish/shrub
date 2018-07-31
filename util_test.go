package shrub

import (
	"testing"
)

func prepend(s string, l []string) []string { return append([]string{s}, l...) }

func assert(t *testing.T, cond bool, msgElems ...string) {
	if !cond {
		t.Error(prepend("assert failed:", msgElems))
	}
}

func require(t *testing.T, cond bool, msgElems ...string) {
	if !cond {
		t.Fatal(prepend("failed require:", msgElems))
	}
}
