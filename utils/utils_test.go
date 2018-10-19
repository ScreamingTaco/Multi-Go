package utils

/*
   Copyright 2018 ScreamingTaco

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

import (
	"testing"
)

func TestbytesToGigabytes(t *testing.T) {
	got := bytesToGigabytes(4034846720)
	want := 4.03

	if got != want {
		t.Errorf("got '%f' want '%f'", got, want)
	}
}
