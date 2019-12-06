func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	amounts := make([]int, amount+1)

	for amt, v := range amounts {
		if amt != 0 && v == 0 {
			continue
		}
		for _, coin := range coins {
			newAmt := amt + coin
			if newAmt > amount {
				continue
			}
			if amounts[newAmt] > v+1 || amounts[newAmt] == 0 {
				amounts[newAmt] = v + 1
			}
		}
	}

	if amounts[amount] == 0 {
		return -1
	}
	return amounts[amount]
}