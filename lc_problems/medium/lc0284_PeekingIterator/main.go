/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

//  type PeekingIterator struct {
//     iter *Iterator
//     val int
// }

// func Constructor(iter *Iterator) *PeekingIterator {
//     return &PeekingIterator{iter, 0}
// }

// func (this *PeekingIterator) hasNext() bool {
//     return this.val > 0 || this.iter.hasNext()
// }

// func (this *PeekingIterator) next() int {
//     if this.val > 0 {
//         result := this.val
//         this.val = 0
//         return result
//     }
//     return this.iter.next()
// }

// func (this *PeekingIterator) peek() int {
//     if this.val > 0 {
//         return this.val
//     }
//     this.val = this.iter.next()
//     return this.val
// }