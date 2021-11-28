package stats

import (
	"fmt"

	"github.com/nwarior/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       101,
			Amount:   10_00,
			Category: "ресторан",
			Status:   types.StatusInProgress,
		},
		{
			ID:       103,
			Amount:   20_00,
			Category: "Burger",
			Status: types.StatusOk,
		},
		{
			ID:       101,
			Amount:   -10_00,
			Category: "nowhere",
			Status: types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))
	// Output: 1500
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       101,
			Amount:   10_00,
			Category: "ресторан",
			Status: types.StatusOk,
		},
		{
			ID:       103,
			Amount:   20_00,
			Category: "ресторан",
			Status: types.StatusInProgress,
		},
		{
			ID:       103,
			Amount:   -10_00,
			Category: "nowhere",
			Status: types.StatusFail,
		},
	}

	fmt.Println(TotalInCategory(payments, "ресторан"))
	// Output: 3000
}