package models

type Countries struct {
	Id               string `json:"id"`
	CountryCode      string `json:"country_code"`
	CountryName      string `json:"country_name"`
	CountryContinent string `json:"country_continent"`

	//Properties
	IsActive  bool  `json:"is_active"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
