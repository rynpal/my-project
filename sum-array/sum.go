package main

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
