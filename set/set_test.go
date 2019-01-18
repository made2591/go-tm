package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet()
	if set == nil {
		t.Errorf("NewSet was incorrect, got: nil, want: %+v.", set)
	}
}

func TestAdd(t *testing.T) {

	set := NewSet()
	set.Add("1")
	if !set.Contains("1") {
		t.Errorf("Add was incorrect, got: %t, want: %t.", false, true)
	}

}

func TestRemove(t *testing.T) {

	set := NewSet()
	set.Add("1")
	set.Remove("1")
	if set.Contains("1") {
		t.Errorf("Remove was incorrect, got: %t, want: %t.", true, false)
	}

}

func TestContains(t *testing.T) {

	set := NewSet()
	set.Add("1")
	if !set.Contains("1") {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", false, true)
	}

	set.Remove("1")
	if set.Contains("1") {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", true, false)
	}

}

func TestCardinality(t *testing.T) {

	set := NewSet()
	set.Add("1")
	set.Add("2")
	set.Add("3")
	set.Add("3")
	if set.Cardinality() != 3 {
		t.Errorf("GetOne was incorrect, got: %d, wanted: %d", set.Cardinality(), 3)
	}

}

func TestIterator(t *testing.T) {

	set := NewSet()
	set.Add("1")
	set.Add("2")
	set.Add("3")
	elem := set.Iterator()
	for _, e := range elem {
		if c := set.Contains(e); !c {
			t.Errorf("Iterator was incorrect, got: %t, want: %t.", c, true)
		}
	}

}
