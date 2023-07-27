package models

type Transaction struct {
	Id         string  `json:"id"`
	SalesId    string  `json:"sales_id"`
	Type       string  `json:"type"`
	SourceType string  `json:"source_type"`
	Text       string  `json:"text"`
	Amount     float64 `json:"amount"`
	StaffId    string  `json:"staff_id"`
}

type CreateTransaction struct {
	SalesId    string  `json:"sales_id"`
	Type       string  `json:"type"`
	SourceType string  `json:"source_type"`
	Text       string  `json:"text"`
	Amount     float64 `json:"amount"`
	StaffId    string  `json:"staff_id"`
}

type UpdateTransaction struct {
	Id         string  `json:"id"`
	SalesId    string  `json:"sales_id"`
	Type       string  `json:"type"`
	SourceType string  `json:"source_type"`
	Text       string  `json:"text"`
	Amount     float64 `json:"amount"`
	StaffId    string  `json:"staff_id"`
}

type TransactionGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type TransactionGetListResponse struct {
	Count        int            `json:"count"`
	Transactions []*Transaction `json:"transactions"`
}

type TransactionPrimaryKey struct {
	Id string `json:"id"`
}
