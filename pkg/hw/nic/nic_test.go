// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: BSD-3-Clause
//

package nic

import (
	"testing"
)

func TestList(t *testing.T) {
	l := List()
	if len(l) == 0 {
		t.Fatal("no nics")
	}
}
