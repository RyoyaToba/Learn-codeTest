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

// best answer? ↓
func isPalindrome2(x int) bool {
	// 負の数や、0 以外で末尾が 0 の数は回文にならない
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rev := 0
	for x > rev {
		// これは 整数を反転する 一般的なアルゴリズムです。
		// x % 10 は x の 最後の桁を取得する。
		// rev * 10 で 既存の rev を1桁左にずらす（繰り上げ）。
		// 最後の桁を rev に 追加。
		rev = rev*10 + x%10
		x /= 10
		// x = 123の場合
		// x % 10 → 3を取得 rev * 10 → 0 * 10 = 0 revの更新 → 0 + 3 = 3
		// x % 10 → 2を取得 rev * 10 → 3 * 10 = 30 revの更新 → 30 + 2 = 32
		// x % 10 → 1を取得 rev * 10 → 32 * 10 = 320 revの更新 → 320 + 1 = 321
	}

	// 偶数桁の場合: x == rev
	// 奇数桁の場合: x == rev/10 (中央の桁を除外)
	return x == rev || x == rev/10
}
