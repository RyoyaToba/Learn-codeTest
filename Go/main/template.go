package main

import (
	. "Learn-codeTest/Go"
	"fmt"
)

// ここに解きたい関数を貼り付ける
func solve() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Println(result)
}

func main() {
	solve()
}
