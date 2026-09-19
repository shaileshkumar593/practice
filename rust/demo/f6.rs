fn main() {
    let x = {
        let a = 10;  // statement
        let b = 20;  // statement

        a + b        // expression → final value
    };

    println!("{}", x);
}