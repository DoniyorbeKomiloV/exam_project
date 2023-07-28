package models

type Tarif struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	AmountForCash float64 `json:"amount_for_cash"`
	AmountForCard float64 `json:"amount_for_card"`
}

type CreateTarif struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	AmountForCash float64 `json:"amount_for_cash"`
	AmountForCard float64 `json:"amount_for_card"`
}

type UpdateTarif struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	AmountForCash float64 `json:"amount_for_cash"`
	AmountForCard float64 `json:"amount_for_card"`
}

type TarifGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TarifGetListResponse struct {
	Count  int      `json:"count"`
	Tarifs []*Tarif `json:"tarifs"`
}

type TarifPrimaryKey struct {
	Id string `json:"id"`
}
