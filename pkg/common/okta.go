package common

const (
	// OktaPrefix is the beginning of all of Okta kinds.
	OktaPrefix = "okta.v1."
)

const (
	// KindOktaUser is for resources that represent Okta users.
	KindOktaUser = OktaPrefix + "User"

	// KindOktaAppUser is for resources that represent Okta app users.
	KindOktaAppUser = OktaPrefix + "AppUser"

	// KindOktaGroup is for resources that represent Okta groups.
	KindOktaGroup = OktaPrefix + "Group"
)

const (
	// LabelOktaID label containing the Okta ID related to a Resource.
	LabelOktaID = "oktaId"
)
