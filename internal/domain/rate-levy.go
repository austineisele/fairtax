package domain



type LevyData struct {
	FiscalYearEnding            int     `json:"fiscal_year_ending"`
	RollYear                    int     `json:"roll_year"`
	SwisCode                    string  `json:"swis_code"`
	Municipality                string  `json:"municipality"`
	County                      string  `json:"county"`
	SchoolCode                  string  `json:"school_code"`
	SchoolName                  string  `json:"school_name"`
	TypeOfValue                 string  `json:"type_of_value_on_whichtax_rates_are_applied"`
	CountyTaxLevy               int     `json:"county_tax_levy"`
	CountyTaxRateOutsideVillage float32 `json:"county_tax_rate_outside_village_per_1000_assessed_value"`
	CountyTaxRateInsideVillage  float32 `json:"county_tax_rate_inside_village_per_1000_assessed_value"`
	MunicipalTaxLevy            float32 `json:"municipality_tax_levy"`
	MunicipalTaxOutsideVillage  float32 `json:"municipal_tax_rate_outside_village_per_1000_assessed_value"`
	MunicipalTaxInsideVillage   float32 `json:"municipal_tax_rate_inside_village_per_1000_assessed_value"`
	SchoolTaxLevy               float32 `json:"school_district_tax_levy"`
	SchoolTaxRatePer1000        float32 `json:"school_district_tax_rate_per_1000_assessed_value"`
}
