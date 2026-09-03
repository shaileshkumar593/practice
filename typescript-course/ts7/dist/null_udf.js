"use strict";
let a1 = null;
console.log(a1 ?? "Guest");
a1 = "Shailesh";
console.log(a1 ?? "Guest");
let b1 = null;
console.log(b1 ?? 100);
b1 = 0;
console.log(b1 ?? 100);
let c = null;
console.log(c ?? true);
c = false;
console.log(c ?? true);
console.log("-----------------");
console.log(5 === 5); // true
console.log("5" === "5"); // true
console.log(null === undefined); // false
console.log(null == undefined); // true
console.log("++++++++++++++++++++++++++++++");
console.log(Boolean(null));
console.log(Boolean(undefined));
