package common

const (
	// ResourcePrefix begins each resource label.
	ResourcePrefix = IndentPrefix + "resource/"
)

const (
	// LabelResourceEmail contains the email associated with Resource.
	LabelResourceEmail = ResourcePrefix + "email"

	// LabelResourceKind contains the kind associated with Resource.
	LabelResourceKind = ResourcePrefix + "kind"

	// LabelResourceID contains the ID associated with Resource.
	LabelResourceID = ResourcePrefix + "id"

	// LabelResourceIconURL contains the icon URL associated with Resource.
	LabelResourceIconURL = ResourcePrefix + "icon"
)

const (
	// ResourceSchema tracks the version of the Resource schema.
	ResourceSchema = "indent.audit.v1beta1.Resource"
)
