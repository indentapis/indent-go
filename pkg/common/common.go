package common

const (
	// IndentPrefix is the beginning of all Indent types.
	IndentPrefix = "indent.com/"

	// IndentKind represents the type of a block.
	IndentKind = IndentPrefix + "kind"

	// LabelSchema identifies the schema for the structure that backs a block.
	LabelSchema = IndentPrefix + "schema"

	// LabelDescription provides an excerpt for a Resource.
	LabelDescription = "description"
)
