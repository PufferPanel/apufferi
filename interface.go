/*
 Copyright 2019 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package apufferi

import (
	"fmt"
)

func ToStringArray(element interface{}) []string {
	if element == nil {
		return nil
	}
	switch element.(type) {
	case string:
		return []string{element.(string)}
	case []string:
		return element.([]string)
	case []interface{}:
		var arr = make([]string, 0)
		for _, element := range element.([]interface{}) {
			arr = append(arr, ToString(element))
		}
		return arr
	default:
		return []string{}
	}
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
