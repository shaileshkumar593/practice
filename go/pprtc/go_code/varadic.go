package main

import("fmt"
		"strings")


func joinstr(element...string)string{

return strings.Join(element,"-")
}

func main(){
element := []string{"geeks","FOR","geeks"}
fmt.Println(joinstr(element...))

}
