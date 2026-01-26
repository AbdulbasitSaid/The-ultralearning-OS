fn main() {
    const STARTING_MISSILES: i32 = 8;
    const READY_AMOUNT: i32 = 2;
    //bJust playing with the bellow line see what the compiler with do.
    // READY_AMOUNT = 1;
    // let mut missiles: i32 = STARTING_MISSILES;
    // let ready: i32 = READY_AMOUNT;
    // A pragmatic way of variable declaration using a tuple
    let (missiles, ready) = (STARTING_MISSILES, READY_AMOUNT);
    println!("Firing {} of my {} missiles...", ready, missiles);
    // for bellow line to work, you need to make #[missiles] in line 6 mutable using the mut keyword

    // missiles = missiles - ready;

    println!("{} missiles left", missiles - ready);
}
