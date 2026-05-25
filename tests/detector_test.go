package tests

import (
	"testing"
	"product-claims/models"
	"product-claims/services"
)

func TestDetectBasicClaims(t *testing.T) {
	product := &models.Product{
		ShortDescription: "Source of fibre and iron. Safe for children.",
		AboutItems: []string{
			"Source of Fibre & Iron",
		},
		ComplianceInfo: models.ComplianceInfo{
			CountryOfOrigin: "India",
			FssaiNo:         "10012011000168",
		},
		Attributes: models.Attributes{
			NutritionalDetails: "Protein 8.0g, Fiber 5.0g, Iron 3.70mg",
			DietType:           "Vegetarian",
		},
	}

	got := services.Detect(product)

	if len(got) == 0 {
		t.Fatal("expected claims, got none")
	}

	t.Logf("found %d claims!", len(got))
	for _, c := range got {
		t.Logf("  → %s: %s", c.ClaimType, c.ClaimValue)
	}
}

func TestDetectNoDuplicates(t *testing.T) {
	product := &models.Product{
		ShortDescription: "Source of fibre and iron.",
		AboutItems: []string{
			"Source of Fibre & Iron",
		},
		ComplianceInfo: models.ComplianceInfo{},
		Attributes:     models.Attributes{},
	}

	got := services.Detect(product)

	seen := map[string]bool{}
	for _, c := range got {
		key := c.ClaimType + c.ClaimValue
		if seen[key] {
			t.Fatalf("duplicate claim found: %s", c.ClaimValue)
		}
		seen[key] = true
	}

	t.Log("no duplicates found!")
}

func TestDetectEmptyProduct(t *testing.T) {
	product := &models.Product{}

	got := services.Detect(product)

	if len(got) != 0 {
		t.Fatalf("expected no claims, got %d", len(got))
	}

	t.Log("empty product returns no claims!")
}
