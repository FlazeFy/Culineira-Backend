package models

type Ingredients struct {
	Id             string `json:"id"`
	RecipeId       string `json:"recipes_id"`
	CommentBody    string `json:"comment_body"`
	CommentFileUrl string `json:"comment_file_url"`

	//Properties
	IsOptional bool  `json:"is_optional"`
	CreatedAt  int64 `json:"created_at"`
	CreatedBy  int64 `json:"created_by"`
	UpdatedAt  int64 `json:"updated_at"`
	UpdatedBy  int64 `json:"updated_by"`
}
