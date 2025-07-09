package Go

/* https://leetcode.com/problems/longest-common-prefix/description/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	// 最小の文字数を求める
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	left, right := 0, minLen
	for left <= right {
		mid := (left + right) / 2
		if isCommonPrefix(strs, mid) {
			left = mid + 1 // もっと長い共通部分を探す
		} else {
			right = mid - 1 // もっと短い共通部分を探す
		}
	}

	return strs[0][:right]
}

// 与えられた配列の文字列がすべて指定のプレフィックスを持つかをチェック
func isCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length] // 最初の文字列のprefixを取得
	for _, str := range strs {
		if len(str) < length || str[:length] != prefix {
			return false
		}
	}
	return true
}
