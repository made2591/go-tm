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
	if c := set.Contains("1"); c == false {
		t.Errorf("Add was incorrect, got: %t, want: %t.", c, true)
	}

}

func TestRemove(t *testing.T) {

	set := NewSet()
	set.Add("1")
	set.Remove("1")
	if c := set.Contains("1"); c == true {
		t.Errorf("Remove was incorrect, got: %t, want: %t.", c, false)
	}

}

func TestContains(t *testing.T) {

	set := NewSet()
	set.Add("1")
	if c := set.Contains("1"); c == false {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", c, true)
	}

	set.Remove("1")
	if c := set.Contains("1"); c == true {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", c, false)
	}

}

func TestIterator(t *testing.T) {

	set := NewSet()
	set.Add("1")
	set.Add("2")
	set.Add("3")
	elem := set.Iterator()
	for e := range elem {
		if c := set.Contains(e); c == false {
			t.Errorf("Iterator was incorrect, got: %t, want: %t.", c, true)
		}
	}

}
