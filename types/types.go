package types

type ReceiptStore interface {
	GetPoints(receipt_id uint64) (uint64, error)
	AddPoints(points uint64) (uint64, error)
	CheckDB()
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptPayload struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}
