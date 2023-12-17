個人的Goでアルゴリズムの基本文法まとめ

# アルゴリズムめっちゃ大事

エンジニアにとってCSの基礎知識を学んでいない筆者が効率的に処理を書くとはを自慢気に語れるようになりたいと思い

個人で学習したアルゴリズムの書き方をまとめる。

解説は必要と思ったタイミングでは入れますが、詳細には書かれていないので学びたい方は更にググるかコメントください。

## 超基本のmap

mapはGoには標準搭載されている。mapはO(1)で値にアクセスできる優れものである。

例として要素数N個のスライス中に重複する要素があるかを判定する関数を考えてみよう。

最初に筆者が考えたのは

- 一番目の要素を二番目以降の要素と比較して重複があるか判定
- 二番目の要素を三番目以降の要素と比較して重複があるか判定
...
- N-1番目の要素をN番目以降の要素と比較して重複があるか判定

この処理だと時間計算量は**O(N^2)**となり効率は良くはありません。

これをmapを使う方法で考えてみます。

- mapを定義
- スライスをループして、mapのキーにindexの要素があるか重複判定。重複がなければmapのキーにindexの要素を追加

これをすることで計算量は**O(N)**となる。空間計算量は**O(N)**。

コードを見てみよう。


```go
func containsDuplicate(nums []int) bool {
    hm := make(map[int]struct{})

    for i := range nums {
        if _, ok := hm[nums[i]]; ok {
            return true
        }
        hm[nums[i]] = struct{}{}
    }
    return false
}
```

mapのvalueにはstruct{}{}を入れているが、もちろんtrueを入れてもOK。


## two pointers

two pointersはleftとrightの2つのポインタを使ったテクニックになる。

パターンは色々あるが回分の判定を例に見てみよう。

回分はある文字列が左から順に読んでも右から順に読んでも一致する状態であり、英語では`palindrome`という。

条件として引数の文字列は小文字の英語のみとします。

```go
func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1

    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}
```

時間計算量は**O(N)**で空間計算量は**O(1)**

## sliding window

two pointersに似ているがsliding windowはスライス内の特定の範囲に着目し、範囲をずらして比較するを繰り返すアルゴリズム。

例えばスライスの中nこの部分スライスの和の最大を求めるコードを考える。

```go

func maxSubarraySum(array []int, n int) int {

	if n > len(array) {
		return 0
	}

	maxSum := 0
	tempSum := 0

	for i := 0; i < n; i++ {

		maxSum += array[i]
	}

	tempSum = maxSum


	for i := n; i < len(array); i++ {
		tempSum = tempSum - array[i-n] + array[i]
		maxSum = max(maxSum, tempSum)
	}
	return maxSum
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
```

時間計算量は**O(N)**で空間計算量は**O(1)**


## stack

```go
func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, ok := pairs[char]
		if !ok {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack) - 1] != pair {
			return false
		}

		stack = stack[:len(stack) - 1]
	}

	return len(stack) == 0
}
```

binary search

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            return mid
        }
    }
    return -1
}```

linked list
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedListHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedListHead
}

// Iterative version
// func reverseList(head *ListNode) *ListNode {
//     var prev *ListNode
//     curr := head
//
//     for curr != nil {
//         tmp := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = tmp
//     }
//
//     return prev
// }
```

heap

```go
```

backtrack

```go
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		ans = append(ans, append([]int{}, curr...))
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)
	return ans
}
```

DP

```go
func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		sum := one + two
		one, two = two, sum
	}

	return two
}
```

greedy
```
func maxSubArray(nums []int) int {
    res := nums[0]
    currentSum := 0

    for i := range nums {
        if currentSum < 0 {
            currentSum = 0
        }

        currentSum += nums[i]
        res = max(currentSum, res)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

interval

```go
import "sort"

func merge(intervals [][]int) [][]int {
	const (
		start = 0
		end   = 1
	)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][start] < intervals[j][start]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		currResLastIndex := len(res) - 1
		currEnd := res[currResLastIndex]
		curr := intervals[i]

		if currEnd[end] < curr[start] {
			res = append(res, curr)
		} else if currEnd[end] < curr[end] {
			res[currResLastIndex][end] = curr[end]
		}
	}

	return res
}
```
