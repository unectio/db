package db

type ComputeDb struct {
	/*
	 * Name by which apilet (and users) sees the compute
	 */
	Name	string		`bson:"name"`
	/*
	 * ID by which the callet knows the compute it works for
	 */
	Cookie	string		`bson:"cookie"`
}
