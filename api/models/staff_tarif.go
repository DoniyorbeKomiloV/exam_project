package models

type Tarif struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateTarif struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type UpdateTarif struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type TarifGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type TarifGetListResponse struct {
	Count  int      `json:"count"`
	Tarifs []*Tarif `json:"tarifs"`
}

type TarifPrimaryKey struct {
	Id string `json:"id"`
}
