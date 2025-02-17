package Go

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
