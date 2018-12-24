package set

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet()
	if set == nil {
		t.Errorf("NewSet was incorrect, got: nil, want: %+v.", set)
	}
}

func TestAdd(t *testing.T) {

	set := NewSet()
	set.Add("1")
	if _, c := set.m["1"]; c == false {
		t.Errorf("Add was incorrect, got: %t, want: %t.", c, true)
	}

}

func TestRemove(t *testing.T) {

	set := NewSet()
	set.m["1"] = struct{}{}
	set.Remove("1")
	if _, c := set.m["1"]; c == true {
		t.Errorf("Remove was incorrect, got: %t, want: %t.", c, false)
	}

}

func TestContains(t *testing.T) {

	set := NewSet()
	set.m["1"] = struct{}{}
	if c := set.Contains("1"); c == false {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", c, false)
	}

}
