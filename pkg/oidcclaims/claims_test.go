package oidcclaims

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	standardClaims = &Standard{
		Sub:               "248289761001",
		Name:              "Jane Doe",
		GivenName:         "Jane",
		FamilyName:        "Doe",
		MiddleName:        "He",
		Nickname:          "Janie",
		PreferredUsername: "jane.doe",
		Profile:           "https://example.com/jane.doe",
		Picture:           "https://example.com/jane.doe/img.png",
		Website:           "https://jane.he.doe.example.net",
		Email:             "jane.doe@example.com",
		EmailVerified:     true,
		Zoneinfo:          "America/Los_Angeles",
		Locale:            "en-US",
		PhoneNumber:       "+1 (310) 123-4567",
		Address: &Address{
			StreetAddress: "1234 Hollywood Blvd.",
			Locality:      "Los Angeles",
			Region:        "CA",
			PostalCode:    "90210",
			Country:       "US",
		},
		UpdatedAt: 13663990,
	}
	userInfoData, _ = json.Marshal(&standardClaims)
)

func TestStandardClaims(t *testing.T) {
	claims := decodeClaims(t, userInfoData)
	assert.Equal(t, standardClaims, claims.Standard)
}

func TestNoStandardClaims(t *testing.T) {
	noStandardData := []byte(`{
	"notAStandardField": true
}`)
	claims := decodeClaims(t, noStandardData)
	assert.Nil(t, claims.Standard)
	assert.Equal(t, map[string]interface{}{"notAStandardField": true}, claims.All)
}

func TestSomeStandardClaims(t *testing.T) {
	noStandardData := []byte(`{
	"notAStandardField": true,
	"given_name": "First",
	"family_name": "Last"
}`)
	claims := decodeClaims(t, noStandardData)
	assert.Equal(t, &Standard{GivenName: "First", FamilyName: "Last"}, claims.Standard)
	assert.Len(t, claims.All, 3)
}

func decodeClaims(t *testing.T, data []byte) *Claims {
	t.Helper()

	claims := new(Claims)
	err := json.Unmarshal(data, claims)
	assert.NoError(t, err)
	return claims
}
