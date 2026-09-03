let a1: string | null = null;
console.log(a1 ?? "Guest");

a1 = "Shailesh";
console.log(a1 ?? "Guest");


let b1: number | null = null;
console.log(b1 ?? 100);

b1 = 0;
console.log(b1 ?? 100);


let c: boolean | null = null;
console.log(c ?? true);

c = false;
console.log(c ?? true);