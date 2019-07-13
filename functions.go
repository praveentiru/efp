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
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

// Concat implement Excel's CONCAT function
func Concat(args ...interface{}) string {
	var b strings.Builder
	for _, v := range args {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			b.WriteString(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			b.WriteString(strconv.FormatInt(v.Int(), 10))
		case reflect.Float32, reflect.Float64:
			b.WriteString(fmt.Sprint(v.Float()))
		}
	}
	return b.String()
}

// Exact implments Excel's EXACT function
func Exact(a string, b string) bool {
	if strings.Compare(a, b) != 0 {
		return false
	}
	return true
}

// Find implements Excel's FIND function
// Deviation: pos is mandatory input
func Find(f, w string, pos int) int {
	if pos > len(w) {
		return -1
	}
	r := strings.Index(w[pos-1:], f)
	if r == -1 {
		return r
	}
	return pos + r
}

// Left implements Excel's LEFT function
// Deviation: No default value for l unlike Excel where it is 1
func Left(s string, l int) string {
	if l > len(s) {
		l = len(s)
	}
	return s[:l]
}

// Len implements Excel's LEN function
func Len(s string) int {
	return len(s)
}

// Lower implements Excel's LOWER function
func Lower(s string) string {
	return strings.ToLower(s)
}

// Mid implements Excel's MID function
func Mid(s string, strt, num int) string {
	l := len(s)
	if strt > l || strt < 1 || num < 0 {
		return ""
	}
	if strt+num > l {
		return s[strt-1:]
	}
	return s[strt-1 : strt-1+num]
}

// Proper implements Excel's PROPER function
func Proper(s string) string {
	type scanState int
	const (
		inString scanState = iota
		outString
	)
	var sb strings.Builder
	var sc scanner.Scanner
	currState := outString
	sc.Init(strings.NewReader(s))
	for rn := sc.Next(); rn != scanner.EOF; rn = sc.Next() {
		if unicode.IsLetter(rn) {
			switch currState {
			case inString:
				sb.WriteRune(unicode.ToLower(rn))
			case outString:
				currState = inString
				sb.WriteRune(unicode.ToUpper(rn))
			}
		} else {
			if currState == inString {
				currState = outString
			}
			sb.WriteRune(rn)
		}
	}
	return sb.String()
}

// Replace implements Excel's REPLACE function
func Replace(old string, strt, num int, newStr string) string {
	strt = strt - 1
	l := len(old)
	if strt+num > l {
		num = l - strt
		if strt > l {
			strt = l
			num = 0
		}
	}
	left := old[:strt]
	right := old[strt+num:]
	var sb strings.Builder
	sb.WriteString(left)
	sb.WriteString(newStr)
	sb.WriteString(right)
	fmt.Println(sb.String())
	return sb.String()
}

// Rept implements Excel's REPT function
func Rept(s string, r int) string {
	var sb strings.Builder
	for i := 0; i < r; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// Right implements Excel's RIGHT function
func Right(s string, num int) string {
	if num <= 0 {
		return ""
	}
	l := len(s)
	if num > l {
		return s
	}
	return s[l-num:]
}

// Search implements Excel's SEARCH function
// Deviation: pos parameter is mandatory in GO
func Search(fnd, in string, pos int) int {
	fnd = strings.ToLower(fnd)
	in = strings.ToLower(in)
	return Find(fnd, in, pos)
}

// Substitute implements Excel's SUBSTITUTE function
// Deviation: Instance number is mandatory and if all instances need to be substituted provide 0
func Substitute(src, old, new string, inst int) string {
	if inst > 0 {
		substrA := strings.Split(src, old)
		var sb strings.Builder
		for i, val := range substrA {
			if i > 0 {
				switch i {
				case inst:
					sb.WriteString(new)
				default:
					sb.WriteString(old)
				}
			}
			sb.WriteString(val)
		}
		return sb.String()
	}
	return strings.ReplaceAll(src, old, new)
}

// Trim implements Excel's TRIM function
func Trim(s string) string {
	s = strings.TrimSpace(s)
	type scanState int
	const (
		inSpace scanState = iota
		outSpace
	)
	var sc scanner.Scanner
	var sb strings.Builder
	sc.Init(strings.NewReader(s))
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		str := fmt.Sprint(sc.TokenText(), " ")
		sb.WriteString(str)
	}
	return strings.TrimSpace(sb.String())
}

// Upper implements Excel's UPPER function
func Upper(s string) string {
	return strings.ToUpper(s)
}
