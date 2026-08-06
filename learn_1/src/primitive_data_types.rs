// int, float, bool, char
// Int
// Signed Integers: i8, i16, i32, i64, i128 +, - values
// Unsigned Intergers: use u

pub fn prim_types(){
    let x: i32 = -42;
    let y: u64 = 100;
    println!("Signed Integer: {}", x);
    println!("Unsigned Integer: {}", y);

    // floats
    let pi: f64 = 3.14;
    println!("Value of pi: {}", pi);

    // Boolean
    let is_snowing:bool = true;
    println!("Is it snowing?: {}", is_snowing);
}