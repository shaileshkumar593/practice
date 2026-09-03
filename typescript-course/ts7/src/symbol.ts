let id1  = Symbol("secret")

console.log(id1)

let id2 = Symbol("secret")

console.log(id2)

console.log(id1 == id2 )

console.log("-----------------")
console.log(5 === 5);       // true
console.log("5" === "5");   // true
console.log(null === undefined); // false



console.log(null == undefined); // true


console.log("++++++++++++++++++++++++++++++")
console.log(Boolean(null));
console.log(Boolean(undefined));