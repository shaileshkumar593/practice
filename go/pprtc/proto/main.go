package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

// Auto-generated from swap.proto
// Normally you’d generate this via protoc, but included inline for simplicity.

/*
This struct is automatically compatible with google.golang.org/protobuf/proto methods such as:

proto.Marshal() → convert struct → binary

proto.Unmarshal() → binary → struct

*/

type Numbers struct {
	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return fmt.Sprintf("&Numbers{A:%d, B:%d}", m.A, m.B) }
func (*Numbers) ProtoMessage()    {}

func main() {
	// Original numbers
	nums := &Numbers{A: 10, B: 20}
	fmt.Printf("Before Swap: A = %d, B = %d\n", nums.A, nums.B)

	// Encode to protobuf bytes
	data, err := proto.Marshal(nums)
	if err != nil {
		log.Fatalf("Encoding failed: %v", err)
	}

	// Decode from protobuf bytes
	decoded := &Numbers{}
	if err := proto.Unmarshal(data, decoded); err != nil {
		log.Fatalf("Decoding failed: %v", err)
	}

	// Swap numbers
	decoded.A, decoded.B = decoded.B, decoded.A
	fmt.Printf("After Swap:  A = %d, B = %d\n", decoded.A, decoded.B)
}

/*
   ┌───────────────┐
│ swap.proto    │   ← defines message schema
└──────┬────────┘
       │ (protoc generates)
       ▼
┌────────────────────┐
│ proto.Numbers      │
│ { A int32, B int32 }│
└────────┬───────────┘
         │
         │ proto.Marshal()
         ▼
┌────────────────────┐
│ Binary Data        │  (numbers.bin)
└────────┬───────────┘
         │
         │ proto.Unmarshal()
         ▼
┌────────────────────┐
│ Go Struct Restored │
│ {A: 10, B: 20}     │
└────────┬───────────┘
         │ Swap A,B
         ▼
┌────────────────────┐
│ {A: 20, B: 10}     │
└────────────────────┘
*/
