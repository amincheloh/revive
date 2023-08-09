// Test of empty-blocks.

package fixtures

import (
	"fmt"
)

func g(f func() bool) {
	{ // MATCH /this block is empty, you can remove it/
	}

	_ = func(e error) {} // Must not match

	if ok := f(); ok { // MATCH /this block is empty, you can remove it/
		// only a comment
	} else {
		println("it's NOT empty!")
	}

	if ok := f(); ok {
		println("it's NOT empty!")
	} else { // MATCH /this block is empty, you can remove it/

	}

	for i := 0; i < 10; i++ { // MATCH /this block is empty, you can remove it/

	}

	for { // MATCH /this block is empty, you can remove it/

	}

	for true { // MATCH /this block is empty, you can remove it/

	}

	for true { // MATCH /this block is empty, you can remove it/

	}

	// issue 386
	var c = make(chan int)
	for range c { // DO NOT FAIL
		for range c { // DO NOT FAIL
		}

		// But without types it skips this (too artificial?) one
		var s = "a string"
		for range s { // DO NOT FAIL (false negative)
		}

		select {
		case _, ok := <-c:
			if ok { // MATCH /this block is empty, you can remove it/
			}
		}

		// issue 810
		next := 0
		iter := func(v *int) bool {
			*v = next
			next++
			fmt.Println(*v)
			return next < 10
		}

		z := 0
		for iter(&z) { // Must not match
		}

		for process() { // Must not match
		}

		var it iterator
		for it.next() { // Must not match
		}
	}
}

func process() bool {
	return false
}

type iterator struct{}

func (it *iterator) next() bool {
	return false
}
