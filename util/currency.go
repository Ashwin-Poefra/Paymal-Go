package util

// All supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// Checks if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
