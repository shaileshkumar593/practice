1. Use strings.Builder for Concatenation
Best for building strings from pieces.

Avoids repeated allocations.

var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("Go")
}
result := b.String()
Why fast: Only allocates once or resizes minimally, unlike s += "Go" in a loop.

2. Use []byte or []rune for In-place Modifications
Convert string to []byte or []rune.

Modify elements directly.

Convert back to string once done.

s := "hello"
b := []byte(s)
b[0] = 'H'
s = string(b)
Why fast: Avoids creating new strings for each change.

3. Avoid Repeated Slicing in Loops
Slicing strings repeatedly creates new strings (copies bytes).

Instead, use indexes and ranges.

s := "abcdef"
for i := 0; i < len(s); i++ {
    c := s[i] // byte access
    _ = c
}
4. Use strings Package Functions
Functions like strings.Join, strings.ReplaceAll, strings.Contains are optimized internally.

parts := []string{"a", "b", "c"}
s := strings.Join(parts, ",") // faster than manual concatenation
5. Use copy for Bulk Modifications
When working with slices ([]byte), copy is faster than building strings iteratively.

src := []byte("hello")
dst := make([]byte, len(src))
copy(dst, src)
6. Use unicode/utf8 for Rune-aware Processing
For strings containing Unicode, avoid naive indexing.

Use utf8.DecodeRuneInString or for _, r := range s.

s := "₹50"
for i := 0; i < len(s); {
    r, size := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%c ", r)
    i += size
}
7. Pre-allocate Buffers When Possible
If you know the final size, allocate upfront.

buf := make([]byte, 0, 1024)
buf = append(buf, []byte("hello")...)


| Method                    | Use Case                | Performance Benefit                   |
| ------------------------- | ----------------------- | ------------------------------------- |
| `strings.Builder`         | Concatenation in loops  | Minimal allocations                   |
| `[]byte` / `[]rune`       | In-place modification   | Avoid repeated string copies          |
| `strings.Join`            | Concatenate slices      | Optimized internal implementation     |
| `utf8.DecodeRuneInString` | Unicode-aware iteration | Efficient rune decoding               |
| Pre-allocated buffer      | Large modifications     | Avoids resizing/allocating repeatedly |
| `copy`                    | Duplicate slices        | Faster bulk copy than manual append   |
