package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	StateDB string		= "faas_state"
	FuncCol string		= "Functions"
	TriggerCol string	= "Triggers"
	CodeCol string		= "Codes"
	RepoCol string		= "Repositories"
	RouterCol string	= "Routers"
	ACtxCol string		= "AuthCtxs"
	AppCol string		= "Applications"
	SecretCol string	= "Secrets"
	WebsockCol string	= "Websockets"
	MongoCol string		= "Mongos"

	DataDB string		= "faas_data"
	LogsCol string		= "Logs"
	StatsCol string		= "FuncStats"
	PStatsCol string	= "ProjectStats"
	DeferEventCol string	= "DeferEvents"

	AdminDB string		= "faas_admin"
	UsersCol string		= "users"
	KeysCol string		= "keys"
	ProjectsCol string	= "projects"
	RolesCol string		= "roles"
	MwareCol string		= "mwares"
	ClassesCol string	= "classes"
	CapsCol string		= "caps"
	EventSrcCol string	= "eventsources"
)

var (
	LocFunc		= &mongo.Location{StateDB, FuncCol}
	LocTrigger	= &mongo.Location{StateDB, TriggerCol}
	LocCode		= &mongo.Location{StateDB, CodeCol}
	LocRepo		= &mongo.Location{StateDB, RepoCol}
	LocRouter	= &mongo.Location{StateDB, RouterCol}
	LocAuthCtx	= &mongo.Location{StateDB, ACtxCol}
	LocApps		= &mongo.Location{StateDB, AppCol}
	LocSecret	= &mongo.Location{StateDB, SecretCol}
	LocWebsock	= &mongo.Location{StateDB, WebsockCol}
	LocMongo	= &mongo.Location{StateDB, MongoCol}

	LocLogs		= &mongo.Location{DataDB, LogsCol}
	LocStats	= &mongo.Location{DataDB, StatsCol}
	LocPStats	= &mongo.Location{DataDB, PStatsCol}
	LocDeferEvs	= &mongo.Location{DataDB, DeferEventCol}

	LocUsers	= &mongo.Location{AdminDB, UsersCol}
	LocKeys		= &mongo.Location{AdminDB, KeysCol}
	LocProjects	= &mongo.Location{AdminDB, ProjectsCol}
	LocRoles	= &mongo.Location{AdminDB, RolesCol}
	LocMwares	= &mongo.Location{AdminDB, MwareCol}
	LocClasses	= &mongo.Location{AdminDB, ClassesCol}
	LocCaps		= &mongo.Location{AdminDB, CapsCol}
	LocEventSources	= &mongo.Location{AdminDB, EventSrcCol}
)

type DbObject interface {
	Location() *mongo.Location
	ID() bson.ObjectId
}

func NewID() bson.ObjectId {
	return bson.NewObjectId()
}

type DbCommon struct {
	Id		bson.ObjectId		`bson:"_id,omitempty"`
	Name					`bson:",inline"`
	/*
	 * Search by tags (db/list.go) uses this field, which
	 * ... probably must be indexed too
	 */
	Tags		[]string		`bson:"tags,omitempty"`
	UserData	string			`bson:"userdata,omitempty"`
}
