package models

type NaturalPerson struct {
	ID            int     `json:"id"`
	MonthlyIncome float64 `json:"monthly_income"`
	Age           int     `json:"age"`
	FullName      string  `json:"full_name"`
	PhoneNumber   string  `json:"phone_number"`
	Email         string  `json:"email"`
	Category      string  `json:"category"`
	Balance       float64 `json:"balance"`
}

type LegalPerson struct {
	ID             int     `json:"id"`
	AnnualRevenue  float64 `json:"annual_revenue"`
	Age            int     `json:"age"`
	TradeName      string  `json:"trade_name"`
	PhoneNumber    string  `json:"phone_number"`
	CorporateEmail string  `json:"corporate_email"`
	Category       string  `json:"category"`
	Balance        float64 `json:"balance"`
}
