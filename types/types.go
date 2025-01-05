package types

type Item struct {
	shortDescription string `json:"shortDescription"`
	price            string `json:"shortDescription"`
}

type ReceiptPayload struct {
	retailer     string `json:"retailer"`
	purchaseDate string `json:"purchaseDate"`
	purchaseTime string `json:"purchaseTime"`
	items        []Item `json:"items"`
	total        string `json:total`
}
