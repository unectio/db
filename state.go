package db

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	StateInitializing	string = "init"
	StateBuilding		string = "building"
	StateReady		string = "ready"
	StateDying		string = "dying"
	StateBroken		string = "broken"
	StateLost		string = "lost"
)

func DyingRq() bson.M {
	return StateRq(StateDying)
}

func BrokenRq() bson.M {
	return StateRq(StateBroken)
}

func ReadyRq() bson.M {
	return StateRq(StateReady)
}

func BuildingRq() bson.M {
	return StateRq(StateBuilding)
}

func StateRq(st string) bson.M {
	return bson.M{"state": st}
}
