package efp_test

import (
	"strings"
	"testing"

	"github.com/praveentiru/efp"
)

func TestParse(t *testing.T) {
	tt := []struct {
		name string
		exp  string
		out  interface{}
	}{
		{"CONCAT with strings only", `CONCAT("Hello ", "World")`, "Hello World"},
		{"CONCAT with strings and float", `CONCAT("Hello ", 3.1416)`, "Hello 3.1416"},
		{"CONCAT with strings and int", `CONCAT("Hello ", 42)`, "Hello 42"},
		{"EXACT function", `EXACT("Hello", "Hello")`, true},
		{"FIND function with start position", `FIND("Hello", "l", 2)`, 3},
		{"FIND function without start position", `FIND("Hello", "l")`, 3},
		{"FIND function in lower case", `find("Hello", "l")`, 3},
	}
	var errCnt int
	for _, tu := range tt {
		_, err := efp.Parse(strings.NewReader(tu.exp))
		if err != nil {
			t.Logf("Test Case: %v, Parse Error: %v", tu.name, err.Error())
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v cases", errCnt, len(tt))
	}
}
