package addons

type content struct {
	kind  string `bson:"kind"`
	value string `bson:"value"`
}

type groupFilter struct {
	name string `bson:"group_name"`
}

type group struct {
	name    string    `bson:"group_name"` // composition doesn't work ):
	content []content `bson:"content"`
}
