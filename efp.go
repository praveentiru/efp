// Copyright © 2019 Praveen Tirumandyam praveen.tirumandyam@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package efp

import (
	"fmt"
	"io"
	"strings"

	"github.com/PaesslerAG/gval"
)

var excelLanguage = gval.NewLanguage(
	excelText,
	gval.Base(),
)

var excelText = gval.NewLanguage(
	gval.Function("CONCAT", func(args ...interface{}) string {
		str := make([]string, 0)
		for _, arg := range args {
			str = append(str, toString(arg))
		}
		return Concat(str...)
	}),
	gval.Function("CONCATENATE", func(args ...interface{}) string {
		str := make([]string, 0)
		for _, arg := range args {
			str = append(str, toString(arg))
		}
		return Concat(str...)
	}),
	gval.Function("EXACT", func(a, b interface{}) bool {
		aStr := toString(a)
		bStr := toString(b)
		return Exact(aStr, bStr)
	}),
	gval.Function("FIND", func(src, fnd interface{}, num ...float64) float64 {
		srcStr := toString(src)
		fndStr := toString(fnd)
		startPos := 1
		if len(num) > 0 {
			startPos = int(num[0])
		}
		return float64(Find(srcStr, fndStr, startPos))
	}),
	gval.Function("LEFT", func(s interface{}, num ...float64) string {
		str := toString(s)
		l := 1
		if len(num) > 0 {
			l = int(num[0])
		}
		return Left(str, l)
	}),
	gval.Function("LEN", func(s interface{}) float64 {
		str := toString(s)
		return float64(Len(str))
	}),
	gval.Function("LOWER", func(s interface{}) string {
		str := toString(s)
		return Lower(str)
	}),
	gval.Function("MID", func(s interface{}, strt, num float64) string {
		str := toString(s)
		return Mid(str, int(strt), int(num))
	}),
	gval.Function("PROPER", func(a interface{}) string {
		str := toString(a)
		return Proper(str)
	}),
	gval.Function("REPLACE", func(a interface{}, strt, num float64, b interface{}) string {
		aStr := toString(a)
		bStr := toString(b)
		return Replace(aStr, int(strt), int(num), bStr)
	}),
	gval.Function("REPT", func(s interface{}, num float64) string {
		str := toString(s)
		return Rept(str, int(num))
	}),
	gval.Function("RIGHT", func(s interface{}, num ...float64) string {
		str := toString(s)
		l := 1
		if len(num) > 0 {
			l = int(num[0])
		}
		return Right(str, l)
	}),
	gval.Function("SEARCH", func(fnd, src interface{}, strt ...float64) float64 {
		fndStr := toString(fnd)
		srcStr := toString(src)
		l := 1
		if len(strt) > 0 {
			l = int(strt[0])
		}
		return float64(Search(fndStr, srcStr, l))
	}),
	gval.Function("SUBSTITUTE", func(src, old, new interface{}, num ...float64) string {
		srcStr := toString(src)
		oldStr := toString(old)
		newStr := toString(new)
		n := 0
		if len(num) > 0 {
			n = int(num[0])
		}
		return Substitute(srcStr, oldStr, newStr, n)
	}),
	gval.Function("TRIM", func(s interface{}) string {
		str := toString(s)
		return Trim(str)
	}),
	gval.Function("UPPER", func(s interface{}) string {
		str := toString(s)
		return Upper(str)
	}),
)

// toString leverages fmt library to convert values to string
func toString(v interface{}) string {
	return fmt.Sprint(v)
}

// Parse parses the excel formula provided and returns the Eval interface which can be used to evaluate formula
func Parse(r io.WriterTo) (gval.Evaluable, error) {
	var sb strings.Builder
	var wrTo io.Writer = &sb
	r.WriteTo(wrTo)
	exp := sb.String()

	return excelLanguage.NewEvaluable(exp)
}
