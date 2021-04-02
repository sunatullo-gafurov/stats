package stats

import "github.com/sunatullo-gafurov/bank/pkg/types"

// Avg calculates average payment
func Avg(payments []types.Payment) types.Money {
	count := types.Money(0)
	average := types.Money(0)
	for _, payment := range payments {
		average += payment.Amount
		count++
	}

	return average / count
}

// TotalInCategory calculates sum of all payments with given category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	sum := types.Money(0)

	for _, payment := range payments {

		if payment.Category == category {
			sum += payment.Amount
		}
	}

	return sum
}
