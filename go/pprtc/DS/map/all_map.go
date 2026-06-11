RULES TO REMEMBER
✔ Map KEY must be comparable
Allowed as keys:

All numeric types (int, float64, etc.)

string

bool

array

struct (only if all fields are comparable)

pointer

channel

custom types (type aliases of comparable types)

functions (NO) → not allowed

slices (NO) → not allowed

maps (NO) → not allowed

interfaces (YES) → but they MUST hold comparable values

✔ Map VALUE can be anything
Allowed:

basic types

struct

slice

map

function

interface

pointer

channel

any custom type

🎯 ALL PERMUTATIONS OF MAP KEY × VALUE
Below is a complete catalog — each item is a legal Go map.

======================
✅ 1. BASIC KEY TYPES
======================
KEY: int
map[int]int{}
map[int]string{}
map[int]float64{}
map[int]bool{}
map[int][]int{}
map[int]map[string]int{}
map[int]func(int) int{}
map[int]struct{}{}
map[int]interface{}{}
map[int]User{}               // custom struct
map[int]*User{}              // pointer
KEY: string
map[string]int{}
map[string]string{}
map[string]float64{}
map[string]bool{}
map[string][]string{}
map[string]map[int]bool{}
map[string]func(string) int{}
map[string]struct{}{}
map[string]interface{}{}
map[string]User{}
KEY: bool
map[bool]int{}
map[bool]string{}
map[bool][]string{}
map[bool]map[string]int{}
map[bool]func(int) bool{}
KEY: float64
map[float64]int{}
map[float64]string{}
map[float64][]int{}
map[float64]map[int]string{}
========================
✅ 2. STRUCT AS KEY
========================
Struct is comparable only if all fields are comparable.

type Point struct {
    X int
    Y int
}

map[Point]int{}
map[Point]string{}
map[Point][]int{}
map[Point]map[string]int{}
map[Point]User{}
========================
✅ 3. ARRAY AS KEY
========================
Array is always comparable.

map[[3]int]string{}
map[[10]byte]int{}
map[[2]string]float64{}
=================================
✅ 4. POINTER AS KEY
=================================
Pointers are comparable.

map[*int]string{}
map[*User]int{}
map[*float64][]string{}
================================
✅ 5. CHANNEL AS KEY
================================
Channels are comparable.

map[chan int]string{}
map[chan bool]int{}
================================
✅ 6. INTERFACE AS KEY
================================
Only allowed if stored values are comparable.

map[interface{}]int{}
map[interface{}]string{}
map[interface{}][]string{}       // allowed
⚠️ But storing non-comparable keys inside interface (like slices) will panic.

================================
🔥 VALUE TYPES: ALL PERMUTATIONS
================================
Value can be ANYTHING; here are all categories:

A. Value = basic types
map[string]int{}
map[int]string{}
map[bool]float64{}
B. Value = slice
map[string][]int{}
map[int][]string{}
map[User][]User{}
C. Value = map
map[string]map[int]string{}
map[int]map[string]interface{}{}
D. Value = struct
map[string]User{}
map[int]Point{}
E. Value = function
map[string]func(int) int{}
map[int]func(string) bool{}
F. Value = interface
map[int]interface{}{}
map[string]interface{}{}
G. Value = pointer
map[string]*User{}
map[int]*int{}
H. Value = channel
map[string]chan int{}
map[int]chan bool{}
⭐ MASTER TABLE: Key × Value Permutations
Key Type	Allowed Values	Example
int	ANY	map[int][]string
string	ANY	map[string]map[int]bool
bool	ANY	map[bool]func(int)bool
float64	ANY	map[float64]User
struct	ANY	map[Point]*User
array	ANY	map[[3]int]map[string]int
pointer	ANY	map[*User][]string
channel	ANY	map[chan int]interface{}
interface	ONLY comparable keys	map[interface{}]int
🎉 BONUS: ILLEGAL MAP KEY TYPES (For Interviews)
These types CANNOT be keys:

❌ slice
❌ map
❌ function

Examples:

map[[]int]int{}      // illegal
map[map[int]int]int{}  // illegal
map[func() int]int{}   // illegal
🚀 Want FULL CODE with ALL permutations in one file?
I can generate a full Go program printing all combinations & storing real data.