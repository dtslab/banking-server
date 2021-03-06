package savingsproduct

type SavingsProduct struct {
	ProductName  string
	ProductID    string
	InterestRate map[int]float64
	ProductAlias string
}

var SavingsProductTypeName []string
var SavingsProductType map[string]SavingsProduct

func (p SavingsProduct) GetBaseInterestRateOfMonth(month int) float64 {
	return p.InterestRate[month]
}
