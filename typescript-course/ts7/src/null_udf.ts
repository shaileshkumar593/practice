let a1: string | null = null;
console.log(a1 ?? "Guest");

a1 = "Shailesh"; // not null or undefined 
console.log(a1 ?? "Guest");


let b1: number | null = null;
console.log(b1 ?? 100);

b1 = 0; // not null or undefined 
console.log(b1 ?? 100);


let c: boolean | null = null;
console.log(c ?? true); // true is not null or undefined 

c = false; // not null or undefined 
console.log(c ?? true);

console.log("-----------------")
console.log(5 === 5);       // true
console.log("5" === "5");   // true
console.log(null === undefined); // false



console.log(null == undefined); // true


console.log("++++++++++++++++++++++++++++++")
console.log(Boolean(null));
console.log(Boolean(undefined));