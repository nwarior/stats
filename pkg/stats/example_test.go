package stats

import (
	"github.com/nwarior/bank/pkg/types"
	"fmt"
)

func ExampleAvg(){
	payments := []types.Payment {
		{
			ID: 101,
			Amount: 10_00,
			Category: "ресторан",
		},
		{
			ID: 103,
			Amount: 20_00,
			Category: "Burger",
		},
		{
			ID: 101,
			Amount: -10_00,
			Category: "nowhere",
		},
	}

	fmt.Println(Avg(payments))
	// Output: 1500
}

func ExampleTotalInCategory(){
	payments := []types.Payment {
		{
			ID: 101,
			Amount: 10_00,
			Category: "ресторан",
		},
		{
			ID: 103,
			Amount: 20_00,
			Category: "ресторан",
		},
		{
			ID: 103,
			Amount: -10_00,
			Category: "nowhere",
		},
	}

	fmt.Println(TotalInCategory(payments, "ресторан"))
	// Output: 3000
}