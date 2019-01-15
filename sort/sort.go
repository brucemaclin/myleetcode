package main

import "fmt"

func selectSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if data[j] < data[i] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
func less(data []int, i, j int) bool {
	return data[i] < data[j]
}
func bubbleSort(data []int) {
	length := len(data)
	sorted := false
	for i := length - 1; i > 0 && !sorted; i-- {
		sorted = true
		for j := 0; j < i; j++ {
			if less(data, j+1, j) {
				swap(data, j, j+1)
				sorted = false
			}
		}
	}
}
func insertSort(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && less(data, j, j-1); j-- {
			swap(data, j, j-1)
		}
	}

}
func shellSort(data []int) {
	length := len(data)
	var h = 1
	for h < length/3 {
		h *= 3
	}
	for h > 1 {
		for i := h; i < length; i++ {
			for j := i; j > 0 && less(data, j, j-h); j -= h {
				swap(data, j, j-h)
			}
		}
		h = h / 3
	}
}

func mergeSort(data []int, low, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	mergeSort(data, low, mid)
	mergeSort(data, mid+1, high)
	merge(data, low, mid, high)
}
func merge(data []int, low, mid, high int) {
	var tmp = make([]int, len(data))
	i := low
	j := mid + 1
	for l := low; l <= high; l++ {
		tmp[l] = data[l]
	}
	for k := low; k <= high; k++ {
		if i > mid {
			data[k] = tmp[j]
			j++
		} else if j > high {
			data[k] = tmp[i]
			i++
		} else if data[i] <= tmp[j] {
			data[k] = tmp[i]
			i++
		} else {
			data[k] = tmp[j]
			j++
		}
	}
}
func partion(data []int, lo, hi int) int {
	i := lo
	j := hi + 1
	v := data[i]
	for {
		for {
			i++
			if i == hi {
				break
			}
			if !(data[i] < v) {
				break
			}
		}
		for {
			j--
			if !(data[j] > v) {
				break
			}
		}
		if i >= j {
			break
		}
		swap(data, i, j)
	}
	swap(data, lo, j)
	return j
}
func quickSort(data []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partion(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)

}
func threeWayQuickSort(data []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt := lo
	i := lo + 1
	gt := hi
	v := data[lo]
	for i <= gt {
		cmp := data[i] - v
		if cmp < 0 {
			swap(data, lt, i)
			lt++
			i++
		} else if cmp > 0 {
			swap(data, i, gt)
			gt--
		} else {
			i++
		}
	}
	threeWayQuickSort(data, lo, lt-1)
	threeWayQuickSort(data, gt+1, hi)

}
func main() {
	data := []int{1, 3, 4, 2, 7, 6}
	selectSort(data)
	fmt.Println(data)
	bubbleSort(data)
	fmt.Println(data)
	insertSort(data)
	fmt.Println(data)
	shellSort(data)
	fmt.Println(data)
	mergeSort(data, 0, len(data)-1)
	fmt.Println(data)
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)
	threeWayQuickSort(data, 0, len(data)-1)
	fmt.Println(data)
}
