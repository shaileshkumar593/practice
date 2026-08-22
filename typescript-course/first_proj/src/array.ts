let num :number[] = [70,25,42,25]
console.log(num)


let numArry :Array<number> =[62,24,84]
console.log(numArry)

console.log(numArry.push(2012))
console.log(numArry.push(2026))
numArry.push(2425)

console.log(numArry.pop())


let thrice : number[] = num.map(number => number * 3)
console.log(thrice)

let fltr :number[] = num.filter(number => number > 25);

console.log(fltr)

let sumofval : number = num.reduce(
    (sum, val) => sum + val
);

console.log(sumofval)

