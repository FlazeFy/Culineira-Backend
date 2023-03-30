package models

type (
	UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	CreateUser struct {
		CountryCode string `json:"country_code"`
		UserName    string `json:"username"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Job         string `json:"job"`
		Email       string `json:"email"`
		Password    string `json:"password"`

		//Properties
		IsActive bool `json:"is_active"`
	}
)
