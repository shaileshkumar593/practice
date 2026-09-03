let num1 :number[] = [70,25,42,25]
console.log(num)


let numArry :Array<number> =[62,24,84]
console.log(numArry)

console.log(numArry.push(2012))
console.log(numArry.push(2026))
numArry.push(2425)

console.log(numArry.pop())


let thrice : number[] = num1.map(number => number * 3)
console.log(thrice)

let fltr :number[] = num1.filter(number => number > 25);

console.log(fltr)

let sumofval : number = num1.reduce(
    (sum, val) => sum + val
);

console.log(sumofval)

