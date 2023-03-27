package models

type Comments struct {
	Id             string `json:"id"`
	RecipeId       string `json:"recipes_id"`
	CommentBody    string `json:"comment_body"`
	CommentFileUrl string `json:"comment_file_url"`

	//Properties
	CreatedAt int64 `json:"created_at"`
	CreatedBy int64 `json:"created_by"`
	DeletedAt int64 `json:"deleted_at"`
	DeletedBy int64 `json:"deleted_by"`
}
