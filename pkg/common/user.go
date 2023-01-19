package common

const (
	// KindUser is the resource kind for an Indent user resource.
	KindUser = "indent.blocks.v1.User"

	// KindBotUser is the resource kind for an Indent bot user resource.
	KindBotUser = "indent.v1.Bot"

	// KindEmail is a representation of an Email address.
	KindEmail = "email"
)

const (
	// KindRole is a Role within Indent that can be assigned to a user.
	KindRole = "indent.v1.Role"

	// KindSpacedRole is a Role specific to a certain space.
	KindSpacedRole = "indent.v1.SpacedRole"
)

const (
	// LabelRole is used to indicate the role of a spaced role.
	LabelRole = IndentPrefix + "role"
	// LabelResourceLink contains the resource name for the external resource
	// "backing" this Indent user.
	LabelResourceLink = ResourcePrefix + "link"
	// LabelIsEditor indicates whether the user resource has editor Role.
	LabelIsEditor = ResourcePrefix + "isEditor"
)
