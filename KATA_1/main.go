package main

func Arr(n, x int)[]int{

	var result []int

	for i := 0; i <=n; i++{

		result = append(result, x * i)
	}
	return result
}


