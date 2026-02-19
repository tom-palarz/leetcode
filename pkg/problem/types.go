package problem

type ProblemMetadata struct {
	fullName string
	status   string // TODO new type (str enum)
	problem  Problem
}

// TODO: Get this from GraphQL?
type Problem struct {
	id          int
	slug        string
	difficulty  string
	description string
	companies   []string
}
