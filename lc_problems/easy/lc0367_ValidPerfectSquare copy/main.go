package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

//  func guessNumber(n int) int {
//     low := 1
//     high := n
//     for low < high {
//         mid := low + (high - low) >> 1

//         ans := guess(mid)
//         if ans == 0 {
//             return mid
//         }
//         if ans > 0 {
//             low = mid + 1
//         } else {
//             high = mid - 1
//         }
//     }
//     return low
// }
