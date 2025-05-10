package models

type Movie struct {
	ContentID string `json:"content_id" validate:"required,lte=255"`
}

type Tv struct {
	ContentID string `json:"content_id" validate:"required,lte=255"`
}
