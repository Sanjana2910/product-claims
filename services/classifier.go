package services

import "strings"
// labels each claim category wise(business logic for classification)

// checks if claim text contains any of these words
var claimKeywords = map[string][]string{
	"Nutritional Claims": {
		"protein", "fibre", "fiber", "iron", "calcium",
		"sugar", "carbohydrate", "energy", "vitamin",
		"mineral", "sodium", "fat",
	},
	"Health & Wellness (Non-Medicinal)": {
		"supports immunity", "promotes digestion",
		"gut health", "wellness", "nourish",
	},
	"Safety Claims": {
		"safe for children", "non-toxic",
		"dermatologically tested", "hypoallergenic",
	},
	"Environmental / Sustainability": {
		"eco-friendly", "recyclable", "sustainable",
		"biodegradable", "green",
	},
	"Certification Claims": {
		"fssai", "iso", "bis approved", "certified",
		"compliant", "approved",
	},
	"Manufacturing / Origin": {
		"made in india", "made in", "handcrafted",
		"artisan", "locally sourced",
	},
	"Quality Claims": {
		"premium quality", "high-grade", "superior",
		"finest", "best quality",
	},
	"Performance / Efficacy": {
		"long-lasting", "more effective", "high strength",
		"fast acting", "powerful",
	},
	"Comparative Claims": {
		"better than", "30% better", "uses less",
		"more than", "compared to",
	},
	"Authenticity Claims": {
		"100% original", "authentic", "traditional recipe",
		"genuine", "pure",
	},
	"Awards & Endorsements": {
		"award-winning", "recognized by", "endorsed by",
		"recommended by", "approved by doctors",
	},
	"Price / Value Claims": {
		"best value", "lowest price", "big savings",
		"affordable", "value for money",
	},
	"Medical / Therapeutic (Restricted)": {
		"cures", "treats", "heals", "remedy",
		"therapeutic", "medicinal",
	},
}

// Classify takes a claim text and returns which category it belongs to

func Classify(text string) string {
	lower := strings.ToLower(text)

	for category, keywords := range claimKeywords {
		for _, keyword := range keywords { 
			if strings.Contains(lower, keyword) { //if input and data match
				return category 
			}
		}
	}

	return "" // returns empty string if no category matches

}