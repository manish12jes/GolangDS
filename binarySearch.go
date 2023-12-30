package main

import "fmt"


func binarySearch(arr []int, element int) int{
	low := 0
	high := len(arr) - 1
	mid := (low + high)/2 
	// fmt.Printf("Low = %d, mid = %d , high : %d ", low, mid, high)
	for low <= high {
	//	fmt.Println("In loop - ")
	//	fmt.Printf("Low = %d, mid = %d , high : %d ", low, mid, high)
		if arr[mid] == element {
			return mid
		}
		if element > arr[mid] {
			low = mid + 1
		}else{
			high = mid -1
		}
		mid = (low + high)/2 
	//	fmt.Println("Loop end ")
	}
	return -1
}

func main() {
	arr := []int{3, 5, 8, 29, 56, 78, 89, 123, 126, 138, 345, 567}
	el := 2123
	index := binarySearch(arr, el)
	if index > -1 {
		fmt.Printf("Element %d found at %d postion ", el, index)
	}else{
		fmt.Printf("Element %d not found  ", el)		
	}
}