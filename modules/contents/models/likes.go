package models

type (
	Likes struct {
		Username string `json:"username"`

		//Properties
		CreatedAt string `json:"created_at"`
	}
	CreateLikes struct {
		RecipeId string `json:"recipes_id"`

		//Properties
		CreatedBy string `json:"created_by"`
	}
)
