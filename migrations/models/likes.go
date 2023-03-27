package models

type Likes struct {
	Id        string `json:"id"`
	RecipesId string `json:"recipes_id"`

	//Properties
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"` //user id
}
