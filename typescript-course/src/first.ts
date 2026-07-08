import * as readline from "node:readline";

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

rl.question("Enter your name: ", (name) => {
    console.log(`Hello ${name}`);
    console.log("Hello TypeScript");
    rl.close();
});

