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
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
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
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
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
		{"String not found in starting pos", "CDE", "ABCDEFGH", 11, -1},
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
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
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
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestLen(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  int
	}{
		{"5 characters", "ABCDE", 5},
		{"Null string", "", 0},
		{"String with spaces", "     ABCDE", 10},
	}
	var errCnt int
	for _, tu := range tt {
		l := Len(tu.in)
		if l != tu.out {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, l)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestLower(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Null string", "", ""},
		{"All Uppercase", "ABCDEF", "abcdef"},
		{"Mixed Case", "ABcdEF", "abcdef"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Lower(tu.in)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestMid(t *testing.T) {
	tt := []struct {
		name  string
		iStr  string
		iStrt int
		iNum  int
		out   string
	}{
		{"Three characters starting from 3 position", "ABCDEFGH", 3, 3, "CDE"},
		{"Start greater than length of string", "ABCDEFGH", 11, 3, ""},
		{"Sum of start and num greater than length", "ABCDEFGH", 4, 7, "DEFGH"},
		{"Start less than 1", "ABCDEFGH", 0, 7, ""},
		{"Num negative", "ABCDEFGH", 4, -5, ""},
	}
	var errCnt int
	for _, tu := range tt {
		s := Mid(tu.iStr, tu.iStrt, tu.iNum)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestProper(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Simple Sentence", "hello my world, how are you?", "Hello My World, How Are You?"},
		{"Sentence with caps in middle", "heLLo my woRLd, hOw are yoU?", "Hello My World, How Are You?"},
		{"Sentence with words starting with num", "10BuGar 71-haiKu", "10Bugar 71-Haiku"},
		{"Sentence with num in between", "abcd17edgh", "Abcd17Edgh"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Proper(tu.in)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestReplace(t *testing.T) {
	tt := []struct {
		name    string
		inStr   string
		inStart int
		inNum   int
		rString string
		out     string
	}{
		{"Happy case 1", "Hello World", 3, 6, "India", "HeIndiarld"},
		{"Happy case 2", "Hello World", 7, 5, "India", "Hello India"},
		{"Start position larger than length", "Hello World", 12, 5, " and India", "Hello World and India"},
		{"Start position larger than length", "Hello World", 17, 5, " and India", "Hello World and India"},
		{"Sum of start and num larger than length", "Hello World", 7, 10, "India", "Hello India"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Replace(tu.inStr, tu.inStart, tu.inNum, tu.rString)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestRept(t *testing.T) {
	tt := []struct {
		name  string
		inStr string
		inNum int
		out   string
	}{
		{"Case 1", "*-", 5, "*-*-*-*-*-"},
		{"Case 2", "*-", 0, ""},
		{"Case 3", "*-", -1, ""},
		{"Case 4", "__", 3, "______"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Rept(tu.inStr, tu.inNum)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestRight(t *testing.T) {
	tt := []struct {
		name  string
		inStr string
		inNum int
		out   string
	}{
		{"Happy Case 1", "Hello World", 5, "World"},
		{"Happy Case 2", "Hello World", 0, ""},
		{"Num is equal to string length", "Hello World", 11, "Hello World"},
		{"Num greater than string length", "Hello World", 16, "Hello World"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Right(tu.inStr, tu.inNum)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestSearch(t *testing.T) {
	tt := []struct {
		testName string
		iFnd     string
		iSrch    string
		iPos     int
		out      int
	}{
		{"String Found: Same Case", "B", "ABCDEFGH", 1, 2},
		{"String Not Found: Same Case", "Z", "ABCDEFGH", 1, -1},
		{"Sub-String Found: Same Case", "CDE", "ABCDEFGH", 1, 3},
		{"Sub-String Not Found: Same Case", "GHI", "ABCDEFGH", 1, -1},
		{"String found in starting pos: Same Case", "CDE", "ABCDEFGH", 3, 3},
		{"String not found in starting pos: Same Case", "CDE", "ABCDEFGH", 4, -1},
		{"String Found: Different Case", "b", "ABCDEFGH", 1, 2},
		{"String Not Found: Different Case", "z", "ABCDEFGH", 1, -1},
		{"Sub-String Found: Different Case", "cde", "ABCDEFGH", 1, 3},
		{"Sub-String Not Found: Different Case", "ghi", "ABCDEFGH", 1, -1},
		{"String found in starting pos: Different Case", "cDe", "ABCDEFGH", 3, 3},
		{"String not found in starting pos: Different Case", "Cde", "ABCDEFGH", 8, -1},
	}
	var errCnt int
	for _, tu := range tt {
		oPos := Search(tu.iFnd, tu.iSrch, tu.iPos)
		if oPos != tu.out {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.testName, tu.out, oPos)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestSubstitute(t *testing.T) {
	tt := []struct {
		name   string
		inSrc  string
		inOld  string
		inNew  string
		inInst int
		out    string
	}{
		{"Happy Case", "a,b,c,d,e", ",", "|", 0, "a|b|c|d|e"},
		{"Specific Instance in limit", "a,b,c,d,e", ",", "|", 3, "a,b,c|d,e"},
		{"Specific Instance when found at start", ",a,b,c,d,e", ",", "|", 3, ",a,b|c,d,e"},
		{"Specific Instance when out of limit", ",a,b,c,d,e", ",", "|", 7, ",a,b,c,d,e"},
		{"Substring: Happy Case", "Hello abcWorldabc and abcIndia", "abc", "", 0, "Hello World and India"},
		{"Substring: Specific Instance in limit", "Hello abcWorldabc and abcIndia", "abc", "", 2, "Hello abcWorld and abcIndia"},
		{"Substring: Specific Instance when found at start", "abcHello abcWorldabc and abcIndia", "abc", "", 3, "abcHello abcWorld and abcIndia"},
		{"Substring: Specific Instance when out of limit", "Hello abcWorldabc and abcIndia", "abc", "", 7, "Hello abcWorldabc and abcIndia"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Substitute(tu.inSrc, tu.inOld, tu.inNew, tu.inInst)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestTrim(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Happy case", "  Hello World  ", "Hello World"},
		{"Excess space in sentence", "  Hello	   World  ", "Hello World"},
		{"Tabs in beginning", "	  Hello World	  ", "Hello World"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Trim(tu.in)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestUpper(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Null string", "", ""},
		{"All Lowercase", "abcdef", "ABCDEF"},
		{"Mixed Case", "ABcdEF", "ABCDEF"},
	}
	var errCnt int
	for _, tu := range tt {
		s := Upper(tu.in)
		if strings.Compare(s, tu.out) != 0 {
			t.Logf("Test Name: %v, Expected: %v, Got: %v", tu.name, tu.out, s)
			errCnt++
		}
	}
	if errCnt > 0 {
		t.Errorf("Failed %v of %v test cases", errCnt, len(tt))
	}
}

func TestLanguage(t *testing.T) {
	f1 := gval.Function("CONCAT", Concat)
	f2 := gval.Function("LEFT", func(s string, num float64) string {
		return Left(s, int(num))
	})
	f3 := gval.Function("MID", func(s string, str, num float64) string {
		return Mid(s, int(str), int(num))
	})
	langs := []gval.Language{f1, f2, f3}
	// vars := map[string]interface{}{
	// 	"name":  "Hello",
	// 	"value": "World",
	// 	"type":  "Leaf"}
	val, err := gval.Evaluate(`CONCAT(LEFT("ABCD",2),MID("ABCD",3,2))`, nil, langs...)
	// val, err := gval.Evaluate(`LEFT("ABCD",2)`, nil, langs...)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if strings.Compare(val.(string), "ABCD") != 0 {
		t.Errorf("Expected ABCD Got %v", val.(string))
	}
	// if strings.Compare(val.(string), "SuspensionRearLeaf") != 0 {
	// 	t.Errorf("Expected SuspensionRearLeaf got %s", val)
	// }
}
