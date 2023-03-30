package models

type (
	StepInfo struct {
		StepsImageUrl string `json:"steps_image_url"`
		StepsVideoUrl string `json:"steps_video_url"`

		//Properties
		IsOptional bool `json:"is_optional"`
	}
	Steps struct {
		SortNumber int16  `json:"sort_number"`
		StepsBody  string `json:"steps_body"`
		StepInfo

		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	CreateStep struct {
		RecipeId  string `json:"recipe_id"`
		StepsBody string `json:"steps_body"`
		StepInfo

		CreatedBy  string `json:"created_by"`
		SortNumber int16  `json:"sort_number"`
	}
	UpdateStep struct {
		StepsBody string `json:"steps_body"`
		StepInfo

		UpdatedBy string `json:"updated_by"`
	}
)
