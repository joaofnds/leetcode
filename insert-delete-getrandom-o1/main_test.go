package main

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		set := Constructor()

		if !set.Insert(1) {
			t.Errorf("Insert(1) should return true for new element")
		}
		if set.Insert(1) {
			t.Errorf("Insert(1) should return false for duplicate element")
		}

		if !set.Remove(1) {
			t.Errorf("Remove(1) should return true for existing element")
		}
		if set.Remove(1) {
			t.Errorf("Remove(1) should return false for non-existent element")
		}
		if !set.Insert(1) {
			t.Errorf("Insert(1) should return true after element was removed")
		}

		if set.GetRandom() != 1 {
			t.Errorf("GetRandom() should return 1 when only 1 is in the set")
		}
	})

	t.Run("Multiple operations", func(t *testing.T) {
		set := Constructor()

		set.Insert(1)
		set.Insert(2)
		set.Insert(3)

		if !set.Remove(2) {
			t.Errorf("Remove(2) should return true")
		}

		found1, found3 := false, false
		for range 100 {
			val := set.GetRandom()
			if val == 1 {
				found1 = true
			} else if val == 3 {
				found3 = true
			} else {
				t.Errorf("GetRandom() returned unexpected value: %d", val)
			}

			if found1 && found3 {
				break
			}
		}

		if !found1 || !found3 {
			t.Errorf("GetRandom() failed to return all elements: found1=%v, found3=%v", found1, found3)
		}
	})

	t.Run("Insert and Remove sequence", func(t *testing.T) {
		set := Constructor()

		if !set.Insert(10) {
			t.Errorf("Insert(10) should return true")
		}
		if !set.Insert(20) {
			t.Errorf("Insert(20) should return true")
		}
		if !set.Insert(30) {
			t.Errorf("Insert(30) should return true")
		}
		if !set.Insert(40) {
			t.Errorf("Insert(40) should return true")
		}

		if !set.Remove(20) {
			t.Errorf("Remove(20) should return true")
		}
		if !set.Remove(10) {
			t.Errorf("Remove(10) should return true")
		}
		if !set.Remove(40) {
			t.Errorf("Remove(40) should return true")
		}

		for range 10 {
			if set.GetRandom() != 30 {
				t.Errorf("GetRandom() should return 30")
			}
		}
	})

	t.Run("Empty set operations", func(t *testing.T) {
		set := Constructor()

		if set.Remove(1) {
			t.Errorf("Remove(1) should return false on empty set")
		}

		_ = set.GetRandom() // GetRandom on empty set should not crash

		// Insert and then Remove to empty the set again
		set.Insert(5)
		set.Remove(5)

		if !set.Insert(5) {
			t.Errorf("Insert(5) should return true after emptying the set")
		}
	})
}
