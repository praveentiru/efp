package efp

import (
	"strings"
	"testing"

	"github.com/PaesslerAG/gval"
)

func TestConcat(t *testing.T) {
	testCase1 := []interface{}{"Hello ", "World"}
	testCase2 := []interface{}{"Hello World ", 1984}
	testCase3 := []interface{}{"Hello Pi - ", 3.1416}
	tt := []struct {
		testName string
		in       []interface{}
		out      string
	}{
		{"Strings only", testCase1, "Hello World"},
		{"Strings and integers", testCase2, "Hello World 1984"},
		{"Strings and flot", testCase3, "Hello Pi - 3.1416"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Concat(tu.in...)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test: %s, Expected: %s, Got: %s", tu.testName, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v cases", errCnt)
	}
}

func TestExact(t *testing.T) {
	tt := []struct {
		testName string
		inA      string
		inB      string
		out      bool
	}{
		{"Match", "Hello", "Hello", true},
		{"Mis-Match", "Hello", "World", false},
	}
	var errCnt int
	for _, tu := range tt {
		b := Exact(tu.inA, tu.inB)
		if b != tu.out {
			t.Logf("TestCase: %v, Expected: %v, Got: %v", tu.testName, tu.out, b)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v cases", errCnt)
	}
}

func TestFind(t *testing.T) {
	tt := []struct {
		testName string
		iFnd     string
		iSrch    string
		iPos     int
		out      int
	}{
		{"String Found", "B", "ABCDEFGH", 1, 2},
		{"String Not Found", "Z", "ABCDEFGH", 1, -1},
		{"Sub-String Found", "CDE", "ABCDEFGH", 1, 3},
		{"Sub-String Not Found", "GHI", "ABCDEFGH", 1, -1},
		{"String found in starting pos", "CDE", "ABCDEFGH", 3, 3},
		{"String not found in starting pos", "CDE", "ABCDEFGH", 4, -1},
	}
	var errCnt int
	for _, tu := range tt {
		oPos := Find(tu.iFnd, tu.iSrch, tu.iPos)
		if oPos != tu.out {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.testName, tu.out, oPos)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v test cases", errCnt)
	}
}

func TestLeft(t *testing.T) {
	tt := []struct {
		testName string
		inStr    string
		inLen    int
		out      string
	}{
		{"Characters extracted 0", "ABCDEFGH", 0, ""},
		{"Characters extracted less than length", "ABCDEFGH", 3, "ABC"},
		{"Characters extracted more than length", "ABCDEFGH", 10, "ABCDEFGH"},
	}
	var errCnt int
	for _, tu := range tt {
		c := Left(tu.inStr, tu.inLen)
		if strings.Compare(c, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.testName, tu.out, c)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v cases", errCnt)
	}
}

func TestLanguage(t *testing.T) {
	f := gval.Function("EXACT", Exact)
	vars := map[string]interface{}{
		"name":  "Hello",
		"value": "World",
		"type":  "Leaf"}
	// val, err := f.Evaluate("CONCAT(name, value, type)", vars)
	val, err := gval.Evaluate("EXACT(name, value)", vars, f)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if val.(bool) {
		t.Errorf("Expected False Got True")
	}
	// if strings.Compare(val.(string), "SuspensionRearLeaf") != 0 {
	// 	t.Errorf("Expected SuspensionRearLeaf got %s", val)
	// }
}
