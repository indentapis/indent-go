package common

const (
	// LabelView contains a view to be returned.
	LabelView = "resource.indentapis.com/v1/ResourcePolicy.view"
)

const (
	// ViewPeople is the view id for listing all human users in a space.
	ViewPeople = "people"

	// ViewEmail is the view id for listing all user emails in a space.
	ViewEmail = "email"

	// ViewRequestable is the view id for listing requestable resources.
	ViewRequestable = "requestable"

	// ViewRecipients is the view id for listing potential recipients of access.
	ViewRecipients = "recipients"

	// ViewRequestableModal is the view id for listing requestable resources in a modal.
	ViewRequestableModal = ViewRequestable + ":modal"
)
