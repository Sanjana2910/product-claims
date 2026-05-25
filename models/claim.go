package models

// Claim is one identified claim from the product
type Claim struct {
	ID         string `json:"id"`
	ClaimType  string `json:"claimType"`
	ClaimValue string `json:"claimValue"`
	Status     string `json:"status"`
}