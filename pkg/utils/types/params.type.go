package types

type FilterParams struct {
	Genres     string `json:"genres"`
	Providers  string `json:"providers"`
	Page       string `json:"page"`
	From       string `json:"from"`
	To         string `json:"to"`
	Sort       string `json:"sort"`
	RuntimeGte string `json:"runtime_gte"`
	RuntimeLte string `json:"runtime_lte"`
}
