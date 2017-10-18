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

package cmrdr

import (
	"encoding/json"
	"strings"
)

// ParseCredentialsFromString opens a dictionary file and returns its contents as a Credentials structure
func ParseCredentialsFromString(content string) (Credentials, error) {
	var creds Credentials

	// Unmarshal content of JSON file into data structure
	err := json.Unmarshal([]byte(content), &creds)
	if err != nil {
		return creds, err
	}

	return creds, nil
}

// ParseRoutesFromString opens a dictionary file and returns its contents as a Routes structure
func ParseRoutesFromString(content string) Routes {
	return strings.Split(content, "\n")
}
