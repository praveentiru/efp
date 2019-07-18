package efp_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/praveentiru/efp"
)

// TODO: Add test cases where parsing fails
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
			str, err := eval.EvalString(context.Background(), nil)
			if err != nil {
				t.Logf("Test Case: %v, EvalString returned error, Error: %v", tu.name, err.Error())
				errCnt++
			}
			if strings.Compare(str, v.String()) != 0 {
				t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, str)
				errCnt++
			}
		case reflect.Float64:
			fl, err := eval.EvalFloat64(context.Background(), nil)
			if err != nil {
				t.Logf("Test Case: %v, EvalFloat returned error, Error: %v", tu.name, err.Error())
				errCnt++
			}
			if int(fl) != int(v.Float()) {
				t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, fl)
				errCnt++
			}
		case reflect.Bool:
			b, err := eval.EvalBool(context.Background(), nil)
			if err != nil {
				t.Logf("Test Case: %v, EvalBool returned error, Error: %v", tu.name, err.Error())
				errCnt++
			}
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

func TestNestedStringFunc(t *testing.T) {
	tt := []struct {
		name string
		exp  string
		out  string
	}{
		{"CONCAT with LEFT, MID and RIGHT", `SUBSTITUTE(CONCATENATE(LEFT("Hello World", 5),MID("Hello World", 6, 1),RIGHT("Hello World", 5)), "World", "India")`, "Hello India"},
		{"IF with true return", `IF("Hello" = "Hello", "Hello World", "Hello India")`, "Hello World"},
		{"IF with false return", `IF("Hello" = "hello", "Hello World", "Hello India")`, "Hello India"},
		// {"IF with numerical compariso", `IF(10000 > 1000, "Hello World", "Hello India")`, "Hello World"},
		// {"IF with numerical compariso", `IF(10000 > 1000, 45, 35)`, "45"},
		{"Nested IFs", `IF("Hello" = "hello", "Bah!!!", IF(TRUE, "Hello India", "Bah!!!"))`, "Hello India"},
		{"NOT and False", `IF("Hello" = "hello", "Bah!!!", IF(NOT(FALSE), "Hello India", "Bah!!!"))`, "Hello India"},
	}
	var errCnt int
	for _, tu := range tt {
		eval, err := efp.Parse(strings.NewReader(tu.exp))
		if err != nil {
			t.Logf("Test Case: %v, Expression parse failed, Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		s, err := eval.EvalString(context.Background(), nil)
		if err != nil {
			t.Logf("Test Case: %v, Expression evaluation failed, Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v cases", errCnt, len(tt))
	}
}

func TestBoolFunctions(t *testing.T) {
	tt := []struct {
		name string
		exp  string
		out  bool
	}{
		{"Equality operator for strings", `("Hello" = "Hello")`, true},
		{"Equality check for numbers", `(1 = 1)`, true},
		{"Case mis-match in equality check", `("Hello" = "hello")`, false},
		{"AND with true return", `AND("Hello" = "Hello", 1=1)`, true},
		{"AND with false return", `AND(1=1, "Hello" = "hello")`, false},
		{"OR with true return", `OR("Hello" = "hello", 1=1)`, true},
		{"OR with false return", `OR(1=5, "Hello" = "hello")`, false},
	}
	var errCnt int
	for _, tu := range tt {
		eval, err := efp.Parse(strings.NewReader(tu.exp))
		if err != nil {
			t.Logf("Test Case: %v, Expression parse failed, Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		ret, err := eval.EvalBool(context.Background(), nil)
		if err != nil {
			t.Logf("Test Case: %v, Expression evaluation failed, Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		if tu.out != ret {
			t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, ret)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v cases", errCnt, len(tt))
	}
}

func TestParameterExpressions(t *testing.T) {
	tParam := make(map[string]interface{}, 0)
	tParam["name"] = "India"
	tParam["weight"] = 10
	tParam["location"] = "FA"
	tParam["source"] = "Maharastra"
	tt := []struct {
		name string
		exp  string
		in   map[string]interface{}
		out  string
	}{
		{"IF with equality check with string", `IF(name = "India", "Hello India", "Hello World")`, tParam, "Hello India"},
	}
	var errCnt int
	for _, tu := range tt {
		eval, err := efp.Parse(strings.NewReader(tu.exp))
		if err != nil {
			t.Logf("Test Case: %v, Expression parse failed, Error: %v", tu.name, err.Error())
			errCnt++
			continue
		}
		s, err := eval.EvalString(context.Background(), tu.in)
		if err != nil {
			t.Logf("Test Case: %v, Expression evaluation failed, Error: %v, Input Parameters: %v", tu.name, err.Error(), tParam)
			errCnt++
			continue
		}
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Case: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v cases", errCnt, len(tt))
	}
}
