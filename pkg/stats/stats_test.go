package stats

import (
	"reflect"
	"testing"

	"github.com/sunatullo-gafurov/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {

	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 200_000_00},
		{ID: 2, Category: "mobile", Amount: 1_000_00},
		{ID: 3, Category: "entertainment", Amount: 2_000_00},
		{ID: 4, Category: "food", Amount: 5_000_00},
		{ID: 5, Category: "auto", Amount: 60_000_00},
		{ID: 6, Category: "entertainment", Amount: 10_000_00},
		{ID: 7, Category: "food", Amount: 8_000_00},
		{ID: 8, Category: "mobile", Amount: 500_00},
		{ID: 9, Category: "auto", Amount: 100_000_00},
	}

	expected := map[types.Category]types.Money{
		"auto":          120_000_00,
		"mobile":        750_00,
		"entertainment": 6_000_00,
		"food":          6_500_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected this %v, got this %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {

	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got this %v", expected, result)
	}
}
