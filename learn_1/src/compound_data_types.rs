// Compound data types
// arrays, tuples, slices, strings

// Arrays
pub fn data_types(){
    let numbers = [1, 2, 3, 4, 5];
    for x in numbers{
        println!("{}", x);
    }

    let numbers:[i32;5] = [1, 2, 3, 4, 5];
    println!("Number array: {:?}", numbers);
    
}