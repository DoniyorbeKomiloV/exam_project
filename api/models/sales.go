package models

type Sales struct {
	Id              string  `json:"id"`
	BranchId        string  `json:"branch_id"`
	ShopAssistantId string  `json:"shop_assistant_id"`
	CashierId       string  `json:"cashier_id"`
	Price           float64 `json:"price"`
	PaymentType     string  `json:"payment_type"`
	Status          string  `json:"status"`
	ClientName      string  `json:"client_name"`
}

type CreateSales struct {
	BranchId        string  `json:"branch_id"`
	ShopAssistantId string  `json:"shop_assistant_id"`
	CashierId       string  `json:"cashier_id"`
	Price           float64 `json:"price"`
	PaymentType     string  `json:"payment_type"`
	Status          string  `json:"status"`
	ClientName      string  `json:"client_name"`
}

type UpdateSales struct {
	Id              string  `json:"id"`
	BranchId        string  `json:"branch_id"`
	ShopAssistantId string  `json:"shop_assistant_id"`
	CashierId       string  `json:"cashier_id"`
	Price           float64 `json:"price"`
	PaymentType     string  `json:"payment_type"`
	Status          string  `json:"status"`
	ClientName      string  `json:"client_name"`
}

type SalesGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type SalesGetListResponse struct {
	Count int      `json:"count"`
	Sales []*Sales `json:"sales"`
}

type SalesPrimaryKey struct {
	Id string `json:"id"`
}
