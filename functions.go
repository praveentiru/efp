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
