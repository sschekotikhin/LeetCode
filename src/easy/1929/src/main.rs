use std::vec::Vec;

fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
    let mut result = nums.clone();
    result.append(&mut nums.clone());

    result
}

fn main() {
    assert_eq!(get_concatenation(vec![1, 2, 1]), vec![1, 2, 1, 1, 2, 1]);
    assert_eq!(get_concatenation(vec![1, 3, 2, 1]), vec![1, 3, 2, 1, 1, 3, 2, 1]);
}
