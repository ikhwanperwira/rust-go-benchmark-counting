use std::time::Instant;

fn main() {
    let start_time = Instant::now();

    for i in 1..=1000000 {
        println!("{}", i);
    }

    let duration = start_time.elapsed();
    println!("Execution time: {:?}", duration);
}
