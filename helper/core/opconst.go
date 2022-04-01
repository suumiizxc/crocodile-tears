package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PolarisConst struct {
	polaris_company string
	polaris_lang    string
	polaris_role    string
	polaris_url     string
}

func (pct *PolarisConst) GetLang() string {
	return pct.polaris_lang
}

func (pct *PolarisConst) GetCompany() string {
	return pct.polaris_company
}
func (pct *PolarisConst) GetRole() string {
	return pct.polaris_role
}
func (pct *PolarisConst) GetUrl() string {
	return pct.polaris_url
}

func (pct *PolarisConst) GetPolarisConst(name string) string {
	var Info = map[string]string{
		"polaris_company": pct.polaris_company,
		"polaris_lang":    pct.polaris_lang,
		"polaris_role":    pct.polaris_role,
		"polaris_url":     pct.polaris_url,
	}

	return Info[name]
}

var PC PolarisConst

const (
	OP_test = 123456
)

func PolarisVar() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Failed env : %v", err))
	}
	pc := PolarisConst{}
	pc.polaris_company = os.Getenv("POLARIS_COMPANY")
	pc.polaris_lang = os.Getenv("POLARIS_LANG")
	pc.polaris_role = os.Getenv("POLARIS_ROLE")
	pc.polaris_url = os.Getenv("POLARIS_URL")
	PC = pc
}
