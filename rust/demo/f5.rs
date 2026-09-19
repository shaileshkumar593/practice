fn main() {
  let logged_in = true;
  let is_admin = false;

  println!("Is regular user: {}", logged_in && !is_admin);
  println!("Has any access: {}", logged_in || is_admin);
  println!("Not logged in: {}", !logged_in);
}