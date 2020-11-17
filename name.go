/////////////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2019-2020, Unectio Inc, All Right Reserved.
//
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
/////////////////////////////////////////////////////////////////////////////////

package db

import (
	"errors"
	"unicode"

	"github.com/unectio/util"
	"gopkg.in/mgo.v2/bson"
)

type Name struct {
	Name       string `bson:"name"`
	ProjectRef string `bson:"project"`
	/*
	 * Search by name (a.k.a. lookup) uses this field.
	 * Field is indexed in setupIndexes
	 */
	Cookie string `bson:"cookie"`
}

const (
	NameLenMax    int    = 64
	SharedProject string = "*"
)

func ValidName(n string) error {
	if len(n) == 0 {
		return errors.New("empty")
	}
	if len(n) >= NameLenMax {
		return errors.New("too long")
	}

	/*
	 * Name must contain letters, digits, _-s and .-s, but . cannot
	 * be the first or last character. No regex as / is treated as
	 * unicode.L by it.
	 */
	for i, l := range n {
		if !((l == '.' && i != 0 && i != len(n)-1) ||
			unicode.Is(unicode.Letter, l) ||
			unicode.Is(unicode.Digit, l) ||
			l == '_') {
			return errors.New("Invalid character in name")
		}
	}

	return nil
}

func (n *Name) Fill(projectRef, name, extra string) error {
	err := ValidName(name)
	if err != nil {
		return errors.New("bad name value: " + err.Error())
	}

	n.Name = name
	n.ProjectRef = projectRef
	n.Cookie = util.Sha256(n.Str() + "::" + extra)

	return nil
}

func (n *Name) Str() string {
	return n.ProjectRef + "::" + n.Name
}

func (n *Name) Q() bson.M {
	q := Q()
	q["cookie"] = n.Cookie
	return q
}

type IdNameMap map[bson.ObjectId]string

func NewIdNameMap() IdNameMap {
	return make(map[bson.ObjectId]string)
}

func ProjQ(q bson.M, project string) bson.M {
	q["project"] = project
	return q
}
