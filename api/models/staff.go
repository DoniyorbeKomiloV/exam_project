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
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type StaffGetListResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}

type StaffPrimaryKey struct {
	Id string `json:"id"`
}
