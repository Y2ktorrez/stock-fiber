package dtos

type CreateMovimentRequest struct {
	Type        string `json:"type" validate:"required"`
	Amount      int64  `json:"amount" validate:"required"`
	ProductSlug string `json:"product_slug" validate:"required"`
	UserSlug    string `json:"user_slug" validate:"required"`
}

type MovimentResponse struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Date        string `json:"date"`
	ProductSlug string `json:"product_slug"`
	UserSlug    string `json:"user_slug"`
}
