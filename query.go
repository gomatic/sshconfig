package sshconfig

//
type Query []phrase

//
type what int

//
const (
	keyQuery what = iota
	valueQuery
)

//
type phrase struct {
	element    what
	key, value string
}

//
func NewQuery(q string) Query {
	return Query{
		{
			element: keyQuery,
			key:     "host",
			value:   q,
		},
	}
}
