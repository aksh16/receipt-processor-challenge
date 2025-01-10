package types

type ReceiptStore interface {
	GetPoints(receipt_id int64) (int64, error)
	AddReceipt(receipt string) error
}

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
