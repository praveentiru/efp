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
	gval.Function("FIND", func(src, fnd interface{}, num ...float64) int {
		srcStr := toString(src)
		fndStr := toString(fnd)
		startPos := 1
		if len(num) > 0 {
			startPos = int(num[0])
		}
		return Find(srcStr, fndStr, startPos)
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
