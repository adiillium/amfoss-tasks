use std::io;
extern crate regex;
use regex::Regex;

fn main(){
    let mut email = String::new();
    println!("Enter email:");
    match io::stdin().read_line(&mut email) {
        Ok(_n) => {}
        Err(error) => println!("error: {}", error),
    }
    let re = Regex::new(r"[^@]+@[^\.]+\..+").unwrap(); //very simple regex. Maybe wrong on corner cases.
    if re.is_match(&email) {
        println!("Valid email.");
    }
    else {
        println!("Invalid email.");
    }
}