// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package information

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColumn(t *testing.T) {
	Convey("GetColumnsQuery", t, func() {
		st := GetColumnsQuery("foo")
		So(st, ShouldNotBeNil)
		So(st.Params["table_name"], ShouldEqual, "foo")
		// fmt.Println(st)
	})
}
