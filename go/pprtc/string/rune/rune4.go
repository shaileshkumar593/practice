package main

import (
	"fmt"
)

func main() {
	s := "₹50"

	fmt.Println("Byte iteration:") // byte by byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %x -> %c\n", i, s[i], s[i])
	}

	fmt.Println("\nRune iteration:")
	for i, r := range s {
		fmt.Printf("%d: %U -> %c\n", i, r, r)
	}
}

/*
This is critical for Unicode-safe operations, because slicing a Go string by index works on bytes, not characters.

⚙️ Example String
s := "₹50"
Visually, we see 3 characters, but internally:

Character	Unicode	UTF-8 Bytes (hex)	Bytes (decimal)
₹	U+20B9	E2 82 B9	[226, 130, 185]
5	U+0035	35	[53]
0	U+0030	30	[48]
Total bytes = 3 + 1 + 1 = 5 bytes

❌ 1. Unsafe Slicing by Byte Index
fmt.Println(s[:1])
Output:

�
Why?
Because s[:1] returns only 1 byte: 0xE2 — the first byte of the 3-byte UTF-8 sequence for ₹.
That’s an invalid partial rune, so Go replaces it with the Unicode replacement character � (U+FFFD).

⚠️ 2. Partial UTF-8 slice example
fmt.Println(s[:2])
Output:

��
Still invalid — you sliced the first 2 bytes E2 82, missing the 3rd byte B9.
Go again replaces invalid bytes with replacement runes.

✅ 3. Correct Rune-Safe Slicing
To handle Unicode safely, convert the string to a []rune slice:

runes := []rune(s)
fmt.Println(string(runes[:1]))  // ₹
fmt.Println(string(runes[:2]))  // ₹5
fmt.Println(string(runes))      // ₹50
✅ Output:

₹
₹5
₹50
Now it works perfectly because:

[]rune holds each Unicode code point as an int32.

Each element corresponds to a full character (not partial bytes).

Slicing runes operates on characters, not bytes.

*/
/*
 4. Deep Dive — Why Go Uses Bytes by Default
Go strings are defined as:

type string struct {
    data []byte
    len  int
}
Strings are read-only byte sequences, not arrays of characters.

This design gives speed and memory efficiency for ASCII and binary data.

But Unicode (UTF-8) uses variable length, so direct byte slicing can break characters.

🧩 5. Safe Unicode Operations Pattern
✅ Convert once → Work with []rune → Convert back if needed

runes := []rune("₹50")
fmt.Println("Length (bytes):", len("₹50"))      // 5
fmt.Println("Length (runes):", len(runes))      // 3

// Safe character access
fmt.Printf("First char: %c\n", runes[0])        // ₹

// Safe substring
fmt.Println(string(runes[:2]))                  // ₹5
🔥 6. Benchmark: Bytes vs Runes
Operation	UTF-8 String	[]rune
len()	Bytes count	Rune count
s[i]	1 byte	1 full rune
s[:n]	Byte slice	Character slice
Performance	Faster (ASCII)	Slower (Unicode safe)
Memory	Less	More (4 bytes per rune)
🧩 7. Summary — When to Use What
Use Case	Use Type
Binary data, ASCII text	string or []byte
Human text, Unicode-safe manipulation	[]rune
Character counting, slicing, reversing	Convert to []rune
Serialization, hashing, network I/O	Keep as bytes
*/
