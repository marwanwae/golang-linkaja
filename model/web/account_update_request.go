package web

type AccountUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100,min=1"`
}
