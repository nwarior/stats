﻿package stats

import (
	"github.com/nwarior/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	sum = 0

	var index int
	index = 0

	for _, payment := range payments {
		if payment.Amount > 0 && payment.Status != types.StatusFail {
			sum += types.Money(payment.Amount)
			index = index + 1
		}
	}
	return sum / types.Money(index)
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	var sum types.Money
	sum = 0
	
	for _, payment := range payments {
		if payment.Category == category && payment.Amount > 0 && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return sum
}