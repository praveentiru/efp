package efp

import (
	"io"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	exp := "10 > 0"
	var r io.Reader
	r = strings.NewReader(exp)
	val := Parse(&r)
	if strings.Compare(val, "True") != 0 {
		t.Errorf("Expected True but, got False")
	}
}
