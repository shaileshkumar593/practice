/*function throwError(message:string):never{
    throw new Error(message)
}

throwError("404 error") */
 // either return with error or go to infinite loop

 /*function infinityLoop(): never {
    while(true){
    console.log("Running...");
    }
 }

 infinityLoop()*/
 


 function greet(): void{
    console.log("Hello")
 }
 const result = greet()
 console.log(result)