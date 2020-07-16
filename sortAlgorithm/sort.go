package SortAlgorithm

import "sort"

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n-1; j++ {
			if arr[i] > arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
	return arr
}

func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				min = j
			}
		}
		if min != i {
			tmp := arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}
	}
	return arr
}

type ServicesNode struct {
	ID    string
	Value int
}

func OsSort(services []ServicesNode) []ServicesNode {
	sort.Slice(services, func(i, j int) bool {
		return services[i].ID < services[j].ID
	})
	return services
}
