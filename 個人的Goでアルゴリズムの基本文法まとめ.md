個人的Goでアルゴリズムとデータ構造の基本文法まとめ

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

two pointersに似ているがsliding windowはスライス内の特定の範囲に着目し、範囲をずらして比較するを繰り返すアルゴリズムです。

例えばスライスの中n個の部分スライスの和の最大を求めるコードを考える。

```go

func maxSubarraySum(array []int, n int) int {
	// nがスライスの長さよりも大きければ0で終了
	if n > len(array) {
		return 0
	}

	maxSum := 0 // 求めたい最大の和
	tempSum := 0 // n要素のスライスの要素の和

	// 0からn-1番目までの要素の和を求める
	for i := 0; i < n; i++ {
		maxSum += array[i]
	}
	tempSum = maxSum

	for i := n; i < len(array); i++ {
		// i-n番目の要素を引いてi番目を足せば次の部分スライスの我が求められる
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

時間計算量は**O(N)**で空間計算量は**O(1)**となる。

## stack

Stackは最後に入れた要素が最初に取り出されるデータ構造。

LIFO(Last In First Out)ともいわれる。

例えばコードを書いているときに括弧の個数が間違っていないかを判定するときに使える。

与えられた文字列のうち`{}`, `[]`, `()`の対応が正しいかを判定するコードを考えてみます。

文字の種類は`{}[]()`のいずれかのみとします

```go
func isValid(s string) bool {
	// それぞれの括弧閉じるの対応する括弧開けるをmapで定義しておく
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	// stackを定義。今回はスライスで行う。
	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		// 文字が閉じ括弧でない場合、スタックに追加する
		pair, ok := pairs[char]
		if !ok {
			stack = append(stack, char)
			continue
		}

		// スタックが空である場合は無効な括弧の組み合わせのためfalseを返す
		if len(stack) == 0 {
			return false
		}

		// スタックの最後の要素が対応する開き括弧でない場合は無効な括弧の組み合わせのためfalseを返す
		if stack[len(stack) - 1] != pair {
			return false
		}

		// スタックの最後の要素が対応する開き括弧である場合、その要素をスタックから削除
		stack = stack[:len(stack) - 1]
	}

	// スタックが空であればすべての括弧が正しくペアになっているためtrue
	// スタックが空でない場合、一部の開き括弧に対応する閉じ括弧がないためfalseを返す
	return len(stack) == 0
}
```

このようにすると`()[]{}`はtrue、`([)]`はfalseが帰るようになります。

sの文字列の長さをNとすると、時間計算量は**O(N)**、空間計算量は**O(N)**となる

## Queue

Queueは最初に入れられた要素が最初に取り出されるデータ構造。

FIFO(First In First Out)と呼ばれることもあり、Stackとよく対比される。

使い所はtreeのBFS(Breadth First Search)で処理を行いたいとき。

以下の二分木を考える。

```
    1
   / \
  2   3
 / \ / \
4  5 6  7
```

各要素をnodeと称するが、nodeは左右に別のnodeへのアクセスができるとする。

例えば1の要素はleftは2、rightは3といったデータ構造になる。

これを1, 2, 3, ..., 7と順番に処理をしたいときのコードは以下のようになる。

```go
type node struct {
	val int
	left *node
	right *node
}

func bfs(root *node) []int {
	if root == nil {
		return []int{}
	}

	// nodeのポインタのqueueを作成
	queue := []*node{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		// dequeue
		// queueの先頭要素を取得
		current := queue[0]
		// queueの先頭要素を削除
		queue = queue[1:]
		result = append(result, current.val)

		// enqueue
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return result
}

func main() {
	root := &node{val: 1}
	root.left = &node{val: 2}
	root.right = &node{val: 3}
	root.left.left = &node{val: 4}
	root.left.right = &node{val: 5}
	root.right.left = &node{val: 6}
	root.right.right = &node{val: 7}

	fmt.Println(bfs(root))
}
```

先頭の要素を取得して削除する行為を`dequeue`, 末尾に要素を挿入する行為を`enqueue`と言う。

Nをnode数として

- 時間計算量は二分木の各ノードを一度だけ訪問するのみのため時間計算量は**O(N)**
- 空間計算量はキューとしてスライスを使用し、最悪の場合（すべてのノードが同じレベルにある場合）、スライスの長さはノード数と同じになるため空間計算量は**O(N)**


## Binary Search

Binary Searchはソート済みの配列に対してターゲットとなる要素を取り出すアルゴリズム。

intのスライスの中にtargetがあるかを判定し、存在すればtargetを、存在しなければ-1を返す関数を考える。

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for left <= right {
		// 中央をmidとして取得
        mid := (left + right) / 2
        if nums[mid] > target {
			// ターゲットがmidよりも小さい場合は左側にスコープを絞る
            right = mid - 1
        } else if nums[mid] < target {
			// ターゲットがmidよりも大きい場合は右側にスコープを絞る
            left = mid + 1
        } else {
			// targetと同じならmidを返す
            return mid
        }
    }
    return -1
}
```

numsの長さをNとすると

- 時間計算量：二分探索は各ステップで探索範囲を半分に狭めるため**O(log N)**

- 空間計算量：left、right、midの3つの変数のみ使うため**O(1)**

## linked list
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

## heap

```go
```

## backtrack

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

## DP

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

## greedy

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

## interval


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
