package customer

type CountryCode struct {
	CountryCode string  `json:"countryCode"`
	Iso2        string  `json:"iso2"`
	Iso3        string  `json:"iso3"`
	Name        string  `json:"name"`
	Name2       string  `json:"name2"`
	RiskPercent float32 `json:"riskPercent"`
}
