fn main(){
    println!("Hello World");
    let x = 5;
    //x = 10; // Error  cannot assign twice to immutable variable
    println!("{}", x);

    //  help: remove this `mut`
    let mut y = 5;
    println!("{}", y);  // y is never changed after initialization. therefore ask to remove mut 

    let mut z = 12;
    z = 25;
    println!("{}",z);

}