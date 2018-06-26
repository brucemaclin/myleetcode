package main

func isPerfectSquare(num int) bool {
	low := 1
	high := num
	for low <= high {
		middle := int64((low + high) >> 1)
		big := middle * middle
		if int(big) == num {
			return true
		} else if int(big) < num {
			low = int(middle + 1)
		} else {
			high = int(middle - 1)
		}

	}
	return false

}
func main() {

	isPerfectSquare(16)

}
