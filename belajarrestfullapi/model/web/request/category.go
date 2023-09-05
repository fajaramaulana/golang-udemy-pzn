package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=200"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=200"`
}
