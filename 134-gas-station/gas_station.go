package gasstation

// The Idea of this solution is that if the cost and of current index is way more than gas we have there is no point
// to continue from that index to the end

func canCompleteCircuit(gas []int, cost []int) int {

	total, tank, index := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff

		if tank < 0 {
			index = i + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}
	return index
}
