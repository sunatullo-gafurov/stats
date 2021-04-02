package stats

import "github.com/sunatullo-gafurov/bank/v2/pkg/types"

// Avg calculates average payment
func Avg(payments []types.Payment) types.Money {
	count := types.Money(0)
	average := types.Money(0)
	for _, payment := range payments {

		if payment.Status != types.StatusFail {
			average += payment.Amount
			count++
		}
	}

	return average / count
}

// TotalInCategory calculates sum of all payments with given category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	sum := types.Money(0)

	for _, payment := range payments {

		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return sum
}
