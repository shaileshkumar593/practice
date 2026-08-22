let valany:any = "shailesh";
console.log(valany)

valany = 123;

console.log(valany)

valany = true

console.log(valany)


//console.log(valany.toUpperCase()) // throw error 

let valuknown:unknown = "TS";
if (typeof(valuknown) === "string"){
    console.log(valuknown.toUpperCase());
}