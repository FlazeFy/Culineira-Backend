package models

type Steps struct {
	Id            string `json:"id"`
	RecipeId      string `json:"recipes_id"`
	SortNumber    int16  `json:"sort_number"`
	StepsBody     string `json:"steps_body"`
	StepsImageUrl string `json:"comment_image_url"`
	StepsVideoUrl string `json:"comment_video_url"`

	//Properties
	IsOptional bool   `json:"is_optional"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
}
