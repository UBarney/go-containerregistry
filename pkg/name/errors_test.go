// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package name

import (
	"errors"
	"testing"
)

func TestBadName(t *testing.T) {
	_, err := ParseReference("@@")
	if !IsErrBadName(err) {
		t.Errorf("Not an ErrBadName: %v", err)
	}
	var berr *ErrBadName
	if !errors.As(err, &berr) {
		t.Errorf("Not an ErrBadName using errors.As: %v", err)
	}
	if err.Error() != "could not parse reference: @@. detail: parse tag err: repository can only contain the characters `abcdefghijklmnopqrstuvwxyz0123456789_-./`: @@. parse digest err: a digest must contain exactly one '@' separator (e.g. registry/repository@digest) saw: @@" {
		t.Errorf("Unexpected string: %v", err)
	}
}
