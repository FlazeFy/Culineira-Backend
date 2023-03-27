package models

type Steps struct {
	Id            string `json:"id"`
	RecipeId      string `json:"recipes_id"`
	StepsBody     string `json:"steps_body"`
	StepsImageUrl string `json:"comment_image_url"`
	StepsVideoUrl string `json:"comment_video_url"`

	//Properties
	IsOptional bool  `json:"is_optional"`
	CreatedAt  int64 `json:"created_at"`
	CreatedBy  int64 `json:"created_by"`
	UpdatedAt  int64 `json:"updated_at"`
	UpdatedBy  int64 `json:"updated_by"`
}
