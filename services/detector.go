package services // scans the prod text and finds claim

import (
	"strings"
	"product-claims/models"
)

//holds text and where it came from
type textToCheck struct {
	source string
	text   string
}

// Detect scans all fields of a product and returns found claims
func Detect(product *models.Product) []models.Claim {
	var foundClaims []models.Claim //empty, will fill as we find claims

	// collect all text we want to scan
	var allTexts []textToCheck
	if product.ShortDescription != "" { // cannot be empty
    sentences := strings.Split(product.ShortDescription, ".")
    for _, sentence := range sentences {
        cleaned := strings.TrimSpace(sentence)
        if cleaned != "" {// skips empty sentences after trimming
            allTexts = append(allTexts, textToCheck{
                source: "shortDescription",
                text:   cleaned,
            })// adds cleaned sentence to allTexts
        }
    }
}

	// add short description
	// if product.ShortDescription != "" {
	// 	allTexts = append(allTexts, textToCheck{
	// 		source: "shortDescription",
	// 		text:   product.ShortDescription,
	// 	})
	// }

	// add each about item
	for _, item := range product.AboutItems {
		if item != "" {
			allTexts = append(allTexts, textToCheck{
				source: "aboutItems",
				text:   item,
			})
		}
	}

	// add country of origin
	if product.ComplianceInfo.CountryOfOrigin != "" {
		allTexts = append(allTexts, textToCheck{
			source: "complianceInfo",
			text:   "Made in " + product.ComplianceInfo.CountryOfOrigin,
		})
	}

	// if fssai number exists, it means fssai certified
	if product.ComplianceInfo.FssaiNo != "" {
		allTexts = append(allTexts, textToCheck{
			source: "complianceInfo",
			text:   "FSSAI certified",
		})
	}

	// add nutritional details
	if product.Attributes.NutritionalDetails != "" {
		allTexts = append(allTexts, textToCheck{
			source: "attributes",
			text:   product.Attributes.NutritionalDetails,
		})
	}

	// add diet type
	if product.Attributes.DietType != "" {
		allTexts = append(allTexts, textToCheck{
			source: "attributes",
			text:   product.Attributes.DietType,
		})
	}

	//go through every text and try to classify it

	seen := map[string]bool{} // avoid duplicate claims

	for _, item := range allTexts {
		category := Classify(item.text)// classify defined in classifier.go, returns category of claim based on text

		// skip if no category found
		if category == "" {
			continue
		}

		// skip if we already found this exact claim
		key := strings.ToLower(strings.TrimSpace(item.text))

		if seen[key] {
			continue
		}
		seen[key] = true

		foundClaims = append(foundClaims, models.Claim{
			ClaimType:  category,
			ClaimValue: item.text,
			Status:     "IDENTIFIED",
		})
	}

	return foundClaims
}