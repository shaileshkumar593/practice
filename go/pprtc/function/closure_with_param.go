package main

import (
	"fmt"
)

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}

/*

	What Happens When You Call multiplier(2)?
double := multiplier(2)
Step-by-step:

Step 1: factor = 2
Step 2: It returns this function:
func(x int) int {
	return x * factor
}
But here is the important part:

⚠️ That function remembers factor = 2

It forms a closure.

🧠 3️⃣ Memory-Level Explanation
When you call:

double := multiplier(2)
Internally:

Heap:
factor = 2

double → points to function + reference to factor
When you call:

triple := multiplier(3)
New memory is created:

Heap:
factor = 2   (for double)
factor = 3   (for triple)
Two separate closure environments exist.

📦 Visualization
For double:
double:
   function(x)
      uses → factor = 2
For triple:
triple:
   function(x)
      uses → factor = 3
They are completely independent.

🧠 4️⃣ Execution Flow
When you call:
double(5)
It executes:

return 5 * 2
Result: 10

When you call:
triple(5)
It executes:

return 5 * 3
Result: 15

🔥 Important Concept
Closures capture variables, not just values.

Each call to multiplier() creates a new variable factor.

Each returned function remembers its own factor.

🧠 5️⃣ Why Does factor Not Disappear?
Normally:

Local variables live on stack

When function ends → stack frame removed

But here:

The inner function is still using factor.

So Go compiler:

✅ Moves factor to heap
✅ Keeps it alive
✅ Attaches it to returned function

This is called:

🔎 Escape Analysis
You can check using:

go build -gcflags="-m"
🏗 6️⃣ Real Production Use Case
This pattern is heavily used in:

✅ Middleware
func authMiddleware(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking role:", role)
			next.ServeHTTP(w, r)
		})
	}
}
Here role behaves like factor.

✅ Retry Logic
func retry(times int) func(func() error) error {
	return func(fn func() error) error {
		for i := 0; i < times; i++ {
			if err := fn(); err == nil {
				return nil
			}
		}
		return fmt.Errorf("failed")
	}
}
⚡ 7️⃣ Performance Consideration
Closures may:

Allocate memory on heap

Increase GC pressure

Be slightly slower than plain functions

But:

In real systems (API, DB, Kafka), this overhead is negligible.

🎯 8️⃣ Key Interview Points
If interviewer asks:

👉 What is happening here?

You say:

The multiplier function returns a closure that captures the factor variable. Each call creates a separate closure environment, allowing independent state retention.

🚀 Final Mental Model
Think of it like:

multiplier(2) → creates mini object:
{
   factor: 2,
   function(x) → x * factor
}
It behaves like a struct + method without explicitly defining one.

If you want next level, I can show:

🔥 Closure vs Struct comparison

🔥 How this looks in assembly

🔥 Closure memory benchmark

🔥 How to implement same logic without closure

🔥 When NOT to use closure

Tell me how deep you want to go 😄
*/
