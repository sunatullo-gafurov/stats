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

// CategoriesAvgs calculates average payment for each category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	average := make(map[types.Category]types.Money)
	numOfPayments := make(map[types.Category]types.Money)
	sumOfPayments := make(map[types.Category]types.Money)

	for _, payment := range payments {
		sumOfPayments[payment.Category] += payment.Amount
		numOfPayments[payment.Category]++
	}

	for key, value := range sumOfPayments {
		average[key] = value / numOfPayments[key]
	}

	return average
}

func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {

	result := second

	for key, value := range first {
		result[key] -= value
	}

	return result
}
