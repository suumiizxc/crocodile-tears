package industry

import "time"

type Company struct {
	Status            string    `json:"status"`
	IsCompanyCustomer uint      `json:"isCompanyCustomer"`
	FoundedDate       time.Time `json:"foundedDate"`
	IndustryID        uint      `json:"industryId"`
	CustCode          string    `json:"custCode"`
	Name              string    `json:"name"`
	OrgTypeID         uint      `json:"orgTypeId"`
	RegisterMaskCode  string    `json:"registerMaskCode"`
	ShortName         string    `json:"shortName"`
	IndustryName      string    `json:"industryName"`
	Name2             string    `json:"name2"`
	CustSegCode       string    `json:"custSegCode"`
	RegisterCode      string    `json:"registerCode"`
	ShortName2        string    `json:"shortName2"`
	CountryCode       string    `json:"countryCode"`
	Phone             string    `json:"phone"`
	Email             string    `json:"email"`
	CreatedBy         uint      `json:"createdBy"`
	ApprovedBy        uint      `json:"approvedBy"`
}
