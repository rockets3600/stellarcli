// Copyright © 2018 Infostellar, Inc.
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

package printer

// Printer is the interface to output values.
type Printer interface {
	// Write buffered data to the underlying io.Writer.
	Flush()
	// Write an object represented as an array.
	Write(r []interface{})
	// Write a header.
	WriteHeader(t []TemplateItem)
	// Write an object with a template.
	WriteWithTemplate(r []map[string]interface{}, t []TemplateItem)
}
