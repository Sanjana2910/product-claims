package tests

import (
    "testing"
    "product-claims/services"
)

func TestNutritionalClaim(t *testing.T) {
    input := "Source of Fibre & Iron"
    want  := "Nutritional Claims"

    got := services.Classify(input)
    if got != want {
        t.Fatalf("expected %s, got %s", want, got)
    }
    t.Log("nutritional claim passed!")
}

func TestSafetyClaim(t *testing.T) {
    input := "Safe for children"
    want  := "Safety Claims"

    got := services.Classify(input)
    if got != want {
        t.Fatalf("expected %s, got %s", want, got)
    }
    t.Log("safety claim passed!")
}

func TestCertificationClaim(t *testing.T) {
    input := "FSSAI certified"
    want  := "Certification Claims"

    got := services.Classify(input)
    if got != want {
        t.Fatalf("expected %s, got %s", want, got)
    }
    t.Log("certification claim passed!")
}

func TestManufacturingClaim(t *testing.T) {
    input := "Made in India"
    want  := "Manufacturing / Origin"

    got := services.Classify(input)
    if got != want {
        t.Fatalf("expected %s, got %s", want, got)
    }
    t.Log("manufacturing claim passed!")
}

func TestNoMatchClaim(t *testing.T) {
    input := "random text with no claims"
    want  := ""

    got := services.Classify(input)
    if got != want {
        t.Fatalf("expected empty string, got %s", got)
    }
    t.Log("no match test passed!")
}
