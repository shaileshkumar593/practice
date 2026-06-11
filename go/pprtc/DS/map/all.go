ALL TYPES OF MAP DECLARATION
✅ 1. Nil map declaration
var m map[string]int
m == nil

cannot write to it until initialized

reading is allowed (returns zero value)

✅ 2. Empty map (make)
m := make(map[string]int)
✅ 3. Map with initial values
m := map[string]int{
    "a": 1,
    "b": 2,
}
✅ 4. Map with make + capacity
m := make(map[int]string, 10)   // optional
Capacity is not guaranteed, but gives Go a hint.

✅ 5. Map of slices
m := map[string][]int{
    "nums": {1, 2, 3},
}
✅ 6. Map of maps
m := map[string]map[string]int{
    "outer": {"inner": 10},
}
✅ 7. Map of structs
type User struct{ Age int }
m := map[string]User{
    "john": {Age: 30},
}
✅ 8. Map with interface values
m := map[string]interface{}{
    "age": 30,
    "name": "john",
}
✅ 9. Map with function values
m := map[string]func(int) int{
    "square": func(x int) int { return x * x },
}
✅ 10. Map of any custom type
type ID int
m := map[ID]string{
    1: "hello",
}
/*
🎯 Summary Table
| Type  | Declaration                  | Notes         |
| ----- | ---------------------------- | ------------- |
| Slice | `var s []int`                | nil slice     |
| Slice | `s := []int{1,2}`            | literal       |
| Slice | `s := make([]int, 5)`        | pre-sized     |
| Slice | `s := arr[1:3]`              | slicing       |
| Slice | `s := arr[1:3:5]`            | full slice    |
| Map   | `var m map[string]int`       | nil map       |
| Map   | `m := make(map[string]int)`  | initialized   |
| Map   | `m := map[string]int{}`      | literal empty |
| Map   | `m := map[string]int{"a":1}` | with values   |

*/
