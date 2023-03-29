package models

type Users struct {
	Id          string `json:"id"`
	CountryCode string `json:"country_code"`
	SlugName    string `json:"slug_name"`
	UserName    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Job         string `json:"job"`
	Email       string `json:"email"`
	Password    string `json:"password"`

	//Properties
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"` //user id
}
