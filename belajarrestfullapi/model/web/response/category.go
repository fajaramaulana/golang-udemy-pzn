package response

type CategoryResponseById struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at" db:"createdAt"`
}

type CategoryResponseAll struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at" db:"createdAt"`
	UpdatedAt string `json:"updated_at" db:"updatedAt"`
}
