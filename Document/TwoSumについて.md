## Two Sum 問題まとめ

問題概要:
整数配列 nums と整数 target が与えられたとき、
nums 内の2つの数の和が target になるようなインデックスのペアを返す。
ただし、同じ要素を2回使ってはいけない。
解は必ず1つだけ存在すると仮定する。

```Go
package main

import "fmt"

// 愚直解（Brute Force）
// 配列の全てのペアを試して条件を満たすものを返す。
// 時間計算量は O(n^2) で非効率だが最初に考えるべき基礎的な方法。
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 最適解（HashMapを使う方法）
// ループ1回でHashMapに値を保存しつつ、
// 「今必要な補数」が既に登録されているかをチェック。
// 時間計算量は O(n)、空間計算量は O(n)。
func twoSumHashMap(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("=== TwoSum BruteForce ===")
	result1 := twoSumBruteForce(nums, target)
	fmt.Println(result1) // 出力例: [0 1]

	fmt.Println("=== TwoSum HashMap ===")
	result2 := twoSumHashMap(nums, target)
	fmt.Println(result2) // 出力例: [0 1]
}
```

まとめ:

Two Sum は「2つの数の和でtargetになるペアを探す」という基本問題。

- 愚直解（全探索）: 二重ループで全ペアを試す (O(n^2))
- 最適解（HashMap）: ループ1回で補数を管理し高速に探す (O(n))

HashMapを使う解法はコーディングテストで頻出で、
アルゴリズムの基礎力をつけるのに最適。

他にも2ポインタやスライディングウィンドウなど
関連した典型問題が多数あるので、順に学習すると効果的です。
