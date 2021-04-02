package stats

import (
	"fmt"

	"github.com/sunatullo-gafurov/bank/pkg/types"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "grocery",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "sport",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "travel",
		},
	}

	fmt.Println(Avg(payments))
	//Output:
	//2000000
}

func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "grocery",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "sport",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "travel",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "grocery",
		},
	}

	fmt.Println(TotalInCategory(payments, "grocery"))

	//Output:
	//4000000
}
