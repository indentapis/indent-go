package oidcclaims

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	claims := decodeClaims(t, userInfoData)
	resource, err := claims.Resource()
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	assert.Equal(t, claims.Sub, resource.Id)
	assert.Equal(t, claims.Name, resource.DisplayName)
	assert.Equal(t, claims.Email, resource.Email)
}

func TestResourceNoName(t *testing.T) {
	claims := decodeClaims(t, userInfoData)
	claims.Name = ""

	resource, err := claims.Resource()
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	expected := claims.GivenName + " " + claims.MiddleName + " " + claims.FamilyName
	assert.Equal(t, expected, resource.DisplayName)
}

func TestResourceNoStandard(t *testing.T) {
	claims := decodeClaims(t, []byte(`{"non_standard": true}`))
	resource, err := claims.Resource()
	assert.Error(t, err)
	assert.Nil(t, resource)
}
