個人的Goでアルゴリズムの基本文法まとめ

# アルゴリズムめっちゃ大事

エンジニアにとってCSの基礎知識を学んでいない筆者が効率的に処理を書くとはを自慢気に語れるようになりたいと思い

個人で学習したアルゴリズムの書き方をまとめる。

解説は必要と思ったタイミングでは入れますが、詳細には書かれていないので学びたい方は更にググるかコメントください。

## 超基本のmap

mapはGoには標準搭載されている。mapはO(1)で値にアクセスできる優れものである。

例としてスライスの中に重複があるかを判定する関数を考えてみよう。

何も考えずに

```
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

two pointer
```
func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1

    for l < r {
        for l < r && !isAlphaNum(s[l]) {
            l++
        }
        for l < r && !isAlphaNum(s[r]) {
            r--
        }
        if getLowerLetterByte(s[l]) != getLowerLetterByte(s[r]) {
            return false
        }
        l++
        r--
    }
    return true
}

func isAlphaNum(b byte) bool {
    return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func getLowerLetterByte(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        return b + byte('a' - 'A')
    }
    return b
}
```

sliding window

```
package main

import (
	"fmt"
)

func maxSubarraySum(array []int, n int) int {
//Write a check for the edge case if the array's length is smaller than the subarray's length. If it is then we return 0.
	if n > len(array) {
		return 0
	}
	maxSum := 0 //Assign a variable that will hold the highest sum
	tempSum := 0 //Assign a variable that will hold the sum of the current subarray we're looking at.

	//Create a loop that will Loop through 1 time to get the sum of the first set of numbers
	for i := 0; i < n; i++ {
  //Set maxSum to equal the sum of our first loop
		maxSum += array[i]
	}
	//Now that we have a max sum, assign temp sum to max sum
	tempSum = maxSum

	//Loop through the length of the whole array starting from the index of the next number after the first set
	for i := n; i < len(array); i++ {
		//assigning new tempsum, we get the current tempSum, move down the number of indexes == to n, subtract that number, then add the number of the current loops index
    fmt.Println(tempSum, array[i-n], array[i]) //Visual representation when ran to see the values so you can see whats going on.
		tempSum = tempSum - array[i-n] + array[i]
		//see which is larger the current maxSum or the tempSum we're holding, if tempSum is larger maxSum is now the value of TempSum
		maxSum = Max(maxSum, tempSum)
	}
	return maxSum
}

//helper function to get the higher of two integers
func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(maxSubarraySum([]int{-1, -1, -2, 4, 2, 3, 5, 1}, 4))
	fmt.Println(maxSubarraySum([]int{}, 4))
	fmt.Println(maxSubarraySum([]int{100, 200, 300, 400}, 2))
}
```

stack

```
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

```
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

```
```

backtrack

```
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

```
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

```
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
