package testProject

import (
	"github.com/shopspring/decimal"
	"time"
)

func HelloWorld() string {
	return "Hello World"
}

// CalculatePrice converts gross price in grosz to net and gross price in zloty, rounded to two decimal places.
func CalculatePrice(price, tax int64) (net, gross float64) {
	taxRate := decimal.NewFromInt(tax).Mul(decimal.NewFromFloat(0.01)).Add(decimal.NewFromInt(1))
	amountGross := decimal.NewFromInt(price).Mul(decimal.NewFromFloat(0.01))
	amountNet := amountGross.Div(taxRate)
	// Kwoty wykazywane w fakturze zaokrągla się do pełnych groszy, przy czym końcówki poniżej 0,5 grosza pomija się,
	// a końcówki od 0,5 grosza zaokrągla się do 1 grosza. (art. 106e ustawy o VAT)
	net, _ = amountNet.Round(2).Float64()
	gross, _ = amountGross.Round(2).Float64()
	return net, gross
}

//go:generate mock
type Clock interface {
	Now() time.Time
}
