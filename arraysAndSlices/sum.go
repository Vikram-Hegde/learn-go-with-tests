package main

func Sum(arr []int) (sum int) {
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	for _, number := range arr {
		sum += number
	}
	return
}

func SumAll(arr1, arr2 []int) (sum []int) {
	sum1 := 0
	sum2 := 0
	for _, num1 := range arr1 {
		sum1 += num1
	}
	for _, num2 := range arr2 {
		sum2 += num2
	}
	sum = append(sum, sum1, sum2)
	return
}