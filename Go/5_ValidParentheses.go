package Go

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"
Output: true

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	// 開く括弧に対応する閉じる括弧をマップで定義
	opens := map[rune]rune{
		'(': ')', '{': '}', '[': ']',
	}

	// スタック（後入れ先出し: LIFO）
	var stack []rune

	for _, char := range s {
		// 開く括弧ならスタックに追加
		if closing, exists := opens[char]; exists {
			stack = append(stack, closing) // 対応する閉じる括弧をスタックに入れる
		} else {
			// 閉じる括弧が来た場合、スタックのトップと一致するかチェック
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			// 対応する括弧が見つかったのでスタックから削除
			stack = stack[:len(stack)-1]
		}
	}

	// すべての括弧が適切に閉じられているかチェック（スタックが空ならOK）
	return len(stack) == 0
}
