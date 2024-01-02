package common

import (
	"time"
)

const (
	// MinLenReason is the shortest length accepted for a reason
	MinLenReason = 5
)

const (
	// KindPetition is the kind representing a Petition.
	KindPetition = "indent.v1.Petition"
)

const (
	// TimeoutPetitionList is the maximum time to wait for listing petitions.
	TimeoutPetitionList = 9 * time.Second
)

const (
	// PetitionStateOpen denotes the "open" petition state. See [docs](https://indent.dev/docs/petitions#petition-phases).
	PetitionStateOpen = "open"
	// PetitionStateGranted denotes the "granted" petition state. See [docs](https://indent.dev/docs/petitions#petition-phases).
	PetitionStateGranted = "granted"
	// PetitionStateRevoked denotes the "revoked" petition state. See [docs](https://indent.dev/docs/petitions#petition-phases).
	PetitionStateRevoked = "revoked"
	// PetitionStateClosed denotes the "closed" petition state. See [docs](https://indent.dev/docs/petitions#petition-phases).
	PetitionStateClosed = "closed"
)

const (
	// LabelPetitionReplacesID is the label containing the petition that a given petition replaces
	LabelPetitionReplacesID = "PetitionReplaces"
	// LabelPetitionReplacedByID is the label containing the petition that a given petition was replaced by
	LabelPetitionReplacedByID = "PetitionReplacedBy"
)
