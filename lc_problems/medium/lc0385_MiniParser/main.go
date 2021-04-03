/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

//  func getNumber(s string, pos *int) int {
//     result := 0
//     sign := 1
//     if s[*pos] == '-' {
//         sign = -1
//         *pos += 1
//     }
//     for *pos < len(s) && s[*pos] >= '0' && s[*pos] <= '9' {
//         result *= 10
//         result += int(s[*pos] - '0')
//         *pos++
//     }
//     return sign * result
// }

// func makeNestedInteger(num int) *NestedInteger {
//     result := &NestedInteger{}
//     result.SetInteger(num)
//     return result
// }
// func parseNumber(s string, pos *int) *NestedInteger {
//     num := getNumber(s, pos)
//     return makeNestedInteger(num)
// }

// func parseList(s string, pos *int) *NestedInteger {
//     *pos++

//     result := &NestedInteger{}
//     for *pos < len(s) {
//         if s[*pos] == '[' {
//             result.Add(*parseList(s, pos))
//         } else if s[*pos] == ']' {
//             *pos++
//             break
//         } else if s[*pos] == ',' {
//             *pos++
//         } else {
//             result.Add(*parseNumber(s, pos))
//         }
//     }
//     return result
// }

// func deserialize(s string) *NestedInteger {
//     pos := 0
//     if s[pos] == '[' {
//         return parseList(s, &pos)
//     }
//     return parseNumber(s, &pos)
// }