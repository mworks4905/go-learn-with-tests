package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numberSlices ...[]int) []int {
	var allSum []int
	for _, slice := range numberSlices {
		allSum = append(allSum, Sum(slice))
	}
	return allSum
}

func SumAllTails(numberSlices ...[]int) []int {
	var allSumTails []int
	for _, slice := range numberSlices {
		if len(slice) == 0 {
			allSumTails = append(allSumTails, 0)
		} else {
			tailSlice := slice[1:]
			allSumTails = append(allSumTails, Sum(tailSlice))
		}
	}
	return allSumTails
} 

func MasterSum(numberSlices ...[]int) int {
	var masterSum int
	for _, slice := range numberSlices {
		masterSum += Sum(slice)
	}
	return masterSum
}