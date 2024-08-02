package compute

type Query struct {
	Command   int
	Arguments []string
}

func NewQuery(command int, args []string) Query {
	return Query{command, args}
}
