package models

type (
	Steps struct {
		SortNumber    int16  `json:"sort_number"`
		StepsBody     string `json:"steps_body"`
		StepsImageUrl string `json:"steps_image_url"`
		StepsVideoUrl string `json:"steps_video_url"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
	CreateStep struct {
		RecipeId      string `json:"recipe_id"`
		StepsBody     string `json:"steps_body"`
		StepsImageUrl string `json:"steps_image_url"`
		StepsVideoUrl string `json:"steps_video_url"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		CreatedBy  string `json:"created_by"`
		SortNumber int16  `json:"sort_number"`
	}
	UpdateStep struct {
		StepsBody     string `json:"steps_body"`
		StepsImageUrl string `json:"steps_image_url"`
		StepsVideoUrl string `json:"steps_video_url"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		UpdatedBy  string `json:"updated_by"`
	}
)
