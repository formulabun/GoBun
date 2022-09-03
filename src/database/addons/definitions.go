package addons

type content struct {
	Kind  string `bson:"kind"`
	Value string `bson:"value"`
}

type groupFilter struct {
	Name string `bson:"group_name"`
}

type group struct {
	Name    string    `bson:"group_name"` // composition doesn't work ):
	Content []content `bson:"content"`
}
