package models

type (
	Comments struct {
		CommentBody    string `json:"comment_body"`
		CommentFileUrl string `json:"comment_file_url"`
		Username       string `json:"username"`

		//Properties
		CreatedAt string `json:"created_at"`
	}
	CreateComments struct {
		RecipeId       string `json:"recipes_id"`
		CommentBody    string `json:"comment_body"`
		CommentFileUrl string `json:"comment_file_url"`

		//Properties
		CreatedBy string `json:"created_by"`
	}
)
