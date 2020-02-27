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
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type KeyDb struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Kind		string			`bson:"kind"`
	Value		[]byte			`bson:"value"`

	User		bson.ObjectId		`bson:"user,omitempty"`
	Scope		*KeyScopeDb		`bson:"scope,omitempty"`
}

type KeyScopeDb struct {
	Project		bson.ObjectId		`bson:"project,omitempty"`
	RemoteAddr	string			`bson:"remote_addr,omitempty"`
}

func (k *KeyDb)ByKind(kind string) bson.M {
	/* .Kind */
	return bson.M{"kind": kind}
}

func (k *KeyDb)ID() bson.ObjectId { return k.Id }
func (k *KeyDb)Location() * mongo.Location { return LocKeys }

const (
	/*
	 * Self-signed key for apilet
	 */
	KeyKindSelf string		= "jwt.self"
	/*
	 * Key signed by managlet
	 */
	KeyKindServer string		= "jwt.server"
	/*
	 * Key used to encrypt secrets in DB
	 */
	KeyKindSecret string		= "apilet.secrets"
	/*
	 * Key to hmac-sign tokens by apilet
	 */
	KeyKindAccess string		= "apilet.access"
)
