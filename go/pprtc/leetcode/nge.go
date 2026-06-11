package main

import "fmt"

// nextGreaterSteps returns steps to next greater element for each index
func nextGreaterSteps(arr []int) []int {
	n := len(arr)
	res := make([]int, n)

	stack := []int{} // store indices

	// Traverse from right to left
	for i := n - 1; i >= 0; i-- {

		// Remove all elements <= current
		/*
			They can never be next greater

			Current element is bigger → they are useless
		*/
		fmt.Println(stack)

		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		// Assign result
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1] - i
		}

		// Push current index
		stack = append(stack, i)
		fmt.Println(res)
	}

	return res
}

func main() {
	//arr := []int{84, 74, 21, 19, 75, 66, 44, 55, 70, 34, 32, 29, 92, 4, 6, 105, 101}
	arr := []int{4, 5, 2, 10, 8}

	result := nextGreaterSteps(arr)

	fmt.Println("Input :", arr)
	fmt.Println("Output:", result)
}

/*

	Why do we move from RIGHT → LEFT?
🔥 Key Reason
We need information about the FUTURE (right side)
❌ If you go LEFT → RIGHT
👉 At index i, you don’t know:

What is the next greater element on the right?
👉 You would need:

Check all elements on right → O(n²)
✅ If you go RIGHT → LEFT
👉 At index i, you already know:

All elements to the right have been processed
👉 So:

You can use a stack to store useful candidates
🔥 Core Idea
Stack = elements to the RIGHT that are useful
👉 While moving left:

Remove useless elements

Keep only possible "next greater"

🔍 Visual Intuition

https://miro.medium.com/v2/resize%3Afit%3A1400/1%2AViHQy6mhuRifDrskooOXdQ.jpeg

https://miro.medium.com/1%2AZVKYIiFEaaV6GZrwCl5JAA.png

https://cdn-images-1.medium.com/max/600/1%2ABjfehW7Njkm7KvtVe6HFOg.png
4
🔄 Step-by-Step Example
Array:
[84, 74, 75, 66, 70, 92]
Step 1: Start from last
92 → no right → answer = -1
stack = [92]
Step 2: Move left → 70
stack top = 92 (>70) → answer = 92
stack = [92, 70]
Step 3: Move left → 66
stack top = 70 (>66) → answer = 70
stack = [92, 70, 66]
Step 4: Move left → 75
pop 66 (smaller)
pop 70 (smaller)
top = 92 → answer = 92
stack = [92, 75]
🎯 Key Observation
Stack always stores decreasing elements
👉 So:

Top of stack = next greater element

🔥 Why it works
When processing element i:
Stack already contains candidates from RIGHT
👉 So you:

Remove smaller elements (useless)

Top = next greater

Push current

⚡ Time Complexity Insight
Each element:
- pushed once
- popped once
👉 Total:

O(n)
🧠 Interview Answer (Perfect)
👉 If interviewer asks:

“Why traverse from right?”

Answer:

Because we need information about elements on the right.
By traversing from right to left, we ensure that when we
process an element, all elements to its right have already
been processed and are available in the stack.
🚀 Alternative (LEFT → RIGHT?)
👉 Possible, but harder:

You need to track unresolved elements

More complex logic

👉 RIGHT → LEFT is:

Cleaner + intuitive + optimal
🎯 Final Takeaway
We move from right because NGE depends on future elements,
and stack helps us reuse that information efficiently.


arr = [4,5,2,10,8]

Start from right:

i=4 → 8 → stack empty → -1
stack = [4]

i=3 → 10 → pop 8 → stack empty → -1
stack = [3]

i=2 → 2 → next greater = 10
stack = [3,2]

i=1 → 5 → pop 2 → next greater = 10
stack = [3,1]

i=0 → 4 → next greater = 5
stack = [3,1,0]
*/
