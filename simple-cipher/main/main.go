package main

import (
	"fmt"
	"math"
)

/*func main() {
	message := "iamapandabear"
	key := "abcd"
	//	for i := range message {
	//		b := (message[i] - 'a' + 3) % 26
	//		fmt.Printf("%s::%s\n", string(message[i]), string(b+'a'))
	//	}
	//var sb strings.Builder
	//var shiftDistance int
	//shiftDistance = 3
	for i := range message {
		sb.WriteByte('a' + (message[i]-'a'+byte(shiftDistance))%26)
	}
	fmt.Println(sb.String())
	newKey := ""
	if len(message) > len(key) {
		newKey = strings.Repeat(key, len(message)/len(key))
	}
	leftOver := len(message) - (len(message)/len(key))*len(key)
	if leftOver > 0 {
		newKey += key[0:leftOver]
	}
	fmt.Println(newKey)
}

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	mid := 0
	for lo < hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if nums[lo] < target && nums[hi] > target {
		return lo + 1
	}
	return mid
}*/

func sortedSquares(nums []int) []int {

	left, right := 0, len(nums)-1
	result := make([]int, len(nums))

	for index := len(nums) - 1; index >= 0; index-- {
		square := 0
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		result[index] = square * square
	}
	return result

}

func main() {
	//var arr = []int{-4, -1, 0, 3, 10}
	var arr = []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(arr))
}
