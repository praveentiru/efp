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
	"io"

	"github.com/PaesslerAG/gval"
)

// Parse parses the excel formula provided and returns the Eval interface which can be used to evaluate formula
func Parse(r *io.Reader) string {

	vars := map[string]interface{}{"dummy": "dummy"}
	val, _ := gval.Evaluate("10 > 0", vars)
	if val == true {
		return "True"
	}
	return "False"
}
