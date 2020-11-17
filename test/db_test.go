//////////////////////////////////////////////////////////////////////////////
//
// (C) Copyright 2019-2020 by Unectio, Inc.
//
// The information contained herein is confidential, proprietary to Unectio,
// Inc.
//
//////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"

	"github.com/unectio/db"
)

var valid = []string{
	"a_1z",
	"аео", /* %) */
	"_abc",
	"0abc",
	"お_尻",
	"функ_ция",
}

var invalid = []string{
	"a-f",
	"a/f",
	"$abc",
	" abc",
	".abc",
	"abc.",
}

func TestNameValid(t *testing.T) {
	for _, n := range valid {
		if db.ValidName(n) != nil {
			fmt.Printf("Invalid [%s]\n", n)
			t.Fail()
		}
	}

	for _, n := range invalid {
		if db.ValidName(n) == nil {
			fmt.Printf("Valid [%s]\n", n)
			t.Fail()
		}
	}
}
