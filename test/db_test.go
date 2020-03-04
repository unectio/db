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
	"testing"
	"github.com/unectio/db"
)

var valid = []string {
	"a_1z",
	"аео", /* %) */
	"お_尻",
	"_abc",
	"1abc",
}

var invalid = []string {
	"$abc",
	" abc",
	".abc",
	"abc.",
}

func TestNameValid(t *testing.T) {
	for _, n := range valid {
		if db.ValidName(n) != nil {
			t.Fatalf("Invalid [%s]", n)
		}
	}

	for _, n := range invalid {
		if db.ValidName(n) == nil {
			t.Fatalf("Valid [%s]", n)
		}
	}
}
