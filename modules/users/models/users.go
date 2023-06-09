package models

type (
	Users struct {
		CountryCode string `json:"country_code"`
		SlugName    string `json:"slug_name"`
		UserName    string `json:"username"`
		FullName    string `json:"full_name"`

		//Properties
		CreatedAt string `json:"created_at"`
	}
	UserDetail struct {
		CountryCode string `json:"country_code"`
		SlugName    string `json:"slug_name"`
		UserName    string `json:"username"`
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		Job         string `json:"job"`

		//Properties
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	UpdateUser struct {
		CountryCode string `json:"country_code"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Job         string `json:"job"`

		//Properties
		UpdatedBy string `json:"updated_by"`
	}
)
