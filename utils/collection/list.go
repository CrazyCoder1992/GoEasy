package collection

import "goeasy/utils/common"

// List is a collection that supports generic elements and some common
// operations.
type List[T any] struct {
	// values the type of value should be defined as generic T, but
	// built-in functions will be not able to use if you do that.
	// So we defined it as any type although this will result in performance
	// degradation of some functions.
	values []any
}

// Add Append an element to the end of list.
func (l *List[T]) Add(element T) {
	l.values = append(l.values, element)
}

// Remove delete the element from the list by index.
func (l *List[T]) Remove(index int) {
	l.values = append(l.values[0:index], l.values[index+1:]...)
}

// Get the element at the index from list.
func (l *List[T]) Get(index int) T {
	return l.values[index].(T)
}

// GetSize return the size of values.
func (l *List[T]) GetSize() int {
	return len(l.values)
}

// Values return all values of list.
func (l *List[T]) Values() []T {
	return common.ConvAny2Type[T](l.values)
}

// SetValues reset values in list.
func (l *List[T]) SetValues(values []T) {
	l.values = common.ConvType2Any[T](values)
}

// Sort values using the quicksort method.
func (l *List[T]) Sort(comparable func(a T, b T) int) {
	values := l.Values()
	QuickSort(values, comparable)
	l.SetValues(values)
}
