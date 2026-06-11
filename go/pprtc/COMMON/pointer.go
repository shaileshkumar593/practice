package main

import (
	"fmt"
)

func main() {
	var p *int
	var q int = 20
	p = &q

	fmt.Println(p, &p, *p)

	var r = new(int)
	fmt.Println(r, &r, *r)
	var t = &q
	fmt.Println(t, &t, *t)

	var mp *map[string]string
	mp = new(map[string]string)
	fmt.Println(mp, &mp, *mp)

	var mapMake = make(map[string]string)
	fmt.Println(mapMake, &mapMake)

	fmt.Println("slice ")
	var ms *[]int
	ms = new([]int)
	//fmt.Println(ms, &ms, (*ms)[0])//panic: runtime error: index out of range [0] with length 0
	fmt.Println(ms, &ms, *ms)
	//fmt.Println(len(*ms)) //runtime error: index out of range [0] with length 0
	//fmt.Println(cap(*ms))  //panic: runtime error: index out of range [0] with length 0

	/*
		(*ms)[0] = 20
		(*ms)[1] = 29
		fmt.Println(ms, &ms, *ms)
		due to
	*/
	fmt.Println("ms1")
	var ms1 *[]int
	ms1 = new([]int)
	fmt.Println(ms1, &ms1, *ms1)
	fmt.Println()

	*ms1 = make([]int, 5, 10)
	fmt.Println(len(*ms1), &ms1, *ms1)
	(*ms1)[0] = 30
	(*ms1)[1] = 31
	fmt.Println(ms1, &ms1, *ms1)

	var mp2 *map[string]string
	mp2 = new(map[string]string)
	//*mp2 = make(map[string]string) / / if this line is omitted, it will pan "Pan: assignment to entry in nil map"“
	(*mp2)["name"] = "lc"
	fmt.Println((*mp2)["name"])
}

/*
As can be seen above, silce, map, channel and other types are reference types.
 When the reference type is initialized to nil, nil can not be assigned directly,
 nor can new be used to allocate memory, but also need to use make to allocate.*/

/*
/*
/ The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//  Slice: The size specifies the length. The capacity of the slice is
//  equal to its length. A second integer argument may be provided to
//  specify a different capacity; it must be no smaller than the
//  length. For example, make([]int, 0, 10) allocates an underlying array
//  of size 10 and returns a slice of length 0 and capacity 10 that is
//  backed by this underlying array.
//  Map: An empty map is allocated with enough space to hold the
//  specified number of elements. The size may be omitted, in which case
//  a small starting size is allocated.
//  Channel: The channel's buffer is initialized with the specified
//  buffer capacity. If zero, or the size is omitted, the channel is
//  unbuffered.


func make(t Type, size ...IntegerType) Type
*/

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.

//func new(Type) *Type

/*
Simply put, new just allocates memory, not initializes memory; make allocates and initializes memory.
The so-called initialization is to assign an initial value to a type, for example, the character is empty,
 the integer is 0, and the logical value is false.

 It can be seen that its parameter is a type, the return value is a pointer to the memory address
 of this type, and the allocated memory will be set to zero, that is, the zero value of the type,
 that is, the character is empty, the integer is 0, and the logical value is false
*/
