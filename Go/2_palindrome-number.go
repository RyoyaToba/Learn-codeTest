package Go

import (
	"strconv"
)

/*
Given an integer x, return true if x is a
palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

// My first answer↓
func isPalindrome(x int) bool {
	str := strconv.Itoa(x) // 文字列に変換
	dig := len(str)        // 桁数
	half := len(str) / 2   // 半分

	firstHalf := str[:half]
	var secondHalf string

	// 奇数桁か偶数桁かで後半部分の切り方を変える
	if dig%2 == 0 {
		secondHalf = str[half:] // 偶数桁ならそのまま
	} else {
		secondHalf = str[half+1:] // 奇数桁なら中央の文字をスキップ
	}

	for i := 0; i < half; i++ {
		if firstHalf[i] != secondHalf[half-1-i] {
			return false
		}
	}

	return true
}
