// Go program to illustrate how to 
// create the same name methods 
// with non-struct type receivers 
package main 
  
import "fmt"
  
type value_1 string 
type value_2 int
  
// Creating same name function with 
// different types of non-struct receivers 
func (a value_1) display() value_1 { 
  
    return a + "forGeeks"
} 
  
func (p value_2) display() value_2 { 
  
    return p + 298 
} 
  
// Main function 
func main() { 
  
    // Initializing the values 
    res1 := value_1("Geeks") 
    res2 := value_2(234) 
  
    // Display the results 
    fmt.Println("Result 1: ", res1.display()) 
    fmt.Println("Result 2: ", res2.display()) 
} 