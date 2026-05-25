package models
// shape of Maggi JSON in go
//Info about the product
type ComplianceInfo struct{
	CountryOfOrigin string `json:"country_of_origin"`
	FssaiNo string `json:"fssai_no"`
	ManufacturingDate string `json:"manufacturing_date"`


}
//Attributes of the product
type Attributes struct{
		DietType            string `json:"diet_type"`
	NutritionalDetails  string `json:"nutritional_details"`
	Ingredients         string `json:"ingredients"`
	ReadyToCook         bool   `json:"ready_to_cook_yn"`

}

// Claim represents a single product claim
type Product struct {
	ID              string         `json:"id"`
	McrID           string         `json:"mcrId"`
	Title           string         `json:"title"`
	Brand           string         `json:"brand"`
	CategoryName    string         `json:"categoryName"`
	ShortDescription string        `json:"shortDescription"`
	AboutItems      []string       `json:"aboutItems"`
	ComplianceInfo  ComplianceInfo `json:"complianceInfo"`
	Attributes      Attributes     `json:"attributes"`
	Claims          []Claim        `json:"claims"`
}
