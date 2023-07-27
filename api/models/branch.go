package models

type Branch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CreateBranch struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateBranch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type BranchGetListRequest struct {
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	SearchName    string `json:"search_name"`
	SearchAddress string `json:"search_address"`
}

type BranchGetListResponse struct {
	Count    int       `json:"count"`
	Branches []*Branch `json:"branches"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}
