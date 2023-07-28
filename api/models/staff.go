package models

type Staff struct {
	Id       string  `json:"id"`
	BranchId string  `json:"branch_id"`
	TarifId  string  `json:"tarif_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
}

type CreateStaff struct {
	BranchId string  `json:"branch_id"`
	TarifId  string  `json:"tarif_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
}

type UpdateStaff struct {
	Id       string  `json:"id"`
	BranchId string  `json:"branch_id"`
	TarifId  string  `json:"tarif_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
}

type StaffGetListRequest struct {
	Offset          int     `json:"offset"`
	Limit           int     `json:"limit"`
	SearchByBranch  string  `json:"search_by_branch"`
	SearchByTarifId string  `json:"search_by_tarif_id"`
	SearchByType    string  `json:"search_by_type"`
	SearchByName    string  `json:"search_by_name"`
	BalanceFrom     float64 `json:"balance_from"`
	BalanceTo       float64 `json:"balance_to"`
}

type StaffGetListResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}

type StaffPrimaryKey struct {
	Id string `json:"id"`
}
