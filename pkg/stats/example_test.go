package stats

import (
	"fmt"

	"github.com/sunatullo-gafurov/bank/v2/pkg/types"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "grocery",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "sport",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "travel",
			Status:   "FAIL",
		},
	}

	fmt.Println(Avg(payments))
	//Output:
	//1500000
}

func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "grocery",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "sport",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "travel",
			Status:   "FAIL",
		},
	}

	fmt.Println(TotalInCategory(payments, "grocery"))

	//Output:
	//1000000
}
