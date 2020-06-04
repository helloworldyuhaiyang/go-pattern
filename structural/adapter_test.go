package structural

import "testing"

func TestAdapter_Input5V(t *testing.T) {
	familyPower := NewFamilyPower()
	adapter := NewAdapter(*familyPower)
	adapter.Input5V()
}
