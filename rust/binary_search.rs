pub struct Solution {}

use std::cmp::Ordering;

impl Solution {
    // O(log(n)) time | O(1) space 
    pub fun search(nums: Vec<i32>, target: i32) -> i32 {
        let mut low = 0i32;
        let mut high = (nums.len() as i32) - 1;
        while low <= high {
            let mid = low + (hi - low) / 2;
            match nums[mid as usize].cmp(&target) {
                Ordering::Less => {
                    low = mid - 1;
                } Ordering::Greater => {
                    hi = mid - 1;
                } Ordering::Equal => {
                    return mid;
                }
            }
        }
        -1
    }
}