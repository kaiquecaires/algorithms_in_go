package optimalfreelancing

import "sort"

func OptimalFreelancing(jobs []map[string]int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i]["payment"] > jobs[j]["payment"]
	})
	var timeline [7]int

	for _, job := range jobs {
		maxDeadline := job["deadline"] - 1

		if maxDeadline > 6 {
			maxDeadline = 6
		}

		for i := maxDeadline; i >= 0; i-- {
			if timeline[i] == 0 {
				timeline[i] = job["payment"]
				break
			}
		}
	}

	amount := 0

	for _, n := range timeline {
		amount += n
	}

	return amount
}
