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

const (
	StateDB    string = "faas_state"
	FuncCol    string = "Functions"
	TriggerCol string = "Triggers"
	TargetsCol string = "Targets"
	CodeCol    string = "Codes"
	RepoCol    string = "Repositories"
	RouterCol  string = "Routers"
	ACtxCol    string = "AuthCtxs"
	AppCol     string = "Applications"
	SecretCol  string = "Secrets"
	WebsockCol string = "Websockets"
	MongoCol   string = "Mongos"

	DataDB        string = "faas_data"
	LogsCol       string = "Logs"
	StatsCol      string = "FuncStats"
	PStatsCol     string = "ProjectStats"
	DeferEventCol string = "DeferEvents"

	AdminDB     string = "faas_admin"
	UsersCol    string = "users"
	KeysCol     string = "keys"
	ProjectsCol string = "projects"
	RolesCol    string = "roles"
	MwareCol    string = "mwares"
	ClassesCol  string = "classes"
	CapsCol     string = "caps"
	EventSrcCol string = "eventsources"
)

var (
	LocFunc = &mongo.Location{
		Db:  StateDB,
		Col: FuncCol,
	}
	LocTrigger = &mongo.Location{
		Db:  StateDB,
		Col: TriggerCol,
	}
	LocTarget = &mongo.Location{
		Db:  StateDB,
		Col: TargetsCol,
	}
	LocCode = &mongo.Location{
		Db:  StateDB,
		Col: CodeCol,
	}
	LocRepo = &mongo.Location{
		Db:  StateDB,
		Col: RepoCol,
	}
	LocRouter = &mongo.Location{
		Db:  StateDB,
		Col: RouterCol,
	}
	LocAuthCtx = &mongo.Location{
		Db:  StateDB,
		Col: ACtxCol,
	}
	LocApps = &mongo.Location{
		Db:  StateDB,
		Col: AppCol,
	}
	LocSecret = &mongo.Location{
		Db:  StateDB,
		Col: SecretCol,
	}
	LocWebsock = &mongo.Location{
		Db:  StateDB,
		Col: WebsockCol,
	}
	LocMongo = &mongo.Location{
		Db:  StateDB,
		Col: MongoCol,
	}

	LocLogs = &mongo.Location{
		Db:  DataDB,
		Col: LogsCol,
	}
	LocStats = &mongo.Location{
		Db:  DataDB,
		Col: StatsCol,
	}
	LocPStats = &mongo.Location{
		Db:  DataDB,
		Col: PStatsCol,
	}
	LocDeferEvs = &mongo.Location{
		Db:  DataDB,
		Col: DeferEventCol,
	}

	LocUsers = &mongo.Location{
		Db:  AdminDB,
		Col: UsersCol,
	}
	LocKeys = &mongo.Location{
		Db:  AdminDB,
		Col: KeysCol,
	}
	LocProjects = &mongo.Location{
		Db:  AdminDB,
		Col: ProjectsCol,
	}
	LocRoles = &mongo.Location{
		Db:  AdminDB,
		Col: RolesCol,
	}
	LocMwares = &mongo.Location{
		Db:  AdminDB,
		Col: MwareCol,
	}
	LocClasses = &mongo.Location{
		Db:  AdminDB,
		Col: ClassesCol,
	}
	LocCaps = &mongo.Location{
		Db:  AdminDB,
		Col: CapsCol,
	}
	LocEventSources = &mongo.Location{
		Db:  AdminDB,
		Col: EventSrcCol,
	}
)

type DbObject interface {
	Location() *mongo.Location
	ID() bson.ObjectId
}

func NewID() bson.ObjectId {
	return bson.NewObjectId()
}

type DbCommon struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name `bson:",inline"`
	/*
	 * Search by tags (db/list.go) uses this field, which
	 * ... probably must be indexed too
	 */
	Tags     []string `bson:"tags,omitempty"`
	UserData string   `bson:"userdata,omitempty"`
}
