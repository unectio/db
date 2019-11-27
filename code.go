package db

import (
	"github.com/unectio/util/mongo"
	"gopkg.in/mgo.v2/bson"
)

type FuncCodeDb struct {
	DbCommon				`bson:",inline"`

	State		string			`bson:"state"`

	FnId		bson.ObjectId		`bson:"fnid"`
	Lang		string			`bson:"lang"`
	Repo		*FuncRepoLink		`bson:"repo_link,omitempty"`

	Gen		int			`bson:"gen"`

	/*
	 * Key for lookup (deployment -> *code).
	 * XXX The field must be indexed.
	 */
	DepKey		string			`bson:"dep_key"`
}

func (c *FuncCodeDb)ByFn(id bson.ObjectId) bson.M {
	/* .FnId */
	return bson.M{"fnid": id}
}

func (c *FuncCodeDb)ByRepo(id bson.ObjectId) bson.M {
	/* .Repo.Id */
	return bson.M{"repo_link.id": id}
}

func (c *FuncCodeDb)ByDepKey(dk string) bson.M {
	/* .DepKey */
	return bson.M{"dep_key": dk}
}

func (c *FuncCodeDb)ID() bson.ObjectId { return c.Id }
func (c *FuncCodeDb)Location() *mongo.Location { return LocCode }

func (c *FuncCodeDb)Version() string { return c.Id.Hex() }

type FuncRepoLink struct {
	Id		bson.ObjectId		`bson:"id"`
	Path		string			`bson:"path"`
}


func (c *FuncCodeDb)UpdateQ() bson.M {
	/* .Repo & .Gen */
	return bson.M{"repo_link": c.Repo, "gen": c.Gen}
}
