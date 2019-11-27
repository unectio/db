package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type TriggerDb struct {
	DbCommon				`bson:",inline"`

	FnId		bson.ObjectId		`bson:"fnid"`
	CallKey		string			`bson:"call_key"`
	Src		string			`bson:"source"`

	/*
	 * Search key by which event locates the trigger.
	 * XXX This field must be indexed.
	 */
	SearchKey	string			`bson:"key"`
	/*
	 * Sort key by which callet's runner sorts the triggers
	 * to fire event on.
	 */
	SortKey		string			`bson:"sort"`

	URL		*URLTrigDb		`bson:"url,omitempty"`
	Cron		*CronTrigDb		`bson:"cron,omitempty"`
	Websock		*WsTrigDb		`bson:"websock,omitempty"`
	Event		*EventTrigDb		`bson:"event,omitempty"`
}

func (tg *TriggerDb)ByFn(fnid bson.ObjectId) bson.M {
	/* .FnId */
	return bson.M{"fnid": fnid}
}

func (tg *TriggerDb)ID() bson.ObjectId { return tg.Id }
func (tg *TriggerDb)Location() *mongo.Location { return LocTrigger }

type URLTrigDb struct {
	Compute		ComputeDb		`bson:"compute"`
	Cookie		string			`bson:"cookie"`
	Auth		bson.ObjectId		`bson:"auth,omitempty"`
}

func (_ *URLTrigDb)DbKey(id string) string {
	return "url::" + id
}

type CronTrigDb struct {
	Tab		string			`bson:"tab"`
	Args		map[string]string	`bson:"args"`
	Compute		ComputeDb		`bson:"compute"`
}

func (_ *CronTrigDb)DbKey(comp, st string) string {
	return "cron::" + comp + ":" + st
}

/* Update */
func (tg *TriggerDb)Update() *TriggerDbUpd {
	return &TriggerDbUpd{
		tg:	tg,
		q:	bson.M{},
	}
}

type TriggerDbUpd struct {
	tg	*TriggerDb
	q	bson.M
}

func (u *TriggerDbUpd)SortKey(key string) {
	u.tg.SortKey = key
	u.q["sort"] = key
}

func (u *TriggerDbUpd)CallKey(key string) {
	u.tg.CallKey = key
	u.q["call_key"] = key
}

func (u *TriggerDbUpd)Q() bson.M {
	return u.q
}
