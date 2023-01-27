package oidcclaims

// Address represents a physical mailing address. It's defined at https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim.
type Address struct {
	// Full mailing address, formatted for display or use on a mailing label. This field MAY contain multiple lines, separated by newlines.
	// Newlines can be represented either as a carriage return/line feed pair ("\r\n") or as a single line feed character ("\n").
	Formatted string `json:"formatted,omitempty"`
	// Full street address component, which MAY include house number, street name, Post Office Box, and multi-line extended street address
	// information. This field MAY contain multiple lines, separated by newlines. Newlines can be represented either as a carriage return/line
	// feed pair ("\r\n") or as a single line feed character ("\n").
	StreetAddress string `json:"street_address,omitempty"`
	// City or locality component.
	Locality string `json:"locality,omitempty"`
	// State, province, prefecture, or region component.
	Region string `json:"region,omitempty"`
	// Zip code or postal code component.
	PostalCode string `json:"postal_code,omitempty"`
	// Country name component.
	Country string `json:"country,omitempty"`
}
