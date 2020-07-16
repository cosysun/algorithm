package SortAlgorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{4, 1, 7, 5, 9, 11}
	a = BubbleSort(a)
	fmt.Println(a)
}

func TestInsertSort(t *testing.T) {
	a := []int{4, 1, 7, 5, 9, 11}
	a = InsertSort(a)
	fmt.Println(a)
}

func TestSelectSort(t *testing.T) {
	a := []int{4, 1, 7, 5, 9, 11}
	a = SelectSort(a)
	fmt.Println(a)
}

func TestOsSort(t *testing.T) {
	for j := 0; j < 100; j++ {
		services := []ServicesNode{{ID: "production-tw-zfirst-trec-172.16.20.58-0-45880", Value: 1}, {ID: "production-tw-zfirst-trec-172.16.20.57-1-45881", Value: 2},
			{ID: "production-tw-zfirst-trec-172.16.20.55-1-45881", Value: 3}, {ID: "production-tw-zfirst-trec-172.16.20.55-0-45880", Value: 4},
			{ID: "production-tw-zfirst-trec-172.16.20.63-0-45880", Value: 5}, {ID: "production-tw-zfirst-trec-172.16.20.63-0-45880", Value: 6}}
		services = OsSort(services)
		for i := 0; i < len(services); i++ {
			fmt.Println(services[i].Value)
		}
	}

}
