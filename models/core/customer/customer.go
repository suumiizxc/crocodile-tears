package customer

import "time"

type Customer struct {
	CustCode          string    `json:"custCode"`
	CustSegCode       string    `json:"custSegCode"`
	SexCode           uint      `json:"sexCode"`
	TaxExemption      uint      `json:"taxExemption"`
	Status            uint      `json:"status"`
	IsCompanyCustomer uint      `json:"isCompanyCustomer"`
	IndustryID        uint      `json:"industryId"`
	BirthPlaceID      uint      `json:"birthPlaceId"`
	FamilyName        string    `json:"familyName"`
	FamilyName2       string    `json:"familyName2"`
	LastName          string    `json:"lastName"`
	LastName2         string    `json:"lastName2"`
	FirstName         string    `json:"firstName"`
	FirstName2        string    `json:"firstName2"`
	ShortName         string    `json:"shortName"`
	ShortName2        string    `json:"shortName2"`
	RegisterMaskCode  string    `json:"registerMaskCode"`
	RegisterCode      string    `json:"registerCode"`
	BirthDate         time.Time `json:"birthDate"`
	Mobile            string    `json:"mobile"`
	CountryCode       string    `json:"countryCode"`
	EmploymentID      uint      `json:"employmentId"`
	Email             string    `json:"email"`
	IndustryName      string    `json:"industryName"`
	CatID             uint      `json:"catId"`
	TitleID           uint      `json:"titleId"`
	NationalityID     uint      `json:"nationalityId"`
	EthnicGroupID     uint      `json:"ethnicGroupId"`
	LangCode          string    `json:"langCode"`
	MaritalStatus     uint      `json:"maritalStatus"`
	BirthPlaceName    string    `json:"birthPlaceName"`
	BirthPlaceDetail  string    `json:"birthPlaceDetail"`
	EducationID       uint      `json:"educationId"`
	Phone             string    `json:"phone"`
	Fax               string    `json:"fax"`
	CreatedBy         uint      `json:"createdBy"`
	ApprovedBy        uint      `json:"approvedBy"`
	CompanyCode       string    `json:"companyCode"`
	IsVatPayer        uint      `json:"isVatPayer"`
}
