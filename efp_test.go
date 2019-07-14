package efp_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/praveentiru/efp"
)

// TODO: Add failure cases to ensure that parser throws errors when parameters are incorrect
// TODO: Add test cases for nested function calls
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
		{"FIND function with start position", `FIND("l", "Hello", 2)`, 3.0},
		{"FIND function without start position", `FIND("l", "Hello")`, 3.0},
		{"FIND function in lower case", `FIND("l", "Hello")`, 3.0},
		{"LEFT with string and no num chars", `LEFT("Hello World")`, "H"},
		{"LEFT with string and with num chars", `LEFT("Hello World", 5)`, "Hello"},
		{"LEFT with number input in place of text", `LEFT(3.1416, 3)`, "3.1"},
		{"LEN for string", `LEN("Hello World")`, 11.0},
		{"LOWER with string", `LOWER("HeLLo World")`, "hello world"},
		{"MID with string", `MID("Hello World", 3, 3)`, "llo"},
		{"PROPER with string", `PROPER("HEllO wORLd")`, "Hello World"},
		{"REPLACE with string", `REPLACE("Hello World", 7, 5, "India")`, "Hello India"},
		{"REPT with string", `REPT("Hell", 4)`, "HellHellHellHell"},
		{"RIGHT with string and num", `RIGHT("Hell", 2)`, "ll"},
		{"SEARCH with mis-match case", `SEARCH("LL", "Hello World")`, 3.0},
		{"SUBSTITUTE with string", `SUBSTITUTE("Oink Oink Oink", "ink", "inky", 2)`, "Oink Oinky Oink"},
		{"TRIM with spaces", `TRIM("    Hello    World    ")`, "Hello World"},
		{"UPPER with inconsistent csae", `UPPER("Hello India")`, "HELLO INDIA"},
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

func TestEvaluate(t *testing.T) {
	tt := []struct {
		name string
		exp  string
		out  interface{}
	}{
		{"CONCAT with strings only", `CONCAT("Hello ", "World")`, "Hello World"},
		{"CONCAT with strings and float", `CONCAT("Hello ", 3.1416)`, "Hello 3.1416"},
		{"CONCAT with strings and int", `CONCAT("Hello ", 42)`, "Hello 42"},
		{"EXACT function", `EXACT("Hello", "Hello")`, true},
		{"FIND function with start position", `FIND("l", "Hello", 2)`, 3.0},
		{"FIND function without start position", `FIND("l", "Hello")`, 3.0},
		{"FIND function in lower case", `FIND("l", "Hello")`, 3.0},
		{"LEFT with string and no num chars", `LEFT("Hello World")`, "H"},
		{"LEFT with string and with num chars", `LEFT("Hello World", 5)`, "Hello"},
		{"LEFT with number input in place of text", `LEFT(3.1416, 3)`, "3.1"},
		{"LEN for string", `LEN("Hello World")`, 11.0},
		{"LOWER with string", `LOWER("HeLLo World")`, "hello world"},
		{"MID with string", `MID("Hello World", 3, 3)`, "llo"},
		{"PROPER with string", `PROPER("HEllO wORLd")`, "Hello World"},
		{"REPLACE with string", `REPLACE("Hello World", 7, 5, "India")`, "Hello India"},
		{"REPT with string", `REPT("Hell", 4)`, "HellHellHellHell"},
		{"RIGHT with string and num", `RIGHT("Hell", 2)`, "ll"},
		{"SEARCH with mis-match case", `SEARCH("LL", "Hello World")`, 3.0},
		{"SUBSTITUTE with string", `SUBSTITUTE("Oink Oink Oink", "ink", "inky", 2)`, "Oink Oinky Oink"},
		{"TRIM with spaces", `TRIM("    Hello    World    ")`, "Hello World"},
		{"UPPER with inconsistent csae", `UPPER("Hello India")`, "HELLO INDIA"},
	}
	var errCnt int
	for _, tu := range tt {
		eval, err := efp.Parse(strings.NewReader(tu.exp))
		if err != nil {
			t.Logf("Test Case: %v, Parse Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		if eval == nil {
			t.Logf("Test Case: %v, No Parse error but Evaluable is nil", tu.name)
			errCnt++
			continue
		}
		switch v := reflect.ValueOf(tu.out); v.Kind() {
		case reflect.String:
			str, _ := eval.EvalString(context.Background(), nil)
			if strings.Compare(str, v.String()) != 0 {
				t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, str)
				errCnt++
			}
		case reflect.Float64:
			fl, _ := eval.EvalFloat64(context.Background(), nil)
			if int(fl) != int(v.Float()) {
				t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, fl)
				errCnt++
			}
		case reflect.Bool:
			b, _ := eval.EvalBool(context.Background(), nil)
			if b != v.Bool() {
				t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, b)
				errCnt++
			}
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v cases", errCnt, len(tt))
	}
}