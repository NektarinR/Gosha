package sorts

import (
	"reflect"
	"testing"
)

var (
	data = []int{1, 2, 3, 4, 5, 6, 7}
)

func TestSort(t *testing.T) {
	expected := []int{7, 6, 5, 4, 3, 2, 1}
	insertionSort(data)
	if !reflect.DeepEqual(expected, data) {
		t.Errorf("expected %v got %v", expected, data)
	}
}
